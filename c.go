package sc

import (
	"math"
	"math/rand"
)

// Pi is a wrapper around math.Pi
const Pi = C(float32(math.Pi))

// C wraps a float32 and implements the Input interface.
type C float32

// Abs computes the absolute value of a signal.
func (c C) Abs() Input {
	return C(float32(math.Abs(float64(c))))
}

// Absdif returns the absolute value of the difference of two inputs.
func (c C) Absdif(val Input) Input {
	return c.Add(val.Neg()).Abs()
}

// Acos computes the arccosine of a signal.
func (c C) Acos() Input {
	return C(float32(math.Acos(float64(c))))
}

// Add adds another input to the constant.
func (c C) Add(val Input) Input {
	if v, ok := val.(C); ok {
		return C(float32(v) + float32(c))
	}
	return val.Add(c)
}

// Amclip returns 0 when b <= 0, a*b when b > 0.
func (c C) Amclip(val Input) Input {
	return val.GT(C(0)).Mul(c.Mul(val))
}

// AmpDb converts linear amplitude to decibels.
func (c C) AmpDb() Input {
	return C(float32(20 * math.Log10(float64(c))))
}

// Asin computes the arcsine of a signal.
func (c C) Asin() Input {
	return C(float32(math.Asin(float64(c))))
}

// Atan computes the arctangent of a signal.
func (c C) Atan() Input {
	return C(float32(math.Atan(float64(c))))
}

// Atan2 returns the arctangent of y/x.
func (c C) Atan2(val Input) Input {
	if v, ok := val.(C); ok {
		return C(float32(math.Atan2(float64(c), float64(v))))
	}
	return val.Reciprocal().Atan2(c.Reciprocal())
}

// Bilinrand returns a linearly distributed random value between [+in ... -in].
func (c C) Bilinrand() Input {
	return c.Rand2()
}

// Ceil computes the ceiling (next highest integer) of a signal.
func (c C) Ceil() Input {
	return C(float32(math.Ceil(float64(c))))
}

// Clip2 clips input wave a to +/- b
func (c C) Clip2(val Input) Input {
	return c.Max(val.Neg()).Min(val)
}

// Coin returns one or zero with the probability given by the input.
func (c C) Coin() Input {
	if rand.Float64() < float64(c) {
		return C(1)
	}
	return C(0)
}

// Cos computes the cosine of an Input.
func (c C) Cos() Input {
	return C(float32(math.Cos(float64(c))))
}

// Cosh computes the hyperboliccosine of an Input.
func (c C) Cosh() Input {
	return C(float32(math.Cosh(float64(c))))
}

// Cpsmidi converts frequency in Hz to midi note values.
func (c C) Cpsmidi() Input {
	return C(Cpsmidi(float32(c)))
}

// Cpsoct converts cycles per second to decimal octaves.
func (c C) Cpsoct() Input {
	return C(Cpsoct(float32(c)))
}

// Cubed computes the cube of a signal.
func (c C) Cubed() Input {
	return C(float32(c * c * c))
}

// DbAmp converts decibels to linear amplitude.
func (c C) DbAmp() Input {
	return C(float32(math.Pow(10, float64(c)/20)))
}

// Difsqr returns the value of (a*a) - (b*b).
func (c C) Difsqr(val Input) Input {
	return c.Squared().Add(val.Squared().Neg())
}

// Distort performs non-linear distortion on a signal.
func (c C) Distort() Input {
	return C(float32(c) / float32(1+math.Abs(float64(c))))
}

// Div divides one input by another.
// This will panic if val is C(0)
func (c C) Div(val Input) Input {
	if v, ok := val.(C); ok {
		return C(c / v)
	}
	return val.Reciprocal().Mul(c)
}

// Excess returns the difference of the original signal and its clipped form: (a - clip2(a,b)).
func (c C) Excess(val Input) Input {
	return c.Add(c.Clip2(val).Neg())
}

// Exp computes the exponential of a signal.
func (c C) Exp() Input {
	return C(float32(math.Exp(float64(c))))
}

// Expon raises a constant to the power of another input.
// If val is not a C the this method just returns c.
// TODO: fix this
func (c C) Expon(val Input) Input {
	if v, ok := val.(C); ok {
		return C(float32(math.Pow(float64(c), float64(v))))
	}
	return c
}

