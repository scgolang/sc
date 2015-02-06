package sc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// UgenRep
type UgenRep struct {
	Name         string       `json:"name"`
	Rate         int8         `json:"rate"`
	SpecialIndex int16        `json:"specialIndex"`
	Inputs       []*InputRep  `json:"inputs"`
	Outputs      []*OutputRep `json:"outputs"`
}

func (self *UgenRep) AppendInput(i *InputRep) {
	self.Inputs = append(self.Inputs, i)
}

func (self *UgenRep) AddOutput(o *OutputRep) {
	self.Outputs = append(self.Outputs, o)
}

// write a Ugen
func (self *UgenRep) Write(w io.Writer) error {
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
	numInputs := int32(len(self.Inputs))
	we = binary.Write(w, byteOrder, numInputs)
	if we != nil {
		return we
	}
	// one output
	numOutputs := int32(len(self.Outputs))
	we = binary.Write(w, byteOrder, numOutputs)
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

// readUgenRep reads a UgenRep from an io.Reader
func readUgenRep(r io.Reader) (*UgenRep, error) {
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
	inputs := make([]*InputRep, numInputs)
	for i := 0; int32(i) < numInputs; i++ {
		in, err := readInputRep(r)
		if err != nil {
			return nil, err
		}
		inputs[i] = in
	}
	// read outputs
	outputs := make([]*OutputRep, numOutputs)
	for i := 0; int32(i) < numOutputs; i++ {
		out, err := readOutputRep(r)
		if err != nil {
			return nil, err
		}
		outputs[i] = out
	}

	u := UgenRep{
		name.String(),
		rate,
		specialIndex,
		inputs,
		outputs,
	}
	return &u, nil
}
