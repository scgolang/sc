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

// Absdif returns the absolute value of the difference of two inputs.
func (ins Inputs) Absdif(val Input) Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Absdif(val)
	}
	return Inputs(ia)
}

// Acos computes the arccosine of a signal.
func (ins Inputs) Acos() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Acos()
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

// Amclip returns 0 when b <= 0, a*b when b > 0.
func (ins Inputs) Amclip(val Input) Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Amclip(val)
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

// Asin computes the arcsine of a signal.
func (ins Inputs) Asin() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Asin()
	}
	return Inputs(ia)
}

// Atan computes the arctangent of a signal.
func (ins Inputs) Atan() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Atan()
	}
	return Inputs(ia)
}

// Atan2 returns the arctangent of y/x.
func (ins Inputs) Atan2(val Input) Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Atan2(val)
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

// Clip2 clips input wave a to +/- b
func (ins Inputs) Clip2(val Input) Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Clip2(val)
	}
	return Inputs(ia)
}

// Coin returns one or zero with the probability given by the input.
func (ins Inputs) Coin() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Coin()
	}
	return Inputs(ia)
}

// Cos returns the cosine of an Inputs.
func (ins Inputs) Cos() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Cos()
	}
	return Inputs(ia)
}

// Cosh returns the hyperbolic cosine of an Inputs.
func (ins Inputs) Cosh() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Cosh()
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

// Difsqr returns the value of (a*a) - (b*b).
func (ins Inputs) Difsqr(val Input) Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Difsqr(val)
	}
	return Inputs(ia)
}

// Distort performs non-linear distortion on a signal.
func (ins Inputs) Distort() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Distort()
	}
	return Inputs(ia)
}

// Div divides one input by another.
func (ins Inputs) Div(val Input) Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Div(val)
	}
	return Inputs(ia)
}

// Excess returns the difference of the original signal and its clipped form: (a - clip2(a,b)).
func (ins Inputs) Excess(val Input) Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Excess(val)
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

// Expon raises this input to the power of another.
func (ins Inputs) Expon(val Input) Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Expon(val)
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

// Fold2 folds input wave a to +/- b
func (ins Inputs) Fold2(val Input) Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Fold2(val)
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

// GCD computes the gcd of one Input and another.
func (ins Inputs) GCD(val Input) Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.GCD(val)
	}
	return Inputs(ia)
}

// GT computes x > y.
func (ins Inputs) GT(val Input) Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.GT(val)
	}
	return Inputs(ia)
}

// GTE computes x >= y.
func (ins Inputs) GTE(val Input) Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.GTE(val)
	}
	return Inputs(ia)
}

// Hypot returns the square root of the sum of the squares of a and b.
// Or equivalently, the distance from the origin to the point (x, y).
func (ins Inputs) Hypot(val Input) Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Hypot(val)
	}
	return Inputs(ia)
}

// HypotApx returns an approximation of the square root of the sum of the squares of x and y.
func (ins Inputs) HypotApx(val Input) Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.HypotApx(val)
	}
	return Inputs(ia)
}

// InputArray provides access to the list of inputs.
func (ins Inputs) InputArray() []Input {
	return ins
}

// LCM computes the lcm of one Input and another.
func (ins Inputs) LCM(val Input) Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.LCM(val)
	}
	return Inputs(ia)
}

// LT computes x < y.
func (ins Inputs) LT(val Input) Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.LT(val)
	}
	return Inputs(ia)
}

// LTE computes x <= y.
func (ins Inputs) LTE(val Input) Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.LTE(val)
	}
	return Inputs(ia)
}

// Linrand returns a linearly distributed random value between in and zero.
func (ins Inputs) Linrand() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Linrand()
	}
	return Inputs(ia)
}

// Log computes a natural logarithm.
func (ins Inputs) Log() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Log()
	}
	return Inputs(ia)
}

// Log10 computes a base 10 logarithm.
func (ins Inputs) Log10() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Log10()
	}
	return Inputs(ia)
}

// Log2 computes a base 2 logarithm.
func (ins Inputs) Log2() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Log2()
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

// Min returns the minimum of one signal and another.
func (ins Inputs) Min(other Input) Input {
	im := make([]Input, len(ins))
	for i, in := range ins {
		im[i] = in.Min(other)
	}
	return Inputs(im)
}