// Floor computes the floor of a constant.
func (c C) Floor() Input {
	return C(float32(math.Floor(float64(c))))
}

// Fold2 folds input wave a to +/- b
func (c C) Fold2(val Input) Input {
	if v, ok := val.(C); ok {
		return C(fold2(float32(c), float32(v)))
	}
	return c // TODO: fix this
}

// Frac returns the fractional part of a constant.
func (c C) Frac() Input {
	return C(float32(float64(c) - math.Trunc(float64(c))))
}

// GCD computes the gcd of one Input and another.
func (c C) GCD(val Input) Input {
	if v, ok := val.(C); ok {
		return C(gcd(float32(c), float32(v)))
	}
	return val.GCD(c)
}

// GT computes x > y.
func (c C) GT(val Input) Input {
	if v, ok := val.(C); ok {
		if c > v {
			return C(1)
		}
		return C(0)
	}
	return val.LT(c)
}

// GTE computes x >= y.
func (c C) GTE(val Input) Input {
	if v, ok := val.(C); ok {
		if c >= v {
			return C(1)
		}
		return C(0)
	}
	return val.LTE(c)
}

// Hypot returns the square root of the sum of the squares of a and b.
// Or equivalently, the distance from the origin to the point (x, y).
func (c C) Hypot(val Input) Input {
	if v, ok := val.(C); ok {
		return C(float32(math.Hypot(float64(c), float64(v))))
	}
	return val.Hypot(c)
}

// HypotApx returns an approximation of the square root of the sum of the squares of x and y.
// This uses the formula:
//     abs(x) + abs(y) - ((sqrt(2) - 1) * min(abs(x), abs(y)))
func (c C) HypotApx(val Input) Input {
	if v, ok := val.(C); ok {
		var (
			x = float32(math.Abs(float64(c))) + float32(math.Abs(float64(v)))
			y = float32(math.Min(math.Abs(float64(c)), math.Abs(float64(v))))
		)
		return C(x - ((float32(math.Sqrt(2)) - 1) * y))
	}
	return val.HypotApx(c)
}

// LCM computes the least common multiple of one Input and another.
func (c C) LCM(val Input) Input {
	if v, ok := val.(C); ok {
		return C(lcm(float32(c), float32(v)))
	}
	return val.LCM(c)
}

// LT computes x < y.
func (c C) LT(val Input) Input {
	if v, ok := val.(C); ok {
		if c < v {
			return C(1)
		}
		return C(0)
	}
	return val.GT(c)
}

// LTE computes x <= y.
func (c C) LTE(val Input) Input {
	if v, ok := val.(C); ok {
		if c <= v {
			return C(1)
		}
		return C(0)
	}
	return val.GTE(c)
}

// Linrand returns a linearly distributed random value between in and zero.
func (c C) Linrand() Input {
	return c.Rand()
}

// Log computes a natural logarithm.
func (c C) Log() Input {
	return C(math.Log(float64(c)))
}

// Log10 computes a natural logarithm.
func (c C) Log10() Input {
	return C(math.Log10(float64(c)))
}

// Log2 computes a natural logarithm.
func (c C) Log2() Input {
	return C(math.Log2(float64(c)))
}

// Max returns the maximum of one input and another.
func (c C) Max(other Input) Input {
	if v, ok := other.(C); ok {
		return C(maxFloat32(float32(c), float32(v)))
	}
	return other.Max(c)
}

// Midicps converts MIDI note number to cycles per second.
func (c C) Midicps() Input {
	return C(Midicps(float32(c)))
}

// Midiratio converts an interval in MIDI notes into a frequency ratio.
func (c C) Midiratio() Input {
	return C(float32(math.Pow(2, float64(c)/12)))
}

// Min returns the minimum of one signal and another.
func (c C) Min(other Input) Input {
	if v, ok := other.(C); ok {
		return C(minFloat32(float32(c), float32(v)))
	}
	return other.Min(c)
}

