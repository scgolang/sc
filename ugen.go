package sc

import (
	"encoding/binary"
	"io"
)

// UGen done actions, see http://doc.sccode.org/Reference/UGen-doneActions.html
const (
	DoNothing = iota
	Pause
	FreeEnclosing
	FreePreceding
	FreeFollowing
	FreePrecedingGroup
	FreeFollowingGroup
	FreeAllPreceding
	FreeAllFollowing
	FreeAndPausePreceding
	FreeAndPauseFollowing
	DeepFreePreceding
	DeepFreeFollowing
	FreeAllInGroup
	// I do not understand the difference between the last and
	// next-to-last options [bps]
)

// Ugen is a unit generator.
type Ugen struct {
	Name         string      `json:"name"              xml:"name,attr"`
	Rate         int8        `json:"rate"              xml:"rate,attr"`
	SpecialIndex int16       `json:"specialIndex"      xml:"specialIndex,attr"`
	Inputs       []UgenInput `json:"inputs,omitempty"  xml:"Inputs>Input"`
	Outputs      []Output    `json:"outputs,omitempty" xml:"Outputs>Output"`
	NumOutputs   int         `json:"numOutputs"        xml:"numOutputs,attr"`

	inputs []Input
}

// NewUgen is a factory function for creating new Ugen instances.
// Panics if rate is not AR, KR, or IR.
// Panics if numOutputs <= 0.
func NewUgen(name string, rate int8, specialIndex int16, numOutputs int, inputs ...Input) *Ugen {
	CheckRate(rate)

	if numOutputs <= 0 {
		panic("numOutputs must be a positive int")
	}
	// TODO: validate specialIndex
	u := &Ugen{
		Name:         name,
		Rate:         rate,
		SpecialIndex: specialIndex,
		NumOutputs:   numOutputs,
		inputs:       inputs,
	}
	// If any inputs are multi inputs, then this node should get promoted to a multi node.
	for i, input := range inputs {
		switch v := input.(type) {
		case *Ugen:
			// Initialize the Outputs slice.
			input = asOutput(v)
		case MultiInput:
			ins := make(Inputs, len(v.InputArray()))

			// Add outputs to any nodes in a MultiInput.
			for i, in := range v.InputArray() {
				if u, ok := in.(*Ugen); ok {
					ins[i] = asOutput(u)
				}
			}
		}
		u.inputs[i] = input
	}
	return u
}

// Abs computes the absolute value of a signal.
func (u *Ugen) Abs() Input {
	return unaryOpAbs(u.Rate, u, u.NumOutputs)
}

// Absdif returns the absolute value of the difference of two inputs.
func (u *Ugen) Absdif(val Input) Input {
	return binOpAbsdif(u.Rate, u, val, u.NumOutputs)
}

// Acos computes the arccosine of a signal.
func (u *Ugen) Acos() Input {
	return unaryOpAcos(u.Rate, u, u.NumOutputs)
}

// Add adds an input to a ugen node.
func (u *Ugen) Add(val Input) Input {
	return binOpAdd(u.Rate, u, val, u.NumOutputs)
}

// Amclip returns 0 when b <= 0, a*b when b > 0.
func (u *Ugen) Amclip(val Input) Input {
	return binOpAmclip(u.Rate, u, val, u.NumOutputs)
}

// AmpDb converts linear amplitude to decibels.
func (u *Ugen) AmpDb() Input {
	return unaryOpAmpDb(u.Rate, u, u.NumOutputs)
}

// Asin computes the arcsine of a signal.
func (u *Ugen) Asin() Input {
	return unaryOpAsin(u.Rate, u, u.NumOutputs)
}

// Atan computes the arctangent of a signal.
func (u *Ugen) Atan() Input {
	return unaryOpAtan(u.Rate, u, u.NumOutputs)
}

