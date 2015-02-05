package sc

import (
	"encoding/binary"
	"io"
)

// Output ugen output
type Output struct {
	Rate int8 `json:"rate,omitempty"`
}

// Write writes this output to an io.Writer
func (self *Output) Write(w io.Writer) error {
	return binary.Write(w, byteOrder, self.Rate)
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
