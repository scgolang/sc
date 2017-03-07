package sc

// Inputs is a slice of Input.
type Inputs []Input

// Add adds an input to all the inputs.
func (ins Inputs) Add(val Input) Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Add(val)
	}
	return Inputs(ia)
}

// Max returns Inputs that contain the max of all the inputs and the provided Input.
func (ins Inputs) Max(other Input) Input {
	im := make([]Input, len(ins))
	for i, in := range ins {
		im[i] = in.Max(other)
	}
	return Inputs(im)
}

// Mul multiplies all the inputs by another input.
func (ins Inputs) Mul(val Input) Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Mul(val)
	}
	return Inputs(ia)
}

// MulAdd performs a multiplication and addition on all the inputs.
func (ins Inputs) MulAdd(mul, add Input) Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.MulAdd(mul, add)
	}
	return Inputs(ia)
}

// SoftClip adds distortion to the inputs.
func (ins Inputs) SoftClip() Input {
	clipped := make([]Input, len(ins))
	for i, in := range ins {
		clipped[i] = in.SoftClip()
	}
	return Inputs(clipped)
}

// InputArray provides access to the list of inputs.
func (ins Inputs) InputArray() []Input {
	return ins
}

// Multi does multichannel expansion.
// See http://doc.sccode.org/Guides/Multichannel-Expansion.html.
func Multi(inputs ...Input) Inputs {
	return Inputs(inputs)
}
