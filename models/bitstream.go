package models

import (
	"errors"
	"fmt"
	"strings"
)

type BitStream struct {
	Bytes    []byte
	BitCount int
}

func NewBitStream() *BitStream {
	return &BitStream{
		Bytes:    []byte{0},
		BitCount: 0,
	}
}

func (b *BitStream) AppendBit(bit bool) {
	// [11100111, 10011011, 11000000]; bitCount=18
	//                        ^
	// bit = 1
	// byteIndex = 2
	// bitPos = 5
	byteIndex := b.BitCount / 8
	bitPos := 7 - (b.BitCount % 8)

	for byteIndex >= len(b.Bytes) {
		b.Bytes = append(b.Bytes, 0)
	}

	// 00100000
	mask := byte(1 << bitPos)

	if bit {
		// 11000000 | 00100000
		//    ^         ^
		b.Bytes[byteIndex] |= mask
	} else {
		// 11000000 & 00100000
		//    ^         ^
		b.Bytes[byteIndex] &= ^mask
	}

	b.BitCount++
}

func (b *BitStream) Clone() *BitStream {
	newBytes := make([]byte, len(b.Bytes))
	copy(newBytes, b.Bytes)

	b2 := &BitStream{
		Bytes:    newBytes,
		BitCount: b.BitCount,
	}

	return b2
}

func (b *BitStream) GetBytes() []byte {
	// [11100111, 10011011, 11000000]; bitCount=18
	// fullBytes = (18 + 7) / 8 = 25/8 = 3
	fullBytes := (b.BitCount + 7) / 8

	if fullBytes == len(b.Bytes) {
		return b.Bytes
	}

	return b.Bytes[:fullBytes]
}

func (b *BitStream) ReadBitAt(position int) (bool, error) {
	// [11100111, 10011011, 11000000]; bitCount=18
	//  01234567  01234567  01234567
	//              ^
	// position = 11
	// byteIndex = position / 8 = 1
	byteIndex := position / 8
	// bitIndex = position % 8 = 3
	bitIndex := 7 - position%8

	if position > b.BitCount {
		return false, errors.New("bit position out of range :c")
	}

	//          11000000 & 00001000 -> 0 -> false
	//              ^          ^
	// fmt.Printf("b.Bytes   = %08b\nshift     = %08b\n", b.Bytes[byteIndex], 1<<bitIndex)
	return b.Bytes[byteIndex]&(1<<bitIndex) != 0, nil
}

func (b *BitStream) Value() string {
	if b.BitCount == 0 {
		return ""
	}

	var sb strings.Builder
	sb.Grow(b.BitCount)

	for i := range b.BitCount {
		bit, _ := b.ReadBitAt(i)
		if bit {
			sb.WriteRune('1')
		} else {
			sb.WriteRune('0')
		}
	}

	return sb.String()
}

func (b *BitStream) String() string {
	if b.BitCount == 0 {
		return "BitStream <empty>"
	}

	return fmt.Sprintf("%s (%d)", b.Value(), b.BitCount)
}
