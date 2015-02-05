package sc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// Ugen
type Ugen struct {
	Name         string   `json:"name,omitempty"`
	Rate         int8     `json:"rate,omitempty"`
	NumInputs    int32    `json:"numInputs,omitempty"`
	NumOutputs   int32    `json:"numOutputs,omitempty"`
	SpecialIndex int16    `json:"specialIndex,omitempty"`
	Inputs       []Input  `json:"inputs,omitempty"`
	Outputs      []Output `json:"outputs,omitempty"`
}

func (self *Ugen) AddConstant(value float32) {
}

func (self *Ugen) AddUgen(value Ugen) {
}

// write a Ugen
func (self *Ugen) Write(w io.Writer) error {
	// write the synthdef name
	nameLen := len(self.Name)
	bw, we := w.Write(bytes.NewBufferString(self.Name).Bytes())
	if we != nil {
		return we
	}
	if bw != nameLen {
		return fmt.Errorf("could not write Ugen.Name")
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
		name.String(),
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
		name,              // name
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
		} else if ug, isUgen := arg.(Ugen); isUgen {
			u.AddUgen(ug)
		} else {
			return nil, fmt.Errorf("ugen arguments must be float32's or ugen's")
		}
	}

	return &u, nil
}
