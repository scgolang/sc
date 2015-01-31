package gosc

import (
	"encoding/binary"
	"io"
)

type Ugen struct {
	Name Pstring
	Rate int8
	NumInputs int32
	NumOutputs int32
	SpecialIndex int16
	Inputs []InputSpec
	Outputs []OutputSpec
}

type InputSpec struct {
	UgenIndex int32
	OutputIndex int32
}

// write an input
func (self *InputSpec) Write(w io.Writer) error {
	if we := binary.Write(w, byteOrder, self.UgenIndex); we != nil {
		return we
	}
	return binary.Write(w, byteOrder, self.OutputIndex)
}

type OutputSpec struct {
	Rate int8
}

// write an output
func (self *OutputSpec) Write(w io.Writer) error {
	return binary.Write(w, byteOrder, self.Rate)
}

// write a ugen
func (self *Ugen) Write(w io.Writer) error {
	// write the synthdef name
	we := self.Name.Write(w)
	if we != nil {
		return we
	}
	// audio rate
	we = binary.Write(w, byteOrder, self.Rate)
	if we != nil {
		return we
	}
	// one input
	we = binary.Write(w, byteOrder, self.NumInputs)
	if we != nil {
		return we
	}
	// one output
	we = binary.Write(w, byteOrder, self.NumOutputs)
	if we != nil {
		return we
	}
	// special index
	we = binary.Write(w, byteOrder, self.SpecialIndex)
	if we != nil {
		return we
	}
	// inputs
	for _, i := range self.Inputs {
		if we = i.Write(w); we != nil {
			return we
		}
	}
	// outputs
	for _, o := range self.Outputs {
		if we = o.Write(w); we != nil {
			return we
		}
	}
	return nil
}
