package models

import "fmt"

type HSymbol struct {
	Value []rune
	Freq  int
}

func (c HSymbol) Combine(other HSymbol) HSymbol {
    return HSymbol{
        Value: append(c.Value, other.Value...),
        Freq:  c.Freq + other.Freq,
    }
}

func (c HSymbol) String() string {
	return fmt.Sprintf("'%s': %d", string(c.Value), c.Freq)
}

