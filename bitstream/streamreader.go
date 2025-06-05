package bitstream

import (
	"io"
)

type StreamReader struct {
	r        io.Reader
	bitsLeft uint8
	curr     byte
}

func NewStreamReader(r io.Reader) *StreamReader {
	return &StreamReader{
		r:        r,
		bitsLeft: 0,
	}
}

func (sr *StreamReader) Read() (bool, error) {
	if sr.bitsLeft == 0 {
		buf := make([]byte, 1)
		n, err := sr.r.Read(buf)
		if err != nil {
			return false, err
		}
		if n != 1 {
			return false, io.ErrUnexpectedEOF // haha
		}
		sr.curr = buf[0]
		sr.bitsLeft = 8
	}

	// 0b1000000
	//   ^
	bit := (sr.curr & (1 << 7)) != 0
	sr.curr <<= 1
	sr.bitsLeft--


	return bit, nil
}
