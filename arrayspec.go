package sc

import (
	"fmt"
)

// ArraySpec is a convenience type for Klang and Klank.
type ArraySpec [3][]Input

// Abs computes the absolute value of a signal.
func (as ArraySpec) Abs() Input {
	return as.proc(func(i Input) Input {
		return i.Abs()
	})
}

// Absdif returns the absolute value of the difference of two inputs.
func (as ArraySpec) Absdif(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Absdif(val)
	})
}

// Acos computes the arccosine of a signal.
func (as ArraySpec) Acos() Input {
	return as.proc(func(i Input) Input {
		return i.Acos()
	})
}

// Add adds one input to another.
func (as ArraySpec) Add(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Add(val)
	})
}

// Amclip returns 0 when b <= 0, a*b when b > 0.
func (as ArraySpec) Amclip(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Amclip(val)
	})
}

// AmpDb converts linear amplitude to decibels.
func (as ArraySpec) AmpDb() Input {
	return as.proc(func(i Input) Input {
		return i.AmpDb()
	})
}

// Asin computes the arcsine of a signal.
func (as ArraySpec) Asin() Input {
	return as.proc(func(i Input) Input {
		return i.Asin()
	})
}

// Atan computes the arctangent of a signal.
func (as ArraySpec) Atan() Input {
	return as.proc(func(i Input) Input {
		return i.Atan()
	})
}

// Atan2 returns the arctangent of y/x.
func (as ArraySpec) Atan2(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Atan2(val)
	})
}

// Bilinrand returns a linearly distributed random value between [+in ... -in].
func (as ArraySpec) Bilinrand() Input {
	return as.proc(func(i Input) Input {
		return i.Bilinrand()
	})
}

// Ceil computes the ceiling of a signal.
func (as ArraySpec) Ceil() Input {
	return as.proc(func(i Input) Input {
		return i.Ceil()
	})
}

// Coin returns one or zero with the probability given by the input.
func (as ArraySpec) Coin() Input {
	return as.proc(func(i Input) Input {
		return i.Coin()
	})
}

// Clip2 clips input wave a to +/- b
func (as ArraySpec) Clip2(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Clip2(val)
	})
}

// Cos computes the cosine of a signal.
func (as ArraySpec) Cos() Input {
	return as.proc(func(i Input) Input {
		return i.Cos()
	})
}

// Cosh computes the hyperbolic cosine of a signal.
func (as ArraySpec) Cosh() Input {
	return as.proc(func(i Input) Input {
		return i.Cosh()
	})
}

// Cpsmidi converts frequency in Hz to midi note values.
func (as ArraySpec) Cpsmidi() Input {
	return as.proc(func(i Input) Input {
		return i.Cpsmidi()
	})
}

// Cpsoct converts cycles per second to decimal octaves.
func (as ArraySpec) Cpsoct() Input {
	return as.proc(func(i Input) Input {
		return i.Cpsoct()
	})
}

// Cubed computes the cube of a signal.
func (as ArraySpec) Cubed() Input {
	return as.proc(func(i Input) Input {
		return i.Cubed()
	})
}

// DbAmp converts decibels tolinear amplitude.
func (as ArraySpec) DbAmp() Input {
	return as.proc(func(i Input) Input {
		return i.DbAmp()
	})
}

// Difsqr returns the value of (a*a) - (b*b).
func (as ArraySpec) Difsqr(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Difsqr(val)
	})
}

// Distort performs non-linear distortion on a signal.
func (as ArraySpec) Distort() Input {
	return as.proc(func(i Input) Input {
		return i.Distort()
	})
}

// Div divides one input by another.
func (as ArraySpec) Div(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Div(val)
	})
}

// Excess returns the difference of the original signal and its clipped form: (a - clip2(a,b)).
func (as ArraySpec) Excess(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Excess(val)
	})
}

// Exp computes the exponential of a signal.
func (as ArraySpec) Exp() Input {
	return as.proc(func(i Input) Input {
		return i.Exp()
	})
}

// Expon raises an Input to the power of another.
func (as ArraySpec) Expon(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Expon(val)
	})
}

