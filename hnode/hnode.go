package hnode

import (
	"fmt"
	"huffman-coding/models"
	"sort"
	"unicode/utf8"
)

type HNode struct {
	Left   *HNode
	Right  *HNode
	Value  models.HSymbol
	Coding map[rune]*models.BitStream
	text   string
}

func BuildTree(text string) *HNode {
	chars := textToSymbols(text)
	trees := make([]*HNode, len(chars))
	for i, char := range chars {
		trees[i] = &HNode{Value: char}
	}

	for len(trees) > 1 {
		sort.Slice(trees, func(i, j int) bool {
			return trees[i].Value.Freq < trees[j].Value.Freq
		})

		l := trees[0]
		r := trees[1]

		comb := &HNode{
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

func (node *HNode) GenCodes() map[rune]*models.BitStream {
	codes := make(map[rune]*models.BitStream)
	genCodesRecusrive(node, models.NewBitStream(), codes)
	return codes
}

//TODO:
// doesn't work for one character e.g. 'a' because the rec func gets
//  called once, with an empty slice. (handle case)

func genCodesRecusrive(node *HNode, bs *models.BitStream, codes map[rune]*models.BitStream) {
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

func textToSymbols(text string) []models.HSymbol {
	m := make(map[rune]int)
	for _, c := range text {
		m[c]++
	}

	var chars []models.HSymbol
	for char, count := range m {
		chars = append(chars, models.HSymbol{Value: []rune{char}, Freq: count})
	}

	return chars
}

func (node *HNode) Encode(result *models.BitStream) error {
	bitStream := models.NewBitStream()

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

func (node *HNode) SerializeTree() *models.BitStream {
	bs := models.NewBitStream()

	if isLeaf(node) {
		// Means this is a leaf node thing
		bs.AppendBit(true)

		for _, r := range node.Value.Value {
			buf := make([]byte, 4)
			n := utf8.EncodeRune(buf, r)
			utf8bytes := buf[:n]

			// Length (2 bits)
			bs.AppendInt(uint32(n-1), 2)

			// BYTES BYTES BYTES BYTES or byte actually
			for _, b := range utf8bytes {
				bs.AppendInt(uint32(b), 8)
			}
		}
	} else {
		// Means this is a tree node. so just encode left, then right.
		// i mean, recursive panicking is fun, right...
		bs.AppendBit(false)
		l := node.Left.SerializeTree()
		fmt.Printf("l: %v\n", l)
		if err := bs.AppendBitStream(l); err != nil {
			// TODO: dont panic.
			panic(err)
		}
		r := node.Right.SerializeTree()
		fmt.Printf("r: %v\n", r)
		if err := bs.AppendBitStream(r); err != nil {
			panic(err)
		}

	}
	return bs
}

func (node *HNode) Decode(encoded *models.BitStream) (string, error) {
	reverseMap := make(map[string]rune)
	for r, bs := range node.Coding {
		reverseMap[bs.String()] = r
	}

	var result []rune
	curr := models.NewBitStream()

	for i := range encoded.BitCount {
		bit, _ := encoded.ReadBitAt(i)
		curr.AppendBit(bit)

		if r, ok := reverseMap[curr.String()]; ok {
			result = append(result, r)
			curr = models.NewBitStream()
		}
	}

	return string(result), nil
}

// func (node *HNode) EncodeTree(result *models.BitStream) error {
// 	bs := models.NewBitStream()
//
// 	if isLeaf(node) {
// 		bs.AppendBit(true)
// 		v := node.Value.Value[0]
// 	}
//
// 	return nil
// }
