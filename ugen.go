package sc

import (
	"encoding/binary"
	"io"
)

// UGen done actions, see http://doc.sccode.org/Reference/UGen-doneActions.html
const (
	DoNothing = iota
	Pause
	FreeEnclosing
	FreePreceding
	FreeFollowing
	FreePrecedingGroup
	FreeFollowingGroup
	FreeAllPreceding
	FreeAllFollowing
	FreeAndPausePreceding
	FreeAndPauseFollowing
	DeepFreePreceding
	DeepFreeFollowing
	FreeAllInGroup
	// I do not understand the difference between the last and
	// next-to-last options [bps]
)

// Ugen is a unit generator.
type Ugen struct {
	Name         string      `json:"name"              xml:"name,attr"`
	Rate         int8        `json:"rate"              xml:"rate,attr"`
	SpecialIndex int16       `json:"specialIndex"      xml:"specialIndex,attr"`
	Inputs       []UgenInput `json:"inputs,omitempty"  xml:"Inputs>Input"`
	Outputs      []Output    `json:"outputs,omitempty" xml:"Outputs>Output"`
	NumOutputs   int         `json:"-"                 xml:"-"`

	inputs []Input
}

// NewUgen is a factory function for creating new Ugen instances.
// Panics if rate is not AR, KR, or IR.
// Panics if numOutputs <= 0.
func NewUgen(name string, rate int8, specialIndex int16, numOutputs int, inputs ...Input) *Ugen {
	CheckRate(rate)

	if numOutputs <= 0 {
		panic("numOutputs must be a positive int")
	}
	// TODO: validate specialIndex
	u := &Ugen{
		Name:         name,
		Rate:         rate,
		SpecialIndex: specialIndex,
		NumOutputs:   numOutputs,
		inputs:       inputs,
	}
	// If any inputs are multi inputs, then this node should get promoted to a multi node.
	for i, input := range inputs {
		switch v := input.(type) {
		case *Ugen:
			// Initialize the Outputs slice.
			input = asOutput(v)
		case MultiInput:
			ins := make(Inputs, len(v.InputArray()))

			// Add outputs to any nodes in a MultiInput.
			for i, in := range v.InputArray() {
				if u, ok := in.(*Ugen); ok {
					ins[i] = asOutput(u)
				}
			}
		}
		u.inputs[i] = input
	}
	return u
}

// Add adds an input to a ugen node.
func (u *Ugen) Add(val Input) Input {
	return binOpAdd(u.Rate, u, val, u.NumOutputs)
}

// Max computes the maximum of one Input and another.
func (u *Ugen) Max(other Input) Input {
	return binOpMax(u.Rate, u, other, u.NumOutputs)
}

// Midicps converts MIDI note number to cycles per second.
func (u *Ugen) Midicps() Input {
	return unaryOpMidicps(u.Rate, u, u.NumOutputs)
}

// Mul multiplies the ugen node by an input.
func (u *Ugen) Mul(val Input) Input {
	return binOpMul(u.Rate, u, val, u.NumOutputs)
}

// MulAdd multiplies and adds inputs to a ugen node.
func (u *Ugen) MulAdd(mul, add Input) Input {
	return mulAdd(u.Rate, u, mul, add, u.NumOutputs)
}

// Neg is a convenience operator that multiplies a signal by -1.
func (u *Ugen) Neg() Input {
	return unaryOpNeg(u.Rate, u, u.NumOutputs)
}

// SoftClip adds distortion to a ugen.
func (u *Ugen) SoftClip() Input {
	return unaryOpSoftClip(u.Rate, u, u.NumOutputs)
}

// Write writes a Ugen
func (u *Ugen) Write(w io.Writer) error {
	// write the synthdef name
	if err := newPstring(u.Name).Write(w); err != nil {
		return err
	}
	// write rate
	if err := binary.Write(w, byteOrder, u.Rate); err != nil {
		return err
	}
	// write inputs
	numInputs := int32(len(u.Inputs))
	if err := binary.Write(w, byteOrder, numInputs); err != nil {
		return err
	}
	// write outputs
	numOutputs := int32(len(u.Outputs))
	if err := binary.Write(w, byteOrder, numOutputs); err != nil {
		return err
	}
	// special index
	if err := binary.Write(w, byteOrder, u.SpecialIndex); err != nil {
		return err
	}
	// inputs
	for _, i := range u.Inputs {
		if err := i.Write(w); err != nil {
			return err
		}
	}
	// outputs
	for _, o := range u.Outputs {
		if err := o.Write(w); err != nil {
			return err
		}
	}
	return nil
}

// readUgen reads a ugen from an io.Reader
func readUgen(r io.Reader) (*Ugen, error) {
	var (
		numInputs    int32
		numOutputs   int32
		specialIndex int16
		rate         int8
	)
	// read name
	name, err := readPstring(r)
	if err != nil {
		return nil, err
	}
	// read calculation rate
	if err := binary.Read(r, byteOrder, &rate); err != nil {
		return nil, err
	}
	// read number of inputs
	if err := binary.Read(r, byteOrder, &numInputs); err != nil {
		return nil, err
	}
	// read number of outputs
	if err := binary.Read(r, byteOrder, &numOutputs); err != nil {
		return nil, err
	}
	// read special index
	if err := binary.Read(r, byteOrder, &specialIndex); err != nil {
		return nil, err
	}
	var (
		inputs  = make([]UgenInput, numInputs)
		outputs = make([]Output, numOutputs)
	)
	// read inputs
	for i := 0; int32(i) < numInputs; i++ {
		in, err := readInput(r)
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
	return &Ugen{
		Name:         name.String(),
		Rate:         rate,
		SpecialIndex: specialIndex,
		Inputs:       inputs,
		Outputs:      outputs,
	}, nil
}

// asOutput initializes the outputs array of the ugen node.
func asOutput(u *Ugen) *Ugen {
	if u.Outputs == nil {
		u.Outputs = make([]Output, u.NumOutputs)
		for i := range u.Outputs {
			u.Outputs[i] = Output(u.Rate)
		}
	}
	return u
}

func cloneUgen(v *Ugen) *Ugen {
	u := *v
	return &u
}