// Floor computes the floor of a signal.
func (as ArraySpec) Floor() Input {
	return as.proc(func(i Input) Input {
		return i.Floor()
	})
}

// Fold2 folds input wave a to +/- b
func (as ArraySpec) Fold2(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Fold2(val)
	})
}

// Frac returns the fractional part of a signal.
func (as ArraySpec) Frac() Input {
	return as.proc(func(i Input) Input {
		return i.Frac()
	})
}

// GCD computes the gcd of one Input and another.
func (as ArraySpec) GCD(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.GCD(val)
	})
}

// GT computes x > y.
func (as ArraySpec) GT(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.GT(val)
	})
}

// GTE computes x >= y.
func (as ArraySpec) GTE(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.GTE(val)
	})
}

// Hypot returns the square root of the sum of the squares of a and b.
// Or equivalently, the distance from the origin to the point (x, y).
func (as ArraySpec) Hypot(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Hypot(val)
	})
}

// HypotApx returns an approximation of the square root of the sum of the squares of x and y.
func (as ArraySpec) HypotApx(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.HypotApx(val)
	})
}

// LCM computes the least common multiple of one Input and another.
func (as ArraySpec) LCM(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.LCM(val)
	})
}

// LT computes x < y.
func (as ArraySpec) LT(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.LT(val)
	})
}

// LTE computes x <= y.
func (as ArraySpec) LTE(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.LTE(val)
	})
}

// Linrand returns a linearly distributed random value between in and zero.
func (as ArraySpec) Linrand() Input {
	return as.proc(func(i Input) Input {
		return i.Linrand()
	})
}

// Log computes a natural logarithm.
func (as ArraySpec) Log() Input {
	return as.proc(func(i Input) Input {
		return i.Log()
	})
}

// Log10 computes a base 10 logarithm.
func (as ArraySpec) Log10() Input {
	return as.proc(func(i Input) Input {
		return i.Log10()
	})
}

// Log2 computes a base 2 logarithm.
func (as ArraySpec) Log2() Input {
	return as.proc(func(i Input) Input {
		return i.Log2()
	})
}

// Max returns the maximum of one input and another.
func (as ArraySpec) Max(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Max(val)
	})
}

// Midiratio converts an interval in MIDI notes into a frequency ratio.
func (as ArraySpec) Midiratio() Input {
	return as.proc(func(i Input) Input {
		return i.Midiratio()
	})
}

// Midicps converts from MIDI note values to cycles per second.
func (as ArraySpec) Midicps() Input {
	return as.proc(func(i Input) Input {
		return i.Midicps()
	})
}

// Min returns the minimum of one signal and another.
func (as ArraySpec) Min(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Min(val)
	})
}

// Moddif returns the smaller of the great circle distances between the two points.
func (as ArraySpec) Moddif(y, mod Input) Input {
	return as.proc(func(i Input) Input {
		return i.Moddif(y, mod)
	})
}

// Modulo computes the modulo of one signal and another.
func (as ArraySpec) Modulo(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Modulo(val)
	})
}

// Mul multiplies one input and another.
func (as ArraySpec) Mul(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Mul(val)
	})
}

// MulAdd computes (as * m) + a.
func (as ArraySpec) MulAdd(m, a Input) Input {
	return as.proc(func(i Input) Input {
		return i.MulAdd(m, a)
	})
}

// Neg negates an input.
func (as ArraySpec) Neg() Input {
	return as.proc(func(i Input) Input {
		return i.Neg()
	})
}

// Octcps converts decimal octaves to cycles per second.
func (as ArraySpec) Octcps() Input {
	return as.proc(func(i Input) Input {
		return i.Octcps()
	})
}

// Pow raises an Input to the power of another.
func (as ArraySpec) Pow(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Pow(val)
	})
}

// Rand returns an evenly distributed random value between this and zero.
func (as ArraySpec) Rand() Input {
	return as.proc(func(i Input) Input {
		return i.Rand()
	})
}

