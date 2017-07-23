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

// Abs computes the absolute value of a signal.
func (ins Inputs) Abs() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Abs()
	}
	return Inputs(ia)
}

// Add adds an input to all the inputs.
func (ins Inputs) Add(val Input) Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Add(val)
	}
	return Inputs(ia)
}

// AmpDb converts linear amplitude to decibels.
func (ins Inputs) AmpDb() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.AmpDb()
	}
	return Inputs(ia)
}

// Bilinrand returns a linearly distributed random value between [+in ... -in].
func (ins Inputs) Bilinrand() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Bilinrand()
	}
	return Inputs(ia)
}

// Ceil computes the ceiling (next highest integer) of a signal.
func (ins Inputs) Ceil() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Ceil()
	}
	return Inputs(ia)
}

// Cpsmidi converts frequency in Hz to midi note values.
func (ins Inputs) Cpsmidi() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Cpsmidi()
	}
	return Inputs(ia)
}

// Cpsoct converts cycles per second to decimal octaves.
func (ins Inputs) Cpsoct() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Cpsoct()
	}
	return Inputs(ia)
}

// Cubed computes the cube of a signal.
func (ins Inputs) Cubed() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Cubed()
	}
	return Inputs(ia)
}

// DbAmp converts decibels to linear amplitude.
func (ins Inputs) DbAmp() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.DbAmp()
	}
	return Inputs(ia)
}

// Exp computes the exponential of a signal.
func (ins Inputs) Exp() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Exp()
	}
	return Inputs(ia)
}

// Floor computes the floor (next lowest integer) of a signal.
func (ins Inputs) Floor() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Floor()
	}
	return Inputs(ia)
}

// Frac computes the fractional part of a signal.
func (ins Inputs) Frac() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Frac()
	}
	return Inputs(ia)
}

// InputArray provides access to the list of inputs.
func (ins Inputs) InputArray() []Input {
	return ins
}

// Linrand returns a linearly distributed random value between in and zero.
func (ins Inputs) Linrand() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Linrand()
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

// Midicps converts MIDI note number to cycles per second.
func (ins Inputs) Midicps() Input {
	converted := make([]Input, len(ins))
	for i, in := range ins {
		converted[i] = in.Midicps()
	}
	return Inputs(converted)
}

// Midiratio converts an interval in MIDI notes into a frequency ratio.
func (ins Inputs) Midiratio() Input {
	converted := make([]Input, len(ins))
	for i, in := range ins {
		converted[i] = in.Midiratio()
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

// Octcps converts decimal octaves to cycles per second.
func (ins Inputs) Octcps() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Octcps()
	}
	return Inputs(ia)
}

// Rand returns an evenly distributed random value between this and zero.
func (ins Inputs) Rand() Input {
	converted := make([]Input, len(ins))
	for i, in := range ins {
		converted[i] = in.Rand()
	}
	return Inputs(converted)
}

// Rand2 returns an evenly distributed random value between [+this ... - this].
func (ins Inputs) Rand2() Input {
	converted := make([]Input, len(ins))
	for i, in := range ins {
		converted[i] = in.Rand2()
	}
	return Inputs(converted)
}

// Ratiomidi converts a frequency ratio to an interval in MIDI notes.
func (ins Inputs) Ratiomidi() Input {
	converted := make([]Input, len(ins))
	for i, in := range ins {
		converted[i] = in.Ratiomidi()
	}
	return Inputs(converted)
}

// Reciprocal computes the reciprocal of a signal.
func (ins Inputs) Reciprocal() Input {
	converted := make([]Input, len(ins))
	for i, in := range ins {
		converted[i] = in.Reciprocal()
	}
	return Inputs(converted)
}

// Sign computes the sign of a signal.
func (ins Inputs) Sign() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Sign()
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

// Sqrt computes the square root of a signal.
func (ins Inputs) Sqrt() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Sqrt()
	}
	return Inputs(ia)
}

// Squared computes the square of a signal.
func (ins Inputs) Squared() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Squared()
	}
	return Inputs(ia)
}

// Multi does multichannel expansion.
// See http://doc.sccode.org/Guides/Multichannel-Expansion.html.
func Multi(inputs ...Input) Inputs {
	return Inputs(inputs)
}
