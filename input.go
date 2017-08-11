package sc

import (
	"encoding/binary"
	"io"
)

// Input is implemented by any value that can serve as a
// ugen input. This includes synthdef parameters,
// constants, and other ugens.
type Input interface {
	// Abs computes the absolute value of a signal.
	Abs() Input

	// Absdif returns the absolute value of the difference of two inputs.
	Absdif(Input) Input

	// Acos computes the arccosine of a signal.
	Acos() Input

	// Amclip returns 0 when b <= 0, a*b when b > 0.
	Amclip(Input) Input

	// Asin computes the arcsine of a signal.
	Asin() Input

	// Atan computes the arctangent of a signal.
	Atan() Input

	// Atan2 returns the arctangent of y/x.
	Atan2(val Input) Input

	// Add adds one Input to another.
	Add(val Input) Input

	// AmpDb converts linear amplitude to decibels.
	AmpDb() Input

	// Bilinrand returns a linearly distributed random value between [+in ... -in].
	Bilinrand() Input

	// Ceil computes the ceiling (next lowest integer) of a signal.
	Ceil() Input

	// Clip2 clips input wave a to +/- b
	Clip2(Input) Input

	// Coin returns one or zero with the probability given by the input.
	Coin() Input

	// Cos computes the cosine of an Input.
	Cos() Input

	// Cosh computes the hyperbolic cosine of an Input.
	Cosh() Input

	// Cpsmidi converts frequency in Hz to midi note values.
	Cpsmidi() Input

	// Cpsoct converts cycles per second to decimal octaves.
	Cpsoct() Input

	// Cubed raises a signal to the power of 3.
	Cubed() Input

	// DbAmp converts decibels to linear amplitude.
	DbAmp() Input

	// Difsqr returns the value of (a*a) - (b*b).
	Difsqr(Input) Input

	// Distort performs non-linear distortion on a signal.
	Distort() Input

	// Div divides one input by another.
	Div(Input) Input

	// Exp computes the exponential of a signal.
	Exp() Input

	// Expon raises this input to the power of another.
	// When used with UGens which produce a negative signal this function extends
	// the usual definition of exponentiation and returns neg(neg(a) ** b).
	// This allows exponentiation of negative signal values by noninteger exponents.
	// For the normal behaviour use Pow (see below).
	Expon(Input) Input

	// Excess returns the difference of the original signal and its clipped form: (a - clip2(a,b)).
	Excess(Input) Input

	// Floor computes the floor (next lowest integer) of a signal.
	Floor() Input

	// Fold2 folds input wave a to +/- b
	Fold2(Input) Input

	// Frac returns the fractional part of a signal.
	Frac() Input

	// GCD computes the gcd of one Input and another.
	GCD(Input) Input

	// GT computes x > y.
	GT(Input) Input

	// GT computes x >= y.
	GTE(Input) Input

	// Hypot returns the square root of the sum of the squares of a and b.
	// Or equivalently, the distance from the origin to the point (x, y).
	Hypot(Input) Input

	// HypotApx returns an approximation of the square root of the sum of the squares of x and y.
	HypotApx(Input) Input

	// LCM computes the least common multiple of one Input and another.
	LCM(Input) Input

	// LT computes x < y.
	LT(Input) Input

	// LT computes x <= y.
	LTE(Input) Input

	// Linrand returns a linearly distributed random value between in and zero.
	Linrand() Input

	// Log computes a natural logarithm.
	Log() Input

	// Log10 computes a base 10 logarithm.
	Log10() Input

	// Log2 computes a base 2 logarithm.
	Log2() Input

	// Max returns the max of one signal and another.
	Max(other Input) Input

	// Midicps converts MIDI note number to cycles per second.
	Midicps() Input

	// Midiratio converts an interval in MIDI notes into a frequency ratio.
	Midiratio() Input

	// Min returns the minimum of one signal and another.
	Min(Input) Input

	// Moddif returns the smaller of the great circle distances between the two points.
	Moddif(Input, Input) Input

	// Modulo computes the modulo of one signal and another.
	Modulo(Input) Input

	// Mul multiplies one Input by another.
	Mul(val Input) Input

	// MulAdd multiplies and adds an Input using two others.
	MulAdd(mul, add Input) Input

	// Neg is a convenience operator that multiplies a signal by -1.
	Neg() Input

	// Octcps converts decimal octaves to cycles per second.
	Octcps() Input

	// Pow raises an Input to the power of another.
	Pow(Input) Input

	// Rand returns an evenly distributed random value between this and zero.
	Rand() Input

	// Rand2 returns an evenly distributed random value between [+this ... - this].
	Rand2() Input

	// Ratiomidi converts a frequency ratio to an interval in MIDI notes.
	Ratiomidi() Input

	// Reciprocal computes the reciprocal of a signal.
	Reciprocal() Input

	// Ring1 is ring modulation plus first source.
	Ring1(val Input) Input

	// Ring2 is ring modulation plus both sources.
	Ring2(val Input) Input

	// Ring3 returns the value of (a*a *b)
	Ring3(val Input) Input

	// Ring4 returns the value of ((a*a *b) - (a*b*b)).
	Ring4(val Input) Input

	// Round performs quantization by rounding. Rounds a to the nearest multiple of b.
	Round(Input) Input

	// Scaleneg returns a*b when a < 0, otherwise a.
	Scaleneg(Input) Input

	// Sign computes the sign of a signal.
	// This returns -1 when a < 0, +1 when a > 0, 0 when a is 0.
	Sign() Input

	// Sin computes the sine of an Input.
	Sin() Input

	// Sinh computes the hyperbolic sine of an Input.
	Sinh() Input

	// SoftClip distorts a signal with a perfectly linear range from -0.5 to 0.5
	SoftClip() Input

	// Sqrdif computes the square of the difference between the two inputs.
	Sqrdif(Input) Input

	// Sqrsum computes the square of the sum of the two inputs.
	Sqrsum(Input) Input

	// Sqrt computes the square root of a signal.
	// The definition of square root is extended for signals so that sqrt(a)
	// when a < 0 returns -sqrt(-a).
	Sqrt() Input

	// Squared raises a signal to the power of 2.
	Squared() Input

	// Sum3rand returns a value from a gaussian-like random distribution between in and zero.
	Sum3rand() Input

	// Sumsqr returns the value of (a*a) + (b*b).
	Sumsqr(Input) Input

	// Tan computes the tangent of an Input.
	Tan() Input

	// Tanh computes the hyperbolic tangent of an Input.
	Tanh() Input

	// Thresh returns 0 when a < b, otherwise a.
	Thresh(Input) Input

	// Trunc performs quantization by truncation. Truncate a to a multiple of b.
	Trunc(Input) Input

	// Wrap2 wraps input wave to +/- b
	Wrap2(Input) Input
}

func readInput(r io.Reader) (UgenInput, error) {
	var ui UgenInput
	if err := binary.Read(r, byteOrder, &ui.UgenIndex); err != nil {
		return ui, err
	}
	if err := binary.Read(r, byteOrder, &ui.OutputIndex); err != nil {
		return ui, err
	}
	return ui, nil
}