// Rand2 returns an evenly distributed random value between [+this ... - this].
func (as ArraySpec) Rand2() Input {
	return as.proc(func(i Input) Input {
		return i.Rand2()
	})
}

// Ratiomidi converts a frequency ratio to an interval in MIDI notes.
func (as ArraySpec) Ratiomidi() Input {
	return as.proc(func(i Input) Input {
		return i.Ratiomidi()
	})
}

// Reciprocal computes the reciprocal of a signal.
func (as ArraySpec) Reciprocal() Input {
	return as.proc(func(i Input) Input {
		return i.Reciprocal()
	})
}

// Ring1 is ring modulation plus first source.
func (as ArraySpec) Ring1(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Ring1(val)
	})
}

// Ring2 is ring modulation plus both sources.
func (as ArraySpec) Ring2(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Ring2(val)
	})
}

// Ring3 returns the value of (a*a *b)
func (as ArraySpec) Ring3(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Ring3(val)
	})
}

// Ring4 returns the value of ((a*a *b) - (a*b*b)).
func (as ArraySpec) Ring4(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Ring4(val)
	})
}

// Round performs quantization by rounding. Rounds a to the nearest multiple of b.
func (as ArraySpec) Round(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Round(val)
	})
}

// Scaleneg returns a*b when a < 0, otherwise a.
func (as ArraySpec) Scaleneg(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Scaleneg(val)
	})
}

// Sign returns the sign of a signal.
func (as ArraySpec) Sign() Input {
	return as.proc(func(i Input) Input {
		return i.Sign()
	})
}

// Sin computes the sine of a signal.
func (as ArraySpec) Sin() Input {
	return as.proc(func(i Input) Input {
		return i.Sin()
	})
}

// Sinh computes the hyperbolic sine of a signal.
func (as ArraySpec) Sinh() Input {
	return as.proc(func(i Input) Input {
		return i.Sinh()
	})
}

// SoftClip computes nonlinear distortion of an input.
func (as ArraySpec) SoftClip() Input {
	return as.proc(func(i Input) Input {
		return i.SoftClip()
	})
}

// Sqrdif computes the square of the difference between the two inputs.
func (as ArraySpec) Sqrdif(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Sqrdif(val)
	})
}

// Sqrsum computes the square of the sum of the two inputs.
func (as ArraySpec) Sqrsum(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Sqrsum(val)
	})
}

// Sqrt returns the square root of a signal.
func (as ArraySpec) Sqrt() Input {
	return as.proc(func(i Input) Input {
		return i.Sqrt()
	})
}

// Squared computes the square of a signal.
func (as ArraySpec) Squared() Input {
	return as.proc(func(i Input) Input {
		return i.Squared()
	})
}

// Sum3rand returns a value from a gaussian-like random distribution between in and zero.
func (as ArraySpec) Sum3rand() Input {
	return as.proc(func(i Input) Input {
		return i.Sum3rand()
	})
}

// Sumsqr returns the value of (a*a) + (b*b).
func (as ArraySpec) Sumsqr(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Sumsqr(val)
	})
}

// Tan computes the tangent of a signal.
func (as ArraySpec) Tan() Input {
	return as.proc(func(i Input) Input {
		return i.Tan()
	})
}

// Tanh computes the hyperbolic tangent of a signal.
func (as ArraySpec) Tanh() Input {
	return as.proc(func(i Input) Input {
		return i.Tanh()
	})
}

// Thresh returns 0 when a < b, otherwise a.
func (as ArraySpec) Thresh(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Thresh(val)
	})
}

// Trunc performs quantization by truncation. Truncate a to a multiple of b.
func (as ArraySpec) Trunc(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Trunc(val)
	})
}

// Wrap2 wraps input wave to +/-b
func (as ArraySpec) Wrap2(val Input) Input {
	return as.proc(func(i Input) Input {
		return i.Wrap2(val)
	})
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

// proc processes the inputs in this arrayspec
func (as ArraySpec) proc(f func(Input) Input) Input {
	var nas ArraySpec
	for i := range []int{0, 1, 2} {
		for j := range nas[i] {
			nas[i][j] = f(as[i][j])
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