// Atan2 returns the arctangent of y/x.
func (u *Ugen) Atan2(val Input) Input {
	return binOpAtan2(u.Rate, u, val, u.NumOutputs)
}

// Bilinrand returns a linearly distributed random value between [+in ... -in].
func (u *Ugen) Bilinrand() Input {
	return unaryOpBilinrand(u.Rate, u, u.NumOutputs)
}

// Ceil computes the ceiling (next highest integer) of a signal.
func (u *Ugen) Ceil() Input {
	return unaryOpCeil(u.Rate, u, u.NumOutputs)
}

// Clip2 clips input wave a to +/- b
func (u *Ugen) Clip2(val Input) Input {
	return binOpClip2(u.Rate, u, val, u.NumOutputs)
}

// Coin returns one or zero with the probability given by the input.
func (u *Ugen) Coin() Input {
	return unaryOpCoin(u.Rate, u, u.NumOutputs)
}

// Cos returns the cosine of a ugen.
func (u *Ugen) Cos() Input {
	return unaryOpCos(u.Rate, u, u.NumOutputs)
}

// Cosh returns the hyperbolic cosine of a ugen.
func (u *Ugen) Cosh() Input {
	return unaryOpCosh(u.Rate, u, u.NumOutputs)
}

// Cpsmidi converts frequency in Hz to midi note values.
func (u *Ugen) Cpsmidi() Input {
	return unaryOpCpsmidi(u.Rate, u, u.NumOutputs)
}

// Cpsoct converts cycles per second to decimal octaves.
func (u *Ugen) Cpsoct() Input {
	return unaryOpCpsoct(u.Rate, u, u.NumOutputs)
}

// Cubed computes the cube of a signal.
func (u *Ugen) Cubed() Input {
	return unaryOpCubed(u.Rate, u, u.NumOutputs)
}

// DbAmp converts linear amplitude to decibels.
func (u *Ugen) DbAmp() Input {
	return unaryOpDbAmp(u.Rate, u, u.NumOutputs)
}

// Difsqr returns the value of (a*a) - (b*b).
func (u *Ugen) Difsqr(val Input) Input {
	return binOpDifsqr(u.Rate, u, val, u.NumOutputs)
}

// Distort performs non-linear distortion on a signal.
func (u *Ugen) Distort() Input {
	return unaryOpDistort(u.Rate, u, u.NumOutputs)
}

// Div divides one input by another.
func (u *Ugen) Div(val Input) Input {
	return binOpDiv(u.Rate, u, val, u.NumOutputs)
}

// Excess returns the difference of the original signal and its clipped form: (a - clip2(a,b)).
func (u *Ugen) Excess(val Input) Input {
	return binOpExcess(u.Rate, u, val, u.NumOutputs)
}

// Exp computes the exponential of a signal.
func (u *Ugen) Exp() Input {
	return unaryOpExp(u.Rate, u, u.NumOutputs)
}

// Expon raises this input to the power of another.
func (u *Ugen) Expon(val Input) Input {
	return binOpExpon(u.Rate, u, val, u.NumOutputs)
}

// Floor computes the floor (next lowest integer) of a signal.
func (u *Ugen) Floor() Input {
	return unaryOpFloor(u.Rate, u, u.NumOutputs)
}

// Fold2 folds input wave a to +/- b
func (u *Ugen) Fold2(val Input) Input {
	return binOpFold2(u.Rate, u, val, u.NumOutputs)
}

// Frac computes the fractional part of a signal.
func (u *Ugen) Frac() Input {
	return unaryOpFrac(u.Rate, u, u.NumOutputs)
}

// GCD computes the gcd of one Input and another.
func (u *Ugen) GCD(val Input) Input {
	return binOpGCD(u.Rate, u, val, u.NumOutputs)
}

// GT computes x > y.
func (u *Ugen) GT(val Input) Input {
	return binOpGT(u.Rate, u, val, u.NumOutputs)
}

