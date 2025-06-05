package main

import (
	"fmt"
	"huffman-coding/bitstream"
	"huffman-coding/huffman"
	"os"
)

func main() {
	sr := bitstream.NewStreamReader(os.Stdin)
	// b, _ := io.ReadAll(os.Stdin)
	// fmt.Printf("b[0]: %08b\n", b[0])

	// for range 10  {
	// 	b, _ := sr.Read()
	// 	fmt.Printf("b: %v\n", b)
	// }
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

	// huffman.PrettyPrint(tree, "", "")
	// tree.Decode()
	// if err != nil {
	// 	panic(err)
	// }
	// huffman.PrettyPrint(deserialized, "", "")
	//
	// deserialized.Coding = deserialized.GenCodes()
	//
	// s, err := deserialized.Decode(&result)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("s: %v\n", s)
	//
}
