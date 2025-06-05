package huffman

import (
	"errors"
	"fmt"
	"huffman-coding/bitstream"
	// "strings"
	"unicode/utf8"
)

//	func printBytes(bytes []byte, idx int) {
//		fmt.Println("TL765432107654321076543210TL76543210L76543210")
//		for _, b := range bytes {
//			fmt.Printf("%08b", b)
//		}
//		fmt.Println()
//		fmt.Printf("%s^ (%d)\n", strings.Repeat(" ", idx), idx)
//	}

func DeserializeTree(br bitstream.Reader) (*Node, error) {
	// fmt.Printf("br.Bytes: %v\n", br.Bytes)
	// printBytes(br.Bytes, br.ReadIdx)
	// 011110001010000000100111000101100001101100010
	// TL765432107654321076543210TL76543210L76543210

	// first = false
	// readIdx: 1
	// 011110001010000000100111000101100001101100010000
	//  ^ (1)
	// first = true
	// r: “
	// donereadIdx: 31
	// 011110001010000000100111000101100001101100010000
	//                                ^ (31)
	// first = false
	// readIdx: 32
	// 011110001010000000100111000101100001101100010000
	//                                 ^ (32)
	// first = false
	// readIdx: 33
	// 011110001010000000100111000101100001101100010000
	//                                  ^ (33)
	// first = false
	// readIdx: 34
	// 011110001010000000100111000101100001101100010000
	//                                   ^ (34)
	// first = false
	// readIdx: 35
	// TL765432107654321076543210TL76543210L76543210---
	// 011110001010000000100111000101100001101100010000
	//                                    ^ (35)
	//
	// bC = 0 and newBS: 10110001 (8)
	// READCHAR GOT r: 65533
	// r: �
	first, err := br.Read()
	if err != nil {
		return nil, errors.New("BitReader is empty")
	}

	node := &Node{}

	// if is leaf node
	// then read character and pack it to a leaf node
	if first {
		fmt.Println("first = true")
		r, err := readChar(br)
		if err != nil {
			return nil, err
		}
		fmt.Printf("r: %v\n", string(r))
		node.Value = Symbol{
			Value: []rune{r},
		}

		// if tree node
		// parse left node
		// parse right node
		// put them as kids to node
	} else {
		fmt.Println("first = false")
		// fmt.Printf("readIdx: %d\n", br.ReadIdx)
		n, err := DeserializeTree(br)
		if err != nil {
			return nil, err
		}
		node.Left = n
		// fmt.Printf("n: %v\n", n)
		// fmt.Printf("readIdx: %d\n", br.ReadIdx)
		n, err = DeserializeTree(br)
		if err != nil {
			return nil, err
		}
		node.Right = n

		node.Value.Value = append(node.Left.Value.Value, node.Right.Value.Value...)
		// fmt.Printf("n: %v\n", n)
		// fmt.Println("---")
	}

	// fmt.Printf("done\n")
	// printBytes(br.Bytes, br.ReadIdx)
	// _, _ = br.Read()

	return node, nil
}

// func readChar(br *bitstream.BitReader) (rune, error) {
// 	var firstByte []bool
// 	for i := 0; i < 8; i ++ {
// 		bit, err := br.Read()
// 		if err != nil {
// 			return ' ', err
// 		}
// 		firstByte = append(firstByte, bit)
// 	}
//
// 	// count leading 1s
// 	byteCount := 0
// 	for _, bit := range firstByte {
// 		if bit {
// 			byteCount++
// 		}
// 	}
// }

// taka notka
// mam wrazenie ze nie czytam 1 gdy mam leaf node
// dlatego
// a w sumie nie wiem
// dziwne jakies to
func readChar(br bitstream.Reader) (rune, error) {
	newBS := bitstream.NewBitStream()
	byteCount := 0
	for {
		b, err := br.Read()
		// fmt.Printf("reading bit: %v\n", b)
		if err != nil {
			return ' ', err
		}

		if !b {
			break
		}
		byteCount++
		newBS.AppendBit(true)
	}
	newBS.AppendBit(false)

	// fmt.Printf("newBS: %v\n", newBS)

	// grab the rest of first byte
	bCount := byteCount
	for 7-bCount > 0 {
		b, err := br.Read()
		if err != nil {
			return ' ', err
		}
		newBS.AppendBit(b)
		bCount++
	}

	// fmt.Printf("newBS: %v\n", newBS)
	// fmt.Printf("byteCount: %v\n", byteCount)
	if byteCount == 1 {
		for range 6 { // 6 bo czytam ten o dlugosci i ten false potem
			b, err := br.Read()
			if err != nil {
				return ' ', err
			}
			newBS.AppendBit(b)
			// fmt.Printf("bC = 0 and newBS: %v\n", newBS)
		}
	} else {

		for byteCount > 1 {
			for range 8 {
				b, err := br.Read()
				if err != nil {
					return ' ', err
				}
				newBS.AppendBit(b)
			}
			byteCount--
		}
	}

	r, _ := utf8.DecodeRune(newBS.Bytes)
	// fmt.Printf("utf8.DecodeRune -> r: %v\n", r)

	// fmt.Printf("newBS: %v\n", newBS)

	return r, nil
}