// Moddif returns the smaller of the great circle distances between the two points.
func (c C) Moddif(y, mod Input) Input {
	var (
		diff    = c.Absdif(y).Modulo(mod)
		modhalf = mod.Mul(C(0.5))
	)
	return modhalf.Add(diff.Absdif(modhalf).Neg())
}

// Modulo computes the modulo of one signal and another.
// If val is not a C, then this method just returns the receiver.
// I'm not sure what a constant modulo a ugen should be.
// Note that Go only supports integers for the modulo operator.
func (c C) Modulo(val Input) Input {
	if v, ok := val.(C); ok {
		return C(float32(int(c) % int(v)))
	}
	return c
}

// Mul multiplies the constant by another input.
func (c C) Mul(val Input) Input {
	if v, ok := val.(C); ok {
		return C(float32(v) * float32(c))
	}
	return val.Mul(c)
}

// MulAdd multiplies and adds at the same time.
func (c C) MulAdd(mul, add Input) Input {
	if m, mok := mul.(C); mok {
		if a, aok := add.(C); aok {
			return C((float32(m) * float32(c)) + float32(a))
		}
		return add.MulAdd(c, mul)
	}
	return mul.MulAdd(c, add)
}

// Neg is a convenience operator that multiplies a signal by -1.
func (c C) Neg() Input {
	return C(float32(c) * -1)
}

// Octcps converts decimal octaves to cycles per second.
func (c C) Octcps() Input {
	return C(Octcps(float32(c)))
}

// Pow raises a constant to the power of another input.
// If val is not a C the this method just returns c.
// TODO: fix this
func (c C) Pow(val Input) Input {
	if v, ok := val.(C); ok {
		return C(float32(math.Pow(float64(c), float64(v))))
	}
	return c
}

// Rand returns an evenly distributed random value between this and zero.
func (c C) Rand() Input {
	return C(float32(math.Trunc(float64(rand.Float32() * float32(c)))))
}

// Rand2 returns an evenly distributed random value between [+this ... - this].
func (c C) Rand2() Input {
	v := (rand.Float32() * 2 * float32(c)) - float32(c)
	return C(float32(math.Trunc(float64(v))))
}

// Ratiomidi converts a frequency ratio to an interval in MIDI notes.
func (c C) Ratiomidi() Input {
	return C(float32(12 * math.Log2(float64(c))))
}

// Reciprocal computes the reciprocal of a signal.
func (c C) Reciprocal() Input {
	return C(1 / float32(c))
}

// Ring1 returns the value of ((a*b) + a).
func (c C) Ring1(val Input) Input {
	if v, ok := val.(C); ok {
		return C((c * v) + c)
	}
	return val.Ring1(c)
}

// Ring2 returns the value of ((a*b) + a + b).
func (c C) Ring2(val Input) Input {
	if v, ok := val.(C); ok {
		return C((c * v) + c + v)
	}
	return val.Ring2(c)
}

// Ring3 returns the value of (a*a*b).
func (c C) Ring3(val Input) Input {
	if v, ok := val.(C); ok {
		return C((c * v) * v)
	}
	return val.Ring3(c)
}

// Ring4 returns the value of ((a*a *b) - (a*b*b)).
func (c C) Ring4(val Input) Input {
	if v, ok := val.(C); ok {
		return C(((c * v) * v) - (c * v * v))
	}
	return val.Ring4(c)
}

// Round performs quantization by rounding. Rounds a to the nearest multiple of b.
func (c C) Round(val Input) Input {
	if v, ok := val.(C); ok {
		return C(Roundf(float32(c), float32(v)))
	}
	return c
}

// Scaleneg returns a*b when a < 0, otherwise a.
func (c C) Scaleneg(val Input) Input {
	if c < 0 {
		return c.Mul(val)
	}
	return c
}

// Sign computes the sign of the constant.
func (c C) Sign() Input {
	if c > 0 {
		return C(1)
	} else if c < 0 {
		return C(-1)
	}
	return C(0)
}

// Sin computes the sine of an Input.
func (c C) Sin() Input {
	return C(float32(math.Sin(float64(c))))
}

// Sinh computes the hyperbolic sine of an Input.
func (c C) Sinh() Input {
	return C(float32(math.Sinh(float64(c))))
}

