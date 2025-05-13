package main

import (
	"fmt"
	"huffman-coding/bitstream"
)

func main() {
	bs := bitstream.NewBitStream()
	bs2 := bitstream.NewBitStream()

	bs.AppendNBits(true, 3)
	bs.AppendBit(false)
	bs.AppendInt(15, 4)
	fmt.Printf("bs: %v\n", bs)

	bs2.AppendBit(true)
	bs2.AppendBit(false)
	bs2.AppendBit(false)
	fmt.Printf("bs2: %v\n", bs2)

	// fmt.Printf("sm: %v%v\n", bs.String()[:len(bs.String())-4], bs2.String()[:len(bs2.String())-4])
	bs.AppendBitStream(bs2)
	fmt.Printf("bs: %v\n", bs)
}
