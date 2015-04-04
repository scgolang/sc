package sc

import (
	"encoding/binary"
	"io"
)

// output structure of ugen output
type output struct {
	Rate int8 `json:"rate" xml:"rate,attr"`
}

// Write writes this output to an io.Writer
func (self *output) Write(w io.Writer) error {
	return binary.Write(w, byteOrder, self.Rate)
}

func readoutput(r io.Reader) (*output, error) {
	var rate int8
	err := binary.Read(r, byteOrder, &rate)
	if err != nil {
		return nil, err
	}
	out := output{rate}
	return &out, nil
}
