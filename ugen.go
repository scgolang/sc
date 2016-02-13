package sc

import (
	"encoding/binary"
	"io"
)

// Ugen defines the interface between synthdefs and ugens.
type Ugen interface {
	// Name returns the name of the ugen node
	Name() string
	// Rate returns the rate of the ugen node
	Rate() int8
	// SpecialIndex returns the special index of the ugen node
	SpecialIndex() int16
	// Inputs returns the inputs of the ugen node.
	// Inputs can be
	// (1) Constant (float32)
	// (2) Control (synthdef param)
	// (3) Ugen
	Inputs() []Input
	// Outputs returns the outputs of the ugen node
	Outputs() []Output
}

// ugen
type ugen struct {
	Name         string   `json:"name" xml:"name,attr"`
	Rate         int8     `json:"rate" xml:"rate,attr"`
	SpecialIndex int16    `json:"specialIndex" xml:"specialIndex,attr"`
	Inputs       []input  `json:"inputs" xml:"Inputs>Input"`
	Outputs      []Output `json:"outputs" xml:"Outputs>Output"`
}

func (u *ugen) AppendInput(i input) {
	u.Inputs = append(u.Inputs, i)
}

// AddOutput ensures that a ugen has an output at u.Rate
// How do you create a ugen with multiple outputs? -bps
func (u *ugen) AddOutput(o Output) {
	u.Outputs = append(u.Outputs, o)
}

// Write writes a Ugen
func (u *ugen) Write(w io.Writer) error {
	// write the synthdef name
	err := newPstring(u.Name).Write(w)
	if err != nil {
		return err
	}
	// write rate
	if err = binary.Write(w, byteOrder, u.Rate); err != nil {
		return err
	}
	// write inputs
	numInputs := int32(len(u.Inputs))
	if err = binary.Write(w, byteOrder, numInputs); err != nil {
		return err
	}
	// write outputs
	numOutputs := int32(len(u.Outputs))
	if err = binary.Write(w, byteOrder, numOutputs); err != nil {
		return err
	}
	// special index
	if err = binary.Write(w, byteOrder, u.SpecialIndex); err != nil {
		return err
	}
	// inputs
	for _, i := range u.Inputs {
		if err = i.Write(w); err != nil {
			return err
		}
	}
	// outputs
	for _, o := range u.Outputs {
		if err = o.Write(w); err != nil {
			return err
		}
	}
	return nil
}

// readugen reads a ugen from an io.Reader
func readugen(r io.Reader) (*ugen, error) {
	var (
		rate         int8
		numInputs    int32
		numOutputs   int32
		specialIndex int16
	)

	// read name
	name, err := readPstring(r)
	if err != nil {
		return nil, err
	}

	// read calculation rate
	if err = binary.Read(r, byteOrder, &rate); err != nil {
		return nil, err
	}

	// read number of inputs
	if err = binary.Read(r, byteOrder, &numInputs); err != nil {
		return nil, err
	}

	// read number of outputs
	if err = binary.Read(r, byteOrder, &numOutputs); err != nil {
		return nil, err
	}

	// read special index
	if err = binary.Read(r, byteOrder, &specialIndex); err != nil {
		return nil, err
	}

	var (
		inputs  = make([]input, numInputs)
		outputs = make([]Output, numOutputs)
	)

	// read inputs
	for i := 0; int32(i) < numInputs; i++ {
		in, err := readinput(r)
		if err != nil {
			return nil, err
		}
		inputs[i] = in
	}
	// read outputs
	for i := 0; int32(i) < numOutputs; i++ {
		out, err := readOutput(r)
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
		make([]input, 0),
		make([]Output, 0),
	}
	return &u
}

func cloneUgen(n Ugen) *ugen {
	u := ugen{
		n.Name(),
		n.Rate(),
		n.SpecialIndex(),
		// inputs get added at synthdef-creation time
		make([]input, 0),
		make([]Output, 0),
	}
	// add outputs
	for _, out := range n.Outputs() {
		u.AddOutput(out)
	}
	return &u
}
