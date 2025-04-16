package main

import (
	"fmt"
	"huffman-coding/hnode"
)

func main() {
	text := "abac"

	tree := hnode.BuildTree(text)
	hnode.PrettyPrint(tree, "", "")

	for r, b := range tree.Coding {
		fmt.Printf("%q: %s\n", r, b)
	}

	var result string
	tree.Encode(&result)
	fmt.Printf("I: '%s'\n", text)
	fmt.Printf("O: '%s'\n", result)
}
