package sc

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Output ugen output
type Output struct {
	rate int8
}

// Rate gets the rate of an output
func (self *Output) Rate() int8 {
	return self.rate
}

// Dump writes information about this output to an io.Writer
func (self *Output) Dump(w io.Writer) error {
	fmt.Fprintf(w, "%-30s %d\n", "Rate", self.rate)
	return nil
}

// Write writes this output to an io.Writer
func (self *Output) Write(w io.Writer) error {
	return binary.Write(w, byteOrder, self.rate)
}

func readOutput(r io.Reader) (*Output, error) {
	var rate int8
	err := binary.Read(r, byteOrder, &rate)
	if err != nil {
		return nil, err
	}
	out := Output{rate}
	return &out, nil
}

func newOutput(rate int8) *Output {
	return &Output{rate}
}
