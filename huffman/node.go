package huffman

import (
	"fmt"
	"huffman-coding/bitstream"
	"sort"
	"unicode/utf8"
)

type Node struct {
	Left   *Node
	Right  *Node
	Value  Symbol
	Coding map[rune]*bitstream.BitStream
	text   string
}

func BuildTree(text string) *Node {
	chars := textToSymbols(text)
	trees := make([]*Node, len(chars))
	for i, char := range chars {
		trees[i] = &Node{Value: char}
	}

	for len(trees) > 1 {
		sort.Slice(trees, func(i, j int) bool {
			return trees[i].Value.Freq < trees[j].Value.Freq
		})

		l := trees[0]
		r := trees[1]

		comb := &Node{
			Left:  l,
			Right: r,
			Value: l.Value.Combine(r.Value),
		}

		trees = append(trees[2:], comb)
	}

	t := trees[0]
	t.text = text
	t.Coding = t.GenCodes()

	return t
}

func (node *Node) GenCodes() map[rune]*bitstream.BitStream {
	codes := make(map[rune]*bitstream.BitStream)
	genCodesRecusrive(node, bitstream.NewBitStream(), codes)
	return codes
}

//TODO:
// doesn't work for one character e.g. 'a' because the rec func gets
//  called once, with an empty slice. (handle case)

func genCodesRecusrive(node *Node, bs *bitstream.BitStream, codes map[rune]*bitstream.BitStream) {
	if isLeaf(node) {
		codes[node.Value.Value[0]] = bs.Clone()
		return
	}

	if node.Left != nil {
		leftCode := bs.Clone()
		leftCode.AppendBit(false)
		genCodesRecusrive(node.Left, leftCode, codes)
	}

	if node.Right != nil {
		rightCode := bs.Clone()
		rightCode.AppendBit(true)
		genCodesRecusrive(node.Right, rightCode, codes)
	}
}

func textToSymbols(text string) []Symbol {
	m := make(map[rune]int)
	for _, c := range text {
		m[c]++
	}

	var chars []Symbol
	for char, count := range m {
		chars = append(chars, Symbol{Value: []rune{char}, Freq: count})
	}

	return chars
}

func (node *Node) Encode(result *bitstream.BitStream) error {
	bitStream := bitstream.NewBitStream()

	for _, c := range node.text {
		if bs, ok := node.Coding[c]; !ok {
			panic("oopsie woopsie~")
		} else {
			for i := range bs.BitCount {
				bit, _ := bs.ReadBitAt(i)
				bitStream.AppendBit(bit)
			}
		}
	}

	*result = *bitStream
	return nil
}

func (node *Node) SerializeTree() *bitstream.BitStream {
	bs := bitstream.NewBitStream()

	if isLeaf(node) {
		// Means this is a leaf node thing
		bs.AppendBit(true)

		for _, r := range node.Value.Value {
			buf := make([]byte, 4)
			n := utf8.EncodeRune(buf, r)
			utf8bytes := buf[:n]
			fmt.Printf("UTF: %s -> ", string(node.Value.Value))
			for _, b := range utf8bytes {
				fmt.Printf("%d %08b", b, b)
			}
			fmt.Println()

			// Length (2 bits)
			// bs.AppendInt(uint32(n-1), 2)

			// BYTES BYTES BYTES BYTES or byte actually
			for _, b := range utf8bytes {
				bs.AppendInt(uint32(b), 8)
				fmt.Printf("bs: %v\n", bs)
			}
		}
	} else {
		// Means this is a tree node. so just encode left, then right.
		// i mean, recursive panicking is fun, right...
		bs.AppendBit(false)
		l := node.Left.SerializeTree()
		// fmt.Printf("l: %v\n", l)
		if err := bs.Add(l); err != nil {
			// TODO: dont panic.
			panic(err)
		}
		r := node.Right.SerializeTree()
		// fmt.Printf("r: %v\n", r)
		if err := bs.Add(r); err != nil {
			panic(err)
		}

	}
	return bs
}

func (node *Node) Decode(encoded *bitstream.BitStream) (string, error) {
	reverseMap := make(map[string]rune)
	for r, bs := range node.Coding {
		reverseMap[bs.String()] = r
	}

	var result []rune
	curr := bitstream.NewBitStream()

	for i := range encoded.BitCount {
		bit, _ := encoded.ReadBitAt(i)
		curr.AppendBit(bit)

		if r, ok := reverseMap[curr.String()]; ok {
			result = append(result, r)
			curr = bitstream.NewBitStream()
		}
	}

	return string(result), nil
}