// GTE computes x >= y.
func (u *Ugen) GTE(val Input) Input {
	return binOpGTE(u.Rate, u, val, u.NumOutputs)
}

// Hypot returns the square root of the sum of the squares of a and b.
// Or equivalently, the distance from the origin to the point (x, y).
func (u *Ugen) Hypot(val Input) Input {
	return binOpHypot(u.Rate, u, val, u.NumOutputs)
}

// HypotApx returns an approximation of the square root of the sum of the squares of x and y.
func (u *Ugen) HypotApx(val Input) Input {
	return binOpHypotApx(u.Rate, u, val, u.NumOutputs)
}

// LCM computes the lcm of one Input and another.
func (u *Ugen) LCM(val Input) Input {
	return binOpLCM(u.Rate, u, val, u.NumOutputs)
}

// LT computes x < y.
func (u *Ugen) LT(val Input) Input {
	return binOpLT(u.Rate, u, val, u.NumOutputs)
}

// LTE computes x <= y.
func (u *Ugen) LTE(val Input) Input {
	return binOpLTE(u.Rate, u, val, u.NumOutputs)
}

// Linrand returns a linearly distributed random value between in and zero.
func (u *Ugen) Linrand() Input {
	return unaryOpLinrand(u.Rate, u, u.NumOutputs)
}

// Log computes a natural logarithm.
func (u *Ugen) Log() Input {
	return unaryOpLog(u.Rate, u, u.NumOutputs)
}

// Log10 computes a base 10 logarithm.
func (u *Ugen) Log10() Input {
	return unaryOpLog10(u.Rate, u, u.NumOutputs)
}

// Log2 computes a base 2 logarithm.
func (u *Ugen) Log2() Input {
	return unaryOpLog2(u.Rate, u, u.NumOutputs)
}

// Max computes the maximum of one Input and another.
func (u *Ugen) Max(other Input) Input {
	return binOpMax(u.Rate, u, other, u.NumOutputs)
}

// Midicps converts MIDI note number to cycles per second.
func (u *Ugen) Midicps() Input {
	return unaryOpMidicps(u.Rate, u, u.NumOutputs)
}

// Midiratio converts an interval in MIDI notes into a frequency ratio.
func (u *Ugen) Midiratio() Input {
	return unaryOpMidiratio(u.Rate, u, u.NumOutputs)
}

// Min returns the minimum of one signal and another.
func (u *Ugen) Min(other Input) Input {
	return binOpMin(u.Rate, u, other, u.NumOutputs)
}

// Moddif returns the smaller of the great circle distances between the two points.
func (u *Ugen) Moddif(y, mod Input) Input {
	return moddif(u.Rate, u, y, mod, u.NumOutputs)
}

// Modulo computes the modulo of one signal and another.
func (u *Ugen) Modulo(val Input) Input {
	return binOpModulo(u.Rate, u, val, u.NumOutputs)
}

// Mul multiplies the ugen node by an input.
func (u *Ugen) Mul(val Input) Input {
	return binOpMul(u.Rate, u, val, u.NumOutputs)
}

// MulAdd multiplies and adds inputs to a ugen node.
func (u *Ugen) MulAdd(mul, add Input) Input {
	return mulAdd(u.Rate, u, mul, add, u.NumOutputs)
}

// Neg is a convenience operator that multiplies a signal by -1.
func (u *Ugen) Neg() Input {
	return unaryOpNeg(u.Rate, u, u.NumOutputs)
}

// Octcps converts decimal octaves to cycles per second.
func (u *Ugen) Octcps() Input {
	return unaryOpOctcps(u.Rate, u, u.NumOutputs)
}

// Pow raises this input to the power of another.
func (u *Ugen) Pow(val Input) Input {
	return binOpPow(u.Rate, u, val, u.NumOutputs)
}

// Rand returns an evenly distributed random value between this and zero.
func (u *Ugen) Rand() Input {
	return unaryOpRand(u.Rate, u, u.NumOutputs)
}

