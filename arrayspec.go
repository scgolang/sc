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

// Distort performs non-linear distortion on a signal.
func (as ArraySpec) Distort() Input {
	return as.proc(func(i Input) Input {
		return i.Distort()
	})
}

// Exp computes the exponential of a signal.
func (as ArraySpec) Exp() Input {
	return as.proc(func(i Input) Input {
		return i.Exp()
	})
}

// Floor computes the floor of a signal.
func (as ArraySpec) Floor() Input {
	return as.proc(func(i Input) Input {
		return i.Floor()
	})
}

// Frac returns the fractional part of a signal.
func (as ArraySpec) Frac() Input {
	return as.proc(func(i Input) Input {
		return i.Frac()
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
