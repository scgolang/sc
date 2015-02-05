package sc

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Input
type Input struct {
	ugenIndex   int32
	outputIndex int32
}

func (self *Input) UgenIndex() int32 {
	return self.ugenIndex
}

func (self *Input) OutputIndex() int32 {
	return self.outputIndex
}

func (self *Input) Dump(w io.Writer) error {
	fmt.Fprintf(w, "%-30s %d\n", "UgenIndex", self.ugenIndex)
	fmt.Fprintf(w, "%-30s %d\n", "OutputIndex", self.outputIndex)
	return nil
}

// Write writes an inputSpec to an io.Writer
func (self *Input) Write(w io.Writer) error {
	if we := binary.Write(w, byteOrder, self.ugenIndex); we != nil {
		return we
	}
	return binary.Write(w, byteOrder, self.outputIndex)
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

