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

func ReadInputSpec(r io.Reader) (*InputSpec, error) {
	var ugenIndex, outputIndex int32
	err := binary.Read(r, byteOrder, &ugenIndex)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, byteOrder, &outputIndex)
	if err != nil {
		return nil, err
	}
	is := InputSpec{ugenIndex, outputIndex}
	return &is, nil
}

type OutputSpec struct {
	Rate int8
}

// write an output
func (self *OutputSpec) Write(w io.Writer) error {
	return binary.Write(w, byteOrder, self.Rate)
}

func ReadOutputSpec(r io.Reader) (*OutputSpec, error) {
	var rate int8
	err := binary.Read(r, byteOrder, &rate)
	if err != nil {
		return nil, err
	}
	outputSpec := OutputSpec{rate}
	return &outputSpec, nil
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

// ReadUgen reads a Ugen from an io.Reader
func ReadUgen(r io.Reader) (*Ugen, error) {
	// read name
	name, err := ReadPstring(r)
	if err != nil {
		return nil, err
	}
	// read calculation rate
	var rate int8
	err = binary.Read(r, byteOrder, &rate)
	if err != nil {
		return nil, err
	}
	// read number of inputs
	var numInputs int32
	err = binary.Read(r, byteOrder, &numInputs)
	if err != nil {
		return nil, err
	}
	// read number of outputs
	var numOutputs int32
	err = binary.Read(r, byteOrder, &numOutputs)
	if err != nil {
		return nil, err
	}
	// read special index
	var specialIndex int16
	err = binary.Read(r, byteOrder, &specialIndex)
	if err != nil {
		return nil, err
	}
	// read inputs
	inputs := make([]InputSpec, numInputs)
	for i := 0; int32(i) < numInputs; i++ {
		inspec, err := ReadInputSpec(r)
		if err != nil {
			return nil, err
		}
		inputs[i] = *inspec
	}
	// read outputs
	outputs := make([]OutputSpec, numOutputs)
	for i := 0; int32(i) < numOutputs; i++ {
		outspec, err := ReadOutputSpec(r)
		if err != nil {
			return nil, err
		}
		outputs[i] = *outspec
	}
	ugen := Ugen{
		*name,
		rate,
		numInputs,
		numOutputs,
		specialIndex,
		inputs,
		outputs,
	}
	return &ugen, nil
}
