package main

import (
	"fmt"
	"huffman-coding/bitstream"
	"huffman-coding/huffman"
	"os"
)

func main() {
	sr := bitstream.NewStreamReader(os.Stdin)

	for {
		bit, err := sr.Read()
		fmt.Printf("bit: %v\n", bit)
		if err != nil {
			panic(err)
		}
		if !bit {
			break
		}
	}

	tree, err := huffman.DeserializeTree(sr)
	if err != nil {
		panic(err)
	}
	huffman.PrettyPrint(tree, "", "")

	tree.Coding = tree.GenCodes()
	text, err := tree.Decode(sr)
	if err != nil {
		panic(err)
	}

	for _, r  := range text {
		fmt.Printf("%s", string(r))
	}
}
