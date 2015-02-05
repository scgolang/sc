package sc

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Ugen interface {
	Index() int32
	Name() string
	Rate() int8
	NumInputs() int32
	NumOutputs() int32
	SpecialIndex() int16
	Inputs() []Input
	Outputs() []Output
}

type Input interface {
	UgenIndex() int32
	OutputIndex() int32
	// Write writes a text representation of an Input
	// to an io.Writer
	Dump(w io.Writer) error
	// Write writes a binary representation of an Input
	// to an io.Writer
	Write(w io.Writer) error
}

type Output interface {
	Rate() int8
	// Write writes a text representation of an Output
	// to an io.Writer
	Dump(w io.Writer) error
	// Write writes a binary representation of an Output
	// to an io.Writer
	Write(w io.Writer) error
}

type UgenNode struct {
	name         Pstring
	rate         int8
	numInputs    int32
	numOutputs   int32
	specialIndex int16
	inputs       []Input
	outputs      []Output
}

func (self *UgenNode) AddChild(child Node) {
}

func (self *UgenNode) Name() string {
	return self.name.String
}

func (self *UgenNode) Rate() int8 {
	return self.rate
}

func (self *UgenNode) NumInputs() int32 {
	return self.numInputs
}

func (self *UgenNode) NumOutputs() int32 {
	return self.numOutputs
}

func (self *UgenNode) SpecialIndex() int16 {
	return self.specialIndex
}

func (self *UgenNode) Inputs() []Input {
	return self.inputs
}

func (self *UgenNode) Outputs() []Output {
	return self.outputs
}

func (self *UgenNode) Dump(w io.Writer) error {
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

// write a UgenNode
func (self *UgenNode) Write(w io.Writer) error {
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

// ReadUgen reads a UgenNode from an io.Reader
func ReadUgenNode(r io.Reader) (*UgenNode, error) {
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
		inspec, err := readInputSpec(r)
		if err != nil {
			return nil, err
		}
		inputs[i] = inspec
	}
	// read outputs
	outputs := make([]Output, numOutputs)
	for i := 0; int32(i) < numOutputs; i++ {
		outspec, err := readOutputSpec(r)
		if err != nil {
			return nil, err
		}
		outputs[i] = outspec
	}
	u := UgenNode{
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

func newUgen(name string, args ...interface{}) Ugen {
	return nil
}

type inputSpec struct {
	ugenIndex   int32
	outputIndex int32
}

func (self *inputSpec) UgenIndex() int32 {
	return self.ugenIndex
}

func (self *inputSpec) OutputIndex() int32 {
	return self.outputIndex
}

func (self *inputSpec) Dump(w io.Writer) error {
	fmt.Fprintf(w, "%-30s %d\n", "UgenIndex", self.ugenIndex)
	fmt.Fprintf(w, "%-30s %d\n", "OutputIndex", self.outputIndex)
	return nil
}

// Write writes an inputSpec to an io.Writer
func (self *inputSpec) Write(w io.Writer) error {
	if we := binary.Write(w, byteOrder, self.ugenIndex); we != nil {
		return we
	}
	return binary.Write(w, byteOrder, self.outputIndex)
}

func readInputSpec(r io.Reader) (Input, error) {
	var ugenIndex, outputIndex int32
	err := binary.Read(r, byteOrder, &ugenIndex)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, byteOrder, &outputIndex)
	if err != nil {
		return nil, err
	}
	is := inputSpec{ugenIndex, outputIndex}
	return &is, nil
}

// OutputSpec ugen output
type outputSpec struct {
	rate int8
}

func (self *outputSpec) Rate() int8 {
	return self.rate
}

// Dump writes information about this output to an io.Writer
func (self *outputSpec) Dump(w io.Writer) error {
	fmt.Fprintf(w, "%-30s %d\n", "Rate", self.rate)
	return nil
}

// Write writes this output to an io.Writer
func (self *outputSpec) Write(w io.Writer) error {
	return binary.Write(w, byteOrder, self.rate)
}

func readOutputSpec(r io.Reader) (Output, error) {
	var rate int8
	err := binary.Read(r, byteOrder, &rate)
	if err != nil {
		return nil, err
	}
	outputSpec := outputSpec{rate}
	return &outputSpec, nil
}
