package sc

import (
	"fmt"
)

// ArraySpec is a convenience type for Klang and Klank.
type ArraySpec [3][]Input

// Abs computes the absolute value of a signal.
func (as ArraySpec) Abs() Input {
	var nas ArraySpec
	for i := range []int{0, 1, 2} {
		for j := range nas[i] {
			nas[i][j] = as[i][j].Abs()
		}
	}
	return nas
}

// Add adds one input to another.
func (as ArraySpec) Add(val Input) Input {
	var nas ArraySpec
	for i := range []int{0, 1, 2} {
		for j := range nas[i] {
			nas[i][j] = as[i][j].Add(val)
		}
	}
	return nas
}

// Ceil computes the ceiling of a signal.
func (as ArraySpec) Ceil() Input {
	var nas ArraySpec
	for i := range []int{0, 1, 2} {
		for j := range nas[i] {
			nas[i][j] = as[i][j].Ceil()
		}
	}
	return nas
}

// Cubed computes the cube of a signal.
func (as ArraySpec) Cubed() Input {
	var nas ArraySpec
	for i := range []int{0, 1, 2} {
		for j := range nas[i] {
			nas[i][j] = as[i][j].Cubed()
		}
	}
	return nas
}

// Floor computes the floor of a signal.
func (as ArraySpec) Floor() Input {
	var nas ArraySpec
	for i := range []int{0, 1, 2} {
		for j := range nas[i] {
			nas[i][j] = as[i][j].Floor()
		}
	}
	return nas
}

// Max returns the maximum of one input and another.
func (as ArraySpec) Max(val Input) Input {
	var nas ArraySpec
	for i := range []int{0, 1, 2} {
		for j := range nas[i] {
			nas[i][j] = as[i][j].Max(val)
		}
	}
	return nas
}

// Midicps converts from MIDI note values to cycles per second.
func (as ArraySpec) Midicps() Input {
	var nas ArraySpec
	for i := range []int{0, 1, 2} {
		for j := range nas[i] {
			nas[i][j] = as[i][j].Midicps()
		}
	}
	return nas
}

// Mul multiplies one input and another.
func (as ArraySpec) Mul(val Input) Input {
	var nas ArraySpec
	for i := range []int{0, 1, 2} {
		for j := range nas[i] {
			nas[i][j] = as[i][j].Mul(val)
		}
	}
	return nas
}

// MulAdd computes (as * m) + a.
func (as ArraySpec) MulAdd(m, a Input) Input {
	var nas ArraySpec
	for i := range []int{0, 1, 2} {
		for j := range nas[i] {
			nas[i][j] = as[i][j].MulAdd(m, a)
		}
	}
	return nas
}

// Neg negates an input.
func (as ArraySpec) Neg() Input {
	var nas ArraySpec
	for i := range []int{0, 1, 2} {
		for j := range nas[i] {
			nas[i][j] = as[i][j].Neg()
		}
	}
	return nas
}

// Reciprocal computes the reciprocal of a signal.
func (as ArraySpec) Reciprocal() Input {
	var nas ArraySpec
	for i := range []int{0, 1, 2} {
		for j := range nas[i] {
			nas[i][j] = as[i][j].Reciprocal()
		}
	}
	return nas
}

// SoftClip computes nonlinear distortion of an input.
func (as ArraySpec) SoftClip() Input {
	var nas ArraySpec
	for i := range []int{0, 1, 2} {
		for j := range nas[i] {
			nas[i][j] = as[i][j].SoftClip()
		}
	}
	return nas
}

// Squared computes the square of a signal.
func (as ArraySpec) Squared() Input {
	var nas ArraySpec
	for i := range []int{0, 1, 2} {
		for j := range nas[i] {
			nas[i][j] = as[i][j].Squared()
		}
	}
	return nas
}

func (as ArraySpec) inputs(freqfirst bool) []Input {
	var ins []Input

	for i, freq := range as[0] {
		if freqfirst {
			ins = append(ins, freq)
		}
		if i >= len(as[1]) {
			ins = append(ins, C(1))
		} else {
			ins = append(ins, as[1][i])
		}
		if i >= len(as[2]) {
			ins = append(ins, C(0))
		} else {
			ins = append(ins, as[2][i])
		}
		if !freqfirst {
			ins = append(ins, freq)
		}
	}
	return ins
}

func (as ArraySpec) normalize() ArraySpec {
	nas := ArraySpec{as[0], as[1], as[2]}

	if as[1] == nil {
		nas[1] = make([]Input, len(as[0]))
	}
	if as[2] == nil {
		nas[2] = make([]Input, len(as[0]))
	}
	for i := range nas[1] {
		if nas[1][i] == nil {
			nas[1][i] = C(1)
		}
		if nas[2][i] == nil {
			nas[2][i] = C(0)
		}
	}
	return nas
}

func getArraySpecInputs(in Input) []ArraySpec {
	var specs []ArraySpec
	switch v := in.(type) {
	default:
		panic(fmt.Sprintf("unexpected Spec type: %T", in))
	case ArraySpec:
		specs = append(specs, v.normalize())
	case Inputs:
		for _, in := range v {
			switch x := in.(type) {
			default:
				panic(fmt.Sprintf("unexpected Spec type in multichannel expansion: %T", in))
			case ArraySpec:
				specs = append(specs, x.normalize())
			}
		}
	}
	return specs
}
