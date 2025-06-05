package bitstream

type Reader interface {
	Read() (bool, error)
}

