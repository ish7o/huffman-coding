package main

import (
	"bufio"
	"fmt"
	"huffman-coding/bitstream"
	"huffman-coding/huffman"
	"io"
	"os"
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

func doShit() (*os.File, map[rune]int, error) {
	file, err := os.CreateTemp("", "huffman")
	if err != nil {
		return nil, nil, err
	}

	rd := bufio.NewReader(os.Stdin)
	// var freq map[rune]int
	freq := make(map[rune]int)

	for {
		r, _, err := rd.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, err
		}
		freq[r]++

		_, err = file.WriteString(string(r))
		if err != nil {
			return nil, nil, err
		}
	}

	if _, err := file.Seek(0, 0); err != nil {
		return nil, nil, err
	}

	return file, freq, nil
}

func main() {
	f, fq, err := doShit()
	if err != nil {
		panic(err)
	}
	defer f.Close()
	defer os.Remove(f.Name())

	tree := huffman.BuildTree(fq)
	// huffman.PrettyPrint(tree, "", "")

	// for r, b := range tree.Coding {
	// 	fmt.Printf("%q: %s\n", r, b)
	// }
	//
	var r bitstream.BitStream

	// 01100010
	// 		2 + 32 + 64
	err = tree.Encode(&r, f)
	if err != nil {
		panic(err)
	}

	// huffman.PrettyPrint(tree, "", "")

	// fmt.Printf("I: '%s' (%d bytes)\n", text, len(text))
	// fmt.Printf("O: 0b%v (%d bits)\n", byteConv(result.Bytes), result.BitCount)
	// fmt.Printf("Compressed (%.3f)\n", float64(len(text))/float64(result.BitCount*8))

	// decoded, _ := tree.Decode(&result)
	// fmt.Printf("Decoded: '%s' (%d bytes)\n", decoded, len(decoded))

	t := tree.SerializeTree()
	// fmt.Printf("byteConv(t.GetBytes()): %v\n", byteConv(t.GetBytes()))
	t.Add(&r)
	fmt.Print(string(t.GetBytes()))
	// fmt.Printf("byteConv(t.GetBytes()): %v\n", byteConv(t.GetBytes()))
	// for _, b := range t.GetBytes() {
	// 	for i := 7; i >= 0; i-- {
	// 		bit := (b >> i) & 1
	// 		fmt.Print(bit)
	// 	}
	// }
	// fmt.Printf("%v", t.GetBytes())



	// os.Stdout.Write(serialized.GetBytes())
	// fmt.Print(serialized.Value())
	// fmt.Println()
	// fmt.Print(r.GetBytes())

	// fmt.Printf("serialized: %v\n", serialized)

	// deserialized, err := huffman.DeserializeTree(serialized.NewReader())
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
