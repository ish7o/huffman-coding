package main

import (
	"fmt"
	"huffman-coding/bitstream"
	"huffman-coding/huffman"
	"strings"
)

func byteConv(bytes []byte) string {
    var sb strings.Builder
    sb.Grow(len(bytes) * 8)

    for _, bt := range bytes {
        sb.WriteString(fmt.Sprintf("%08b", bt))
    }

    return sb.String()
}

func main() {
	text := "“uwuuwuuwuuwuuwu“"
	// text := "aba"

	tree := huffman.BuildTree(text)
	huffman.PrettyPrint(tree, "", "")

	for r, b := range tree.Coding {
		fmt.Printf("%q: %s\n", r, b)
	}

	var result bitstream.BitStream

	// 01100010
	// 		2 + 32 + 64
	tree.Encode(&result)

	fmt.Printf("I: '%s' (%d bytes)\n", text, len(text))
    fmt.Printf("O: 0b%v (%d bits)\n", byteConv(result.Bytes), result.BitCount)
    fmt.Printf("Compressed (%.3f)\n", float64(len(text)) / float64(result.BitCount * 8))

	decoded, _ := tree.Decode(&result)
	fmt.Printf("Decoded: '%s' (%d bytes)\n", decoded, len(decoded))

	serialized := tree.SerializeTree()
	fmt.Printf("serialized: %v\n", serialized)
}
