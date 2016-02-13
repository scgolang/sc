package sc

import (
	"encoding/binary"
	"io"
)

// Input is implemented by any value that can serve as a
// ugen input. This includes synthdef parameters,
// constants, and other ugens.
type Input interface {
	Mul(val Input) Input
	Add(val Input) Input
	MulAdd(mul, add Input) Input
}

// MultiInput is the interface of an input that causes
// cascading multi-channel expansion.
// See http://doc.sccode.org/Guides/Multichannel-Expansion.html
type MultiInput interface {
	Input
	InputArray() []Input
}

type input struct {
	UgenIndex   int32 `json:"ugenIndex" xml:"ugenIndex,attr"`
	OutputIndex int32 `json:"outputIndex" xml:"outputIndex,attr"`
}

// Write writes an input to an io.Writer
func (input *input) Write(w io.Writer) error {
	if we := binary.Write(w, byteOrder, input.UgenIndex); we != nil {
		return we
	}
	return binary.Write(w, byteOrder, input.OutputIndex)
}

func readinput(r io.Reader) (input, error) {
	var ugenIndex, outputIndex int32
	err := binary.Read(r, byteOrder, &ugenIndex)
	if err != nil {
		return input{}, err
	}
	err = binary.Read(r, byteOrder, &outputIndex)
	if err != nil {
		return input{}, err
	}
	return input{ugenIndex, outputIndex}, nil
}

// newInput
func newInput(ugenIndex, outputIndex int32) input {
	return input{ugenIndex, outputIndex}
}
