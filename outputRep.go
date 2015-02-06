package sc

import (
	"encoding/binary"
	"io"
)

// OutputRep structure of ugen output
type OutputRep struct {
	Rate int8 `json:"rate"`
}

// Write writes this output to an io.Writer
func (self *OutputRep) Write(w io.Writer) error {
	return binary.Write(w, byteOrder, self.Rate)
}

func readOutputRep(r io.Reader) (*OutputRep, error) {
	var rate int8
	err := binary.Read(r, byteOrder, &rate)
	if err != nil {
		return nil, err
	}
	out := OutputRep{rate}
	return &out, nil
}