// SoftClip clips the constant to the range [-0.5, 0.5]
func (c C) SoftClip() Input {
	if float32(c) < -0.5 {
		return C(-0.5)
	} else if float32(c) > 0.5 {
		return C(0.5)
	}
	return c
}

// Sqrt computes the square root of a constant.
func (c C) Sqrt() Input {
	return C(math.Sqrt(float64(c)))
}

// Sqrdif computes the square of the difference between the two inputs.
func (c C) Sqrdif(val Input) Input {
	return c.Add(val.Neg()).Squared()
}

// Sqrsum computes the square of the sum of the two inputs.
func (c C) Sqrsum(val Input) Input {
	return c.Add(val).Squared()
}

// Squared computes the square of a signal.
func (c C) Squared() Input {
	return C(float32(c * c))
}

// Sum3rand returns a value from a gaussian-like random distribution between in and zero.
func (c C) Sum3rand() Input {
	return C(rand.NormFloat64())
}

// Sumsqr returns the value of (a*a) + (b*b).
func (c C) Sumsqr(val Input) Input {
	return c.Squared().Add(val.Squared())
}

// Tan computes the tangent of an Input.
func (c C) Tan() Input {
	return C(float32(math.Tan(float64(c))))
}

// Tanh computes the hyperbolic tangent of an Input.
func (c C) Tanh() Input {
	return C(float32(math.Tanh(float64(c))))
}

// Thresh returns 0 when c < val, otherwise c.
func (c C) Thresh(val Input) Input {
	if v, ok := val.(C); ok {
		if c < v {
			return C(0)
		}
		return c
	}
	return val.GTE(c).Mul(c)
}

// Trunc performs quantization by truncation. Truncate c to a multiple of val.
// If val is not a constant, c is returned.
func (c C) Trunc(val Input) Input {
	if v, ok := val.(C); ok {
		return C(Truncf(float32(c), float32(v)))
	}
	return c
}

// Wrap2 wraps input wave to +/-b
func (c C) Wrap2(val Input) Input {
	if v, ok := val.(C); ok {
		return C(wrap2(float32(c), float32(v)))
	}
	return c // TODO: fix this
}

// Roundf rounds a to the nearest multiple of b.
func Roundf(a, b float32) float32 {
	if b == 0 {
		return 0
	}
	var m, v float32

	for {
		if b*m <= a {
			if b*(m+1) > a {
				if math.Abs(float64(a-b*m)) < math.Abs(float64(a-b*(m+1))) {
					v = b * m
				} else {
					v = b * (m + 1)
				}
				break
			}
			m++
			continue
		}
		m--
		if b*m <= a {
			if math.Abs(float64(a-b*m)) < math.Abs(float64(a-b*(m+1))) {
				v = b * m
			} else {
				v = b * (m + 1)
			}
			break
		}
	}
	return v
}

// Truncf returns the next highest multiple of b that is < a.
func Truncf(a, b float32) float32 {
	if b == 0 {
		return 0
	}
	var m, v float32

	for {
		if b*m <= a {
			if b*(m+1) > a {
				v = b * m
				break
			}
			m++
			continue
		}
		m--
		if b*m <= a {
			v = b * m
			break
		}
	}
	return v
}

func fold2(x, y float32) float32 {
	if y < 0 {
		y *= -1
	} else if y == 0 {
		return 0
	}
	if x <= y && x >= -y {
		return x
	}
	if x > y {
		return fold2(y-(x-y), y)
	}
	return fold2(-y+(-y-x), y)
}

func gcd(x, y float32) float32 {
	a, b := int(x), int(y)
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return float32(a)
}

func lcm(x, y float32) float32 {
	a, b := int(x), int(y)
	return float32(a*b) / gcd(x, y)
}

func maxFloat32(f1, f2 float32) float32 {
	if f1 > f2 {
		return f1
	}
	return f2
}

func minFloat32(f1, f2 float32) float32 {
	if f1 < f2 {
		return f1
	}
	return f2
}

func wrap2(x, y float32) float32 {
	if y < 0 {
		y *= -1
	} else if y == 0 {
		return 0
	}
	if x <= y && x >= -y {
		return x
	}
	if x > y {
		return wrap2(x-(2*y), y)
	}
	return wrap2(x+(2*y), y)
}