// Rand2 returns an evenly distributed random value between [+this ... - this].
func (u *Ugen) Rand2() Input {
	return unaryOpRand2(u.Rate, u, u.NumOutputs)
}

// Ratiomidi converts a frequency ratio to an interval in MIDI notes.
func (u *Ugen) Ratiomidi() Input {
	return unaryOpRatiomidi(u.Rate, u, u.NumOutputs)
}

// Reciprocal computes the reciprocal of a signal.
func (u *Ugen) Reciprocal() Input {
	return unaryOpReciprocal(u.Rate, u, u.NumOutputs)
}

// Ring1 is ring modulation plus first source.
func (u *Ugen) Ring1(val Input) Input {
	return binOpRing1(u.Rate, u, val, u.NumOutputs)
}

// Ring2 is ring modulation plus both sources.
func (u *Ugen) Ring2(val Input) Input {
	return binOpRing2(u.Rate, u, val, u.NumOutputs)
}

// Ring3 returns the value of (a*a *b)
func (u *Ugen) Ring3(val Input) Input {
	return binOpRing3(u.Rate, u, val, u.NumOutputs)
}

// Ring4 returns the value of ((a*a *b) - (a*b*b)).
func (u *Ugen) Ring4(val Input) Input {
	return binOpRing4(u.Rate, u, val, u.NumOutputs)
}

// Round performs quantization by rounding. Rounds a to the nearest multiple of b.
func (u *Ugen) Round(val Input) Input {
	return binOpRound(u.Rate, u, val, u.NumOutputs)
}

// Scaleneg returns a*b when a < 0, otherwise a.
func (u *Ugen) Scaleneg(val Input) Input {
	return binOpScaleneg(u.Rate, u, val, u.NumOutputs)
}

// Sign computes the sign of a signal.
func (u *Ugen) Sign() Input {
	return unaryOpSign(u.Rate, u, u.NumOutputs)
}

// Sin returns the sine of a ugen.
func (u *Ugen) Sin() Input {
	return unaryOpSin(u.Rate, u, u.NumOutputs)
}

// Sinh returns the hyperbolic sine of a ugen.
func (u *Ugen) Sinh() Input {
	return unaryOpSinh(u.Rate, u, u.NumOutputs)
}

// SoftClip adds distortion to a ugen.
func (u *Ugen) SoftClip() Input {
	return unaryOpSoftClip(u.Rate, u, u.NumOutputs)
}

// Sqrdif computes the square of the difference between the two inputs.
func (u *Ugen) Sqrdif(val Input) Input {
	return binOpSqrdif(u.Rate, u, val, u.NumOutputs)
}

// Sqrsum computes the square of the sum of the two inputs.
func (u *Ugen) Sqrsum(val Input) Input {
	return binOpSqrsum(u.Rate, u, val, u.NumOutputs)
}

// Sqrt computes the square root of a signal.
func (u *Ugen) Sqrt() Input {
	return unaryOpSqrt(u.Rate, u, u.NumOutputs)
}

// Squared computes the square of a signal.
func (u *Ugen) Squared() Input {
	return unaryOpSquared(u.Rate, u, u.NumOutputs)
}

// Sum3rand returns a value from a gaussian-like random distribution between in and zero.
func (u *Ugen) Sum3rand() Input {
	return unaryOpSum3rand(u.Rate, u, u.NumOutputs)
}

// Sumsqr returns the value of (a*a) + (b*b).
func (u *Ugen) Sumsqr(val Input) Input {
	return binOpSumsqr(u.Rate, u, val, u.NumOutputs)
}

// Tan returns the tangent of a ugen.
func (u *Ugen) Tan() Input {
	return unaryOpTan(u.Rate, u, u.NumOutputs)
}

// Tanh returns the hyperbolic tangent of a ugen.
func (u *Ugen) Tanh() Input {
	return unaryOpTanh(u.Rate, u, u.NumOutputs)
}

