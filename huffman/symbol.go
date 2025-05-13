package huffman

import "fmt"

type Symbol struct {
	Value []rune
	Freq  int
}

func (c Symbol) Combine(other Symbol) Symbol {
	return Symbol{
		Value: append(c.Value, other.Value...),
		Freq:  c.Freq + other.Freq,
	}
}

func (c Symbol) String() string {
	return fmt.Sprintf("'%s': %d", string(c.Value), c.Freq)
}
