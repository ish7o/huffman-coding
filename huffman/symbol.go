package huffman

import (
	"strconv"
	"strings"
)

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
	var b strings.Builder
	b.WriteByte('\'')
	for _, r := range c.Value {
		// so that \n is a \n not an actual linebreak
		s := strconv.QuoteRuneToGraphic(r)
		b.WriteString(s[1:len(s)-1])
	}
	b.WriteByte('\'')
	return b.String()
}
