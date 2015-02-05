package sc

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Ugen
type Ugen struct {
	name         Pstring
	rate         int8
	numInputs    int32
	numOutputs   int32
	specialIndex int16
	inputs       []Input
	outputs      []Output
}

func (self *Ugen) AddConstant(value float32) {
}

func (self *Ugen) AddUgen(value Ugen) {
}

func (self *Ugen) Name() string {
	return self.name.String
}

func (self *Ugen) Rate() int8 {
	return self.rate
}

func (self *Ugen) NumInputs() int32 {
	return self.numInputs
}

func (self *Ugen) NumOutputs() int32 {
	return self.numOutputs
}

func (self *Ugen) SpecialIndex() int16 {
	return self.specialIndex
}

func (self *Ugen) Inputs() []Input {
	return self.inputs
}

func (self *Ugen) Outputs() []Output {
	return self.outputs
}

func (self *Ugen) Dump(w io.Writer) error {
	var e error

	fmt.Fprintf(w, "%-30s %s\n", "Name", self.name.String)
	fmt.Fprintf(w, "%-30s %d\n", "Rate", self.rate)
	fmt.Fprintf(w, "%-30s %d\n", "NumInputs", self.numInputs)
	fmt.Fprintf(w, "%-30s %d\n", "NumOutputs", self.numOutputs)
	fmt.Fprintf(w, "%-30s %d\n", "SpecialIndex", self.specialIndex)
	if self.numInputs > 0 {
		for i := 0; int32(i) < self.numInputs; i++ {
			fmt.Fprintf(w, "\nInput %d:\n", i)
			e = self.inputs[i].Dump(w)
			if e != nil {
				return e
			}
		}
	}
	if self.numOutputs > 0 {
		for i := 0; int32(i) < self.numOutputs; i++ {
			fmt.Fprintf(w, "\nOutput %d:\n", i)
			e = self.outputs[i].Dump(w)
			if e != nil {
				return e
			}
		}
	}
	return nil
}

// write a Ugen
func (self *Ugen) Write(w io.Writer) error {
	// write the synthdef name
	we := self.name.Write(w)
	if we != nil {
		return we
	}
	// audio rate
	we = binary.Write(w, byteOrder, self.rate)
	if we != nil {
		return we
	}
	// one input
	we = binary.Write(w, byteOrder, self.numInputs)
	if we != nil {
		return we
	}
	// one output
	we = binary.Write(w, byteOrder, self.numOutputs)
	if we != nil {
		return we
	}
	// special index
	we = binary.Write(w, byteOrder, self.specialIndex)
	if we != nil {
		return we
	}
	// inputs
	for _, i := range self.inputs {
		if we = i.Write(w); we != nil {
			return we
		}
	}
	// outputs
	for _, o := range self.outputs {
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
	inputs := make([]Input, numInputs)
	for i := 0; int32(i) < numInputs; i++ {
		inspec, err := readInput(r)
		if err != nil {
			return nil, err
		}
		inputs[i] = *inspec
	}
	// read outputs
	outputs := make([]Output, numOutputs)
	for i := 0; int32(i) < numOutputs; i++ {
		outspec, err := readOutput(r)
		if err != nil {
			return nil, err
		}
		outputs[i] = *outspec
	}

	u := Ugen{
		*name,
		rate,
		numInputs,
		numOutputs,
		specialIndex,
		inputs,
		outputs,
	}
	return &u, nil
}

func Ar(name string, args ...interface{}) (*Ugen, error) {
	u := Ugen{
		NewPstring(name),  // name
		2,                 // rate
		0,                 // numInputs
		0,                 // numOutputs
		0,                 // specialIndex
		make([]Input, 0),  // inputs
		make([]Output, 0), // inputs
	}

	for _, arg := range args {
		if fv, isFloat := arg.(float32); isFloat {
			u.AddConstant(fv)
		}
		if ug, isUgen := arg.(Ugen); isUgen {
			u.AddUgen(ug)
		}
	}

	return &u, nil
}
