package hnode

import (
	"bytes"
	"fmt"
	"huffman-coding/models"
	"sort"
);

type HNode struct {
    Left *HNode
    Right *HNode
    Value models.HSymbol
    Coding map[rune][]bool
    text string
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

        for i, t := range trees {
            fmt.Printf("%d. %s\n", i+1, string(t.Value.Value))
        }
        fmt.Println()

        l := trees[0]
        r := trees[1]

        comb := &HNode{
            Left: l,
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

func (node *HNode) GenCodes() map[rune][]bool {
    codes := make(map[rune][]bool)
    genCodesRecusrive(node, []bool{}, codes)
    return codes
}

func genCodesRecusrive(node *HNode, b []bool, codes map[rune][]bool) {
    if node.Left == nil && node.Right == nil && len(node.Value.Value) == 1 {
        codeCopy := make([]bool, len(b))
        copy(codeCopy, b)
        codes[node.Value.Value[0]] = codeCopy
        return
    }

    if node.Left != nil {
        leftCode := append(b, false)
        genCodesRecusrive(node.Left, leftCode, codes)
    }

    if node.Right != nil {
        rightCode := append(b, true)
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

func (node *HNode) Encode(result *string) error {
    var out bytes.Buffer

    for _, c := range node.text {
        if _, ok := node.Coding[c]; !ok {
            panic("oopsie woopsie~")
        }

        v := btb(node.Coding[c])
        out.WriteString(v)
    }

    *result = out.String()
    return nil
}

func btb(b []bool) string {
    var out bytes.Buffer

    for _, v := range b {
        if v {
            out.WriteString("1")
        } else {
            out.WriteString("0")
        }
    }

    return out.String()
}

