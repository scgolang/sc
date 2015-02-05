package sc

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Input
type Input struct {
	UgenIndex   int32 `json:'ugenIndex,omitempty'`
	OutputIndex int32 `json:'outputIndex,omitempty'`
}

func (self *Input) Dump(w io.Writer) error {
	fmt.Fprintf(w, "%-30s %d\n", "UgenIndex", self.UgenIndex)
	fmt.Fprintf(w, "%-30s %d\n", "OutputIndex", self.OutputIndex)
	return nil
}

// Write writes an inputSpec to an io.Writer
func (self *Input) Write(w io.Writer) error {
	if we := binary.Write(w, byteOrder, self.UgenIndex); we != nil {
		return we
	}
	return binary.Write(w, byteOrder, self.OutputIndex)
}

func readInput(r io.Reader) (*Input, error) {
	var ugenIndex, outputIndex int32
	err := binary.Read(r, byteOrder, &ugenIndex)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, byteOrder, &outputIndex)
	if err != nil {
		return nil, err
	}
	is := Input{ugenIndex, outputIndex}
	return &is, nil
}

