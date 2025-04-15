package main

import (
	"fmt"
	"huffman-coding/hnode"
)

func main() {
    text := "A_DEAD_DAD_CEDED_A_BAD_BABE_A_BEADED_ABACA_BED"

    tree := hnode.BuildTree(text)
    // hnode.PrettyPrint(tree, "", "")

    for r, b := range tree.Coding {
        fmt.Printf("%q: ", r)
        for _, bit := range b {
            if bit {
                fmt.Print("1")
            } else {
                fmt.Print("0")
            }
        }
        fmt.Println()
    }

    var result string
    tree.Encode(&result)
    fmt.Printf("I: '%s'\n", text)
    fmt.Printf("O: '%s'\n", result)
}