// Moddif returns the smaller of the great circle distances between the two points.
func (ins Inputs) Moddif(y, mod Input) Input {
	im := make([]Input, len(ins))
	for i, in := range ins {
		im[i] = in.Moddif(y, mod)
	}
	return Inputs(im)
}

// Modulo computes the modulo of one signal and another.
func (ins Inputs) Modulo(val Input) Input {
	converted := make([]Input, len(ins))
	for i, in := range ins {
		converted[i] = in.Modulo(val)
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

// Pow raises this input to the power of another.
func (ins Inputs) Pow(val Input) Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Pow(val)
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

// Ring1 is ring modulation plus first source.
func (ins Inputs) Ring1(val Input) Input {
	converted := make([]Input, len(ins))
	for i, in := range ins {
		converted[i] = in.Ring1(val)
	}
	return Inputs(converted)
}

// Ring2 is ring modulation plus both sources.
func (ins Inputs) Ring2(val Input) Input {
	converted := make([]Input, len(ins))
	for i, in := range ins {
		converted[i] = in.Ring2(val)
	}
	return Inputs(converted)
}

// Ring3 returns the value of (a*a *b)
func (ins Inputs) Ring3(val Input) Input {
	converted := make([]Input, len(ins))
	for i, in := range ins {
		converted[i] = in.Ring3(val)
	}
	return Inputs(converted)
}

// Ring4 returns the value of ((a*a *b) - (a*b*b)).
func (ins Inputs) Ring4(val Input) Input {
	converted := make([]Input, len(ins))
	for i, in := range ins {
		converted[i] = in.Ring4(val)
	}
	return Inputs(converted)
}

// Round performs quantization by rounding. Rounds a to the nearest multiple of b.
func (ins Inputs) Round(val Input) Input {
	converted := make([]Input, len(ins))
	for i, in := range ins {
		converted[i] = in.Round(val)
	}
	return Inputs(converted)
}

// Scaleneg returns a*b when a < 0, otherwise a.
func (ins Inputs) Scaleneg(val Input) Input {
	converted := make([]Input, len(ins))
	for i, in := range ins {
		converted[i] = in.Scaleneg(val)
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

// Sin returns the sine of an Inputs.
func (ins Inputs) Sin() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Sin()
	}
	return Inputs(ia)
}

// Sinh returns the hyperbolic sine of an Inputs.
func (ins Inputs) Sinh() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Sinh()
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

// Sqrdif computes the square of the difference between the two inputs.
func (ins Inputs) Sqrdif(val Input) Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Sqrdif(val)
	}
	return Inputs(ia)
}

// Sqrsum computes the square of the sum of the two inputs.
func (ins Inputs) Sqrsum(val Input) Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Sqrsum(val)
	}
	return Inputs(ia)
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

// Sum3rand returns a value from a gaussian-like random distribution between in and zero.
func (ins Inputs) Sum3rand() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Sum3rand()
	}
	return Inputs(ia)
}

// Sumsqr returns the value of (a*a) + (b*b).
func (ins Inputs) Sumsqr(val Input) Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Sumsqr(val)
	}
	return Inputs(ia)
}

// Tan returns the tangent of an Inputs.
func (ins Inputs) Tan() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Tan()
	}
	return Inputs(ia)
}

// Tanh returns the hyperbolic tangent of an Inputs.
func (ins Inputs) Tanh() Input {
	ia := make([]Input, len(ins))
	for i, in := range ins {
		ia[i] = in.Tanh()
	}
	return Inputs(ia)
}

// Thresh returns 0 when a < b, otherwise a.
func (ins Inputs) Thresh(val Input) Input {
	converted := make([]Input, len(ins))
	for i, in := range ins {
		converted[i] = in.Thresh(val)
	}
	return Inputs(converted)
}

// Trunc performs quantization by truncation. Truncate a to a multiple of b.
func (ins Inputs) Trunc(val Input) Input {
	converted := make([]Input, len(ins))
	for i, in := range ins {
		converted[i] = in.Trunc(val)
	}
	return Inputs(converted)
}

// Wrap2 wraps input wave to +/-b
func (ins Inputs) Wrap2(val Input) Input {
	converted := make([]Input, len(ins))
	for i, in := range ins {
		converted[i] = in.Wrap2(val)
	}
	return Inputs(converted)
}

// Multi does multichannel expansion.
// See http://doc.sccode.org/Guides/Multichannel-Expansion.html.
func Multi(inputs ...Input) Inputs {
	return Inputs(inputs)
}
