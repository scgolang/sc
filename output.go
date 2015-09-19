package sc

import (
	"encoding/binary"
	"io"
)

// Output is a ugen output.
type Output int8

// Rate returns the rate of the output.
func (o Output) Rate() int8 {
	return int8(o)
}

// Write writes the output to an io.Writer.
func (o Output) Write(w io.Writer) error {
	return binary.Write(w, byteOrder, int8(o))
}

// readOutput
func readOutput(r io.Reader) (Output, error) {
	var (
		rate int8
	)
	if err := binary.Read(r, byteOrder, &rate); err != nil {
		return Output(-1), err
	}
	return Output(rate), nil
}
