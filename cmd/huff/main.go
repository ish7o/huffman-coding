package main

import (
	"bufio"
	"huffman-coding/bitstream"
	"huffman-coding/huffman"
	"io"
	"os"
)

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

	var r bitstream.BitStream

	err = tree.Encode(&r, f)
	if err != nil {
		panic(err)
	}

	t := tree.SerializeTree()
	t.Add(&r)

	padding := 8 - (t.BitCount % 8)
	prefix := bitstream.NewBitStream()

	for i := 1; i < padding; i++ {
		prefix.AppendBit(true)
	}
	prefix.AppendBit(false)

	prefix.Add(t)

	os.Stdout.Write(prefix.GetBytes())
}