// Thresh returns 0 when a < b, otherwise a.
func (u *Ugen) Thresh(val Input) Input {
	return binOpThresh(u.Rate, u, val, u.NumOutputs)
}

// Trunc performs quantization by truncation. Truncate a to a multiple of b.
func (u *Ugen) Trunc(val Input) Input {
	return binOpTrunc(u.Rate, u, val, u.NumOutputs)
}

// Wrap2 wraps input wave to +/-b
func (u *Ugen) Wrap2(val Input) Input {
	return binOpWrap2(u.Rate, u, val, u.NumOutputs)
}

// Write writes a Ugen
func (u *Ugen) Write(w io.Writer) error {
	// write the synthdef name
	if err := newPstring(u.Name).Write(w); err != nil {
		return err
	}
	// write rate
	if err := binary.Write(w, byteOrder, u.Rate); err != nil {
		return err
	}
	// write inputs
	numInputs := int32(len(u.Inputs))
	if err := binary.Write(w, byteOrder, numInputs); err != nil {
		return err
	}
	// write outputs
	numOutputs := int32(len(u.Outputs))
	if err := binary.Write(w, byteOrder, numOutputs); err != nil {
		return err
	}
	// special index
	if err := binary.Write(w, byteOrder, u.SpecialIndex); err != nil {
		return err
	}
	// inputs
	for _, i := range u.Inputs {
		if err := i.Write(w); err != nil {
			return err
		}
	}
	// outputs
	for _, o := range u.Outputs {
		if err := o.Write(w); err != nil {
			return err
		}
	}
	return nil
}

func (u *Ugen) inputsOrdered() bool {
	if u.Name == BinOpUgenName {
		switch u.SpecialIndex {
		case BinOpAdd, BinOpGCD, BinOpLCM, BinOpMax, BinOpMin, BinOpMul, BinOpSqrsum, BinOpSumsqr:
			return false
		default:
			return true
		}
	}
	return true
}

// readUgen reads a ugen from an io.Reader
func readUgen(r io.Reader) (*Ugen, error) {
	var (
		numInputs    int32
		numOutputs   int32
		specialIndex int16
		rate         int8
	)
	// read name
	name, err := readPstring(r)
	if err != nil {
		return nil, err
	}
	// read calculation rate
	if err := binary.Read(r, byteOrder, &rate); err != nil {
		return nil, err
	}
	// read number of inputs
	if err := binary.Read(r, byteOrder, &numInputs); err != nil {
		return nil, err
	}
	// read number of outputs
	if err := binary.Read(r, byteOrder, &numOutputs); err != nil {
		return nil, err
	}
	// read special index
	if err := binary.Read(r, byteOrder, &specialIndex); err != nil {
		return nil, err
	}
	var (
		inputs  = make([]UgenInput, numInputs)
		outputs = make([]Output, numOutputs)
	)
	// read inputs
	for i := 0; int32(i) < numInputs; i++ {
		in, err := readInput(r)
		if err != nil {
			return nil, err
		}
		inputs[i] = in
	}
	// read outputs
	for i := 0; int32(i) < numOutputs; i++ {
		out, err := readOutput(r)
		if err != nil {
			return nil, err
		}
		outputs[i] = out
	}
	return &Ugen{
		Name:         name.String(),
		Rate:         rate,
		SpecialIndex: specialIndex,
		Inputs:       inputs,
		Outputs:      outputs,
		NumOutputs:   int(numOutputs),
	}, nil
}

// asOutput initializes the outputs array of the ugen node.
func asOutput(u *Ugen) *Ugen {
	if u.Outputs == nil {
		u.Outputs = make([]Output, u.NumOutputs)
		for i := range u.Outputs {
			u.Outputs[i] = Output(u.Rate)
		}
	}
	return u
}

func cloneUgen(v *Ugen) *Ugen {
	u := *v
	return &u
}
