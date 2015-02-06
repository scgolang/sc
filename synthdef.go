package sc

import (
	"io"
)

// Synthdef
type Synthdef interface {
	// Name returns the name of the synthdef.
	Name() string

	// AppendConstant appends a float to the synthdef's
	// list of constants only if it isn't there already.
	AppendConstant(value float32)

	// AppendUgen appends a Ugen to the synthdef's
	// list of Ugen only if it isn't there already.
	AppendUgen(value *Ugen)

	// Rep returns a structure that can be serialized
	// to the synthdef representation that scsynth supports.
	Rep() *SynthdefRep

	// Dump writes human-readable information about a synthdef
	// to an io.Writer
	Dump(w io.Writer) error
}

type synthdef struct {
	name      string
	constants []float32
	ugens     []*Ugen
}

func (self *synthdef) Name() string {
	return self.name
}

func (self *synthdef) AppendConstant(value float32) {
	for _, c := range self.constants {
		if value == c {
			return
		}
	}
	self.constants = append(self.constants, value)
}

func (self *synthdef) AppendUgen(value *Ugen) {
	for _, u := range self.ugens {
		if value == u {
			return
		}
	}
	self.ugens = append(self.ugens, value)
}

func (self *synthdef) Rep() *SynthdefRep {
	rep := SynthdefRep{
		self.name,
		self.constants,
		make([]float32, 0),   // initialParamValues
		make([]ParamName, 0), // paramNames
		make([]*UgenRep, 0),  // ugens
		make([]Variant, 0),   // variants
	}

	for i, u := range self.ugens {
		ugenRep := self.ugenRep(i, u)
		rep.Ugens = append(rep.Ugens, ugenRep)
	}

	return &rep
}

// ugenRep converts the indexth ugen to a ugen rep
func (self *synthdef) ugenRep(index int, ugen *Ugen) *UgenRep {
	// note that if index != 0 then the ugen will have outputs
	// this is because self.ugens[0] is the root of the ugen tree

	// convert u to a UgenRep and append it
	// to rep.Ugens
	ugenRep := UgenRep{
		ugen.name,
		ugen.rate,
		int16(0), // BUG(briansorahan) where does special index come from?
		make([]*InputRep, 0),
		// I believe a ugen has an output when it is used as an input
		// to another ugen, and that this is how we should populate
		// the UgenRep.Outputs: by looking for other ugens that are using
		// this one (u) as an input.
		// [briansorahan]
		make([]*OutputRep, 0),
	}
	// populate the UgenRep's Inputs array
	for _, in := range ugen.inputs {
		// note that outputIndex is actually "index of constant"
		// if ugenIndex == -1 (i.e. if this input is a constant)
		// otherwise it determines the index of the unit generator's
		// output (as determined by ugenIndex) that is connected to
		// this input
		var ugenIndex, outputIndex int32

		if in.IsConstant() {
			ugenIndex = -1
			// find the constant whose value equals this input
			constantValue := in.ConstantValue()
			for i, c := range self.constants {
				if c == constantValue {
					outputIndex = int32(i)
				}
			}
		} else {
			// find the ugen that equals the value of this input
			ugenInput := in.UgenValue()
			for i, ui := range self.ugens {
				if ui == ugenInput {
					ugenIndex = int32(i)
				}
			}
		}

		// FIXME
		inputRep := InputRep{ugenIndex, outputIndex}
		ugenRep.Inputs = append(ugenRep.Inputs, &inputRep)
	}

	return nil
}

func (self *synthdef) Dump(w io.Writer) error {
	return nil
}

// NewSynthdef creates a new Synthdef from a UgenGraphFunc
func NewSynthdef(name string, f UgenGraphFunc) Synthdef {
	// TODO: pass an interface into the ugen graph func that allows
	//       users to define params for the synthdef
	rootUgen, err := f()
	if err != nil {
		panic(err)
	}

	s := synthdef{
		name,               // name
		make([]float32, 0), // constants
		make([]*Ugen, 0),   // ugens
	}

	flatten(s, rootUgen)

	return &s
}

// flatten the ugen graph into a list of constants
// and a list of ugens
func flatten(s synthdef, root *Ugen) {
	numInputs := len(root.inputs)

	for i := numInputs - 1; i >= 0; i-- {
		input := root.inputs[i]

		if input.IsConstant() {
			s.AppendConstant(input.ConstantValue())
		} else {
			if i > 0 {
				flatten(s, input.UgenValue())
			} else {
				s.AppendUgen(input.UgenValue())
			}
		}
	}
}
