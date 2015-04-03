package sc

import (
	"encoding/binary"
	"io"
)

type input struct {
	UgenIndex   int32 `json:"ugenIndex"`
	OutputIndex int32 `json:"outputIndex"`
}

// Write writes an input to an io.Writer
func (self *input) Write(w io.Writer) error {
	if we := binary.Write(w, byteOrder, self.UgenIndex); we != nil {
		return we
	}
	return binary.Write(w, byteOrder, self.OutputIndex)
}

func readinput(r io.Reader) (*input, error) {
	var ugenIndex, outputIndex int32
	err := binary.Read(r, byteOrder, &ugenIndex)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, byteOrder, &outputIndex)
	if err != nil {
		return nil, err
	}
	is := input{ugenIndex, outputIndex}
	return &is, nil
}

func newInput(ugenIndex, outputIndex int32) *input {
	return &input{ugenIndex, outputIndex}
}
