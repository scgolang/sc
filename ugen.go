package gosc

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Ugen interface {
	Index() int32
	Ar(name string, args ...interface{}) UgenGraph
	Kr(name string, args ...interface{}) UgenGraph
}

type UgenSpec struct {
	Name Pstring
	Rate int8
	NumInputs int32
	NumOutputs int32
	SpecialIndex int16
	Inputs []InputSpec
	Outputs []OutputSpec
}

func (self *UgenSpec) Dump(w io.Writer) error {
	var e error

	fmt.Fprintf(w, "%-30s %s\n", "Name", self.Name.String)
	fmt.Fprintf(w, "%-30s %d\n", "Rate", self.Rate)
	fmt.Fprintf(w, "%-30s %d\n", "NumInputs", self.NumInputs)
	fmt.Fprintf(w, "%-30s %d\n", "NumOutputs", self.NumOutputs)
	fmt.Fprintf(w, "%-30s %d\n", "SpecialIndex", self.SpecialIndex)
	if self.NumInputs > 0 {
		for i := 0; int32(i) < self.NumInputs; i++ {
			fmt.Fprintf(w, "\nInput %d:\n", i)
			e = self.Inputs[i].Dump(w)
			if e != nil {
				return e
			}
		}
	}
	if self.NumOutputs > 0 {
		for i := 0; int32(i) < self.NumOutputs; i++ {
			fmt.Fprintf(w, "\nOutput %d:\n", i)
			e = self.Outputs[i].Dump(w)
			if e != nil {
				return e
			}
		}
	}
	return nil
}

// write a ugen
func (self *UgenSpec) Write(w io.Writer) error {
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
func ReadUgen(r io.Reader) (*UgenSpec, error) {
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
	ugen := UgenSpec{
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

type InputSpec struct {
	UgenIndex int32
	OutputIndex int32
}

func (self *InputSpec) Dump(w io.Writer) error {
	fmt.Fprintf(w, "%-30s %d\n", "UgenIndex", self.UgenIndex)
	fmt.Fprintf(w, "%-30s %d\n", "OutputIndex", self.OutputIndex)
	return nil
}

// Write writes an InputSpec to an io.Writer
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

// OutputSpec ugen output
type OutputSpec struct {
	Rate int8
}

// Dump writes information about this output to an io.Writer
func (self *OutputSpec) Dump(w io.Writer) error {
	fmt.Fprintf(w, "%-30s %d\n", "Rate", self.Rate)
	return nil
}

// Write writes this output to an io.Writer
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
