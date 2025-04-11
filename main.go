package main

import (
	"fmt"
	"huffman-coding/btree"
	"sort"
)

type Char struct {
	Value []rune
	Freq  int
}

func (c Char) String() string {
	return fmt.Sprintf("'%s': %d", string(c.Value), c.Freq)
}

func buildTree(chars []Char) *btree.BTree[Char] {
    trees := make([]*btree.BTree[Char], len(chars))
    for i, char := range chars {
        trees[i] = &btree.BTree[Char]{Value: char}
    }

    for len(trees) > 1 {
        sort.Slice(trees, func(i, j int) bool {
            return trees[i].Value.Freq < trees[j].Value.Freq
        })

        for i, t := range trees {
            fmt.Printf("%d. %s\n", i+1, string(t.Value.Value))
        }
        fmt.Println()
        // sort.Slice(chars, func(i, j int) bool {
        //     return chars[i].Freq < chars[j].Freq
        // })

        l := trees[0]
        r := trees[1]

        comb := &btree.BTree[Char]{
            Left: l,
            Right: r,
            Value: Char{
                Value: append(l.Value.Value, r.Value.Value...),
                Freq: l.Value.Freq + r.Value.Freq,
            },
        }

        trees = append(trees[2:], comb)
    }

    return trees[0]
}

func textToChar(text string) []Char {
	m := make(map[rune]int)
	for _, c := range text {
		m[c]++
	}

	var chars []Char
	for char, count := range m {
		chars = append(chars, Char{Value: []rune{char}, Freq: count})
	}

	return chars
}

func popFirst[T any](s *[]T) (*T, bool) {
	if len(*s) == 0 {
		return nil, false
	}

	popped := (*s)[0]
	*s = (*s)[1:]
	return &popped, true
}

func main() {
	text := "A_DEAD_DAD_CEDED_A_BAD_BABE_A_BEADED_ABACA_BED"
	chars := textToChar(text)

	for _, c := range chars {
		fmt.Println(c)
	}
    btree.PrettyPrint(buildTree(chars), "", "")
}
