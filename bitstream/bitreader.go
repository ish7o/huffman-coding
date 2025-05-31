package bitstream

import (
	"io"
)

type BitReader struct {
	*BitStream
	ReadIdx int
}

func NewBitReader(bs *BitStream) *BitReader {
	return &BitReader{
		BitStream: bs,
		ReadIdx: 0,
	}
}

func (b *BitReader) Read() (bool, error) {
	if b.ReadIdx >= b.BitCount {
		return false, io.EOF
	}

	bit, err := b.ReadBitAt(b.ReadIdx)
	b.ReadIdx++

	// fmt.Printf("! bit: %v\n", bit)
	// defer func (){b.readIdx += 1}() // fancy ha ha
	return bit, err
}


