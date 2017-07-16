package sc

// MultiInput is the interface of an input that causes
// cascading multi-channel expansion.
// See http://doc.sccode.org/Guides/Multichannel-Expansion.html
type MultiInput interface {
	Input
	InputArray() []Input
}

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

// InputArray provides access to the list of inputs.
func (ins Inputs) InputArray() []Input {
	return ins
}

// Max returns Inputs that contain the max of all the inputs and the provided Input.
func (ins Inputs) Max(other Input) Input {
	im := make([]Input, len(ins))
	for i, in := range ins {
		im[i] = in.Max(other)
	}
	return Inputs(im)
}

// Midicps converts MIDI note number to cycles per second.
func (ins Inputs) Midicps() Input {
	converted := make([]Input, len(ins))
	for i, in := range ins {
		converted[i] = in.Midicps()
	}
	return Inputs(converted)
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

// Neg is a convenience operator that multiplies a signal by -1.
func (ins Inputs) Neg() Input {
	converted := make([]Input, len(ins))
	for i, in := range ins {
		converted[i] = in.Neg()
	}
	return Inputs(converted)
}

// SoftClip adds distortion to the inputs.
func (ins Inputs) SoftClip() Input {
	clipped := make([]Input, len(ins))
	for i, in := range ins {
		clipped[i] = in.SoftClip()
	}
	return Inputs(clipped)
}

// Multi does multichannel expansion.
// See http://doc.sccode.org/Guides/Multichannel-Expansion.html.
func Multi(inputs ...Input) Inputs {
	return Inputs(inputs)
}
