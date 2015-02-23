package sc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	. "github.com/briansorahan/sc/types"
	"io"
)

// ugen
type ugen struct {
	Name         string    `json:"name" xml:"name,attr"`
	Rate         int8      `json:"rate" xml:"rate,attr"`
	SpecialIndex int16     `json:"specialIndex" xml:"specialIndex,attr"`
	Inputs       []*input  `json:"inputs" xml:"Inputs>Input"`
	Outputs      []*output `json:"outputs" xml:"Outputs>Output"`
}

func (self *ugen) AppendInput(i *input) {
	self.Inputs = append(self.Inputs, i)
}

// AddOutput ensures that a ugen has an output at self.Rate
// How do you create a ugen with multiple outputs? -bps
func (self *ugen) AddOutput(o *output) {
	self.Outputs = append(self.Outputs, o)
}

// Write writes a Ugen
func (self *ugen) Write(w io.Writer) error {
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

// readugen reads a ugen from an io.Reader
func readugen(r io.Reader) (*ugen, error) {
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
	inputs := make([]*input, numInputs)
	for i := 0; int32(i) < numInputs; i++ {
		in, err := readinput(r)
		if err != nil {
			return nil, err
		}
		inputs[i] = in
	}
	// read outputs
	outputs := make([]*output, numOutputs)
	for i := 0; int32(i) < numOutputs; i++ {
		out, err := readoutput(r)
		if err != nil {
			return nil, err
		}
		outputs[i] = out
	}

	u := ugen{
		name.String(),
		rate,
		specialIndex,
		inputs,
		outputs,
	}
	return &u, nil
}

func newUgen(name string, rate int8) *ugen {
	u := ugen{
		name,
		rate,
		0, // special index
		make([]*input, 0),
		make([]*output, 0),
	}
	return &u
}

func cloneUgen(n UgenNode) *ugen {
	u := ugen{
		n.Name(),
		n.Rate(),
		n.SpecialIndex(),
		// inputs get added at synthdef-creation time
		make([]*input, 0),
		make([]*output, 0),
	}
	// add outputs
	for _, out := range n.Outputs() {
		u.AddOutput(&output{out.Rate()})
	}
	return &u
}
