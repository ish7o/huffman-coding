package hnode

import (
    "bytes"
    "huffman-coding/models"
    "sort"
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
    if node.Left == nil && node.Right == nil && len(node.Value.Value) == 1 {
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

func (node *HNode) Encode(result *string) error {
    var out bytes.Buffer

    for _, c := range node.text {
        if bs, ok := node.Coding[c]; !ok {
            panic("oopsie woopsie~")
        } else {
            out.WriteString(bs.Value())
        }
    }

    *result = out.String()
    return nil
}
