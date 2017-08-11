package sc

// Operator constants.
// These are used as the "special index" for BinaryOpUGen and UnaryOpUGen in synthdef files.
const (
	BinOpAbsdif       = 38
	BinOpAdd          = 0
	BinOpAmclip       = 40
	BinOpAtan2        = 22
	BinOpClip2        = 42
	BinOpDifsqr       = 34
	BinOpDiv          = 4
	BinOpExcess       = 43
	BinOpExpon        = 25
	BinOpFold2        = 44
	BinOpGCD          = 18
	BinOpGT           = 9
	BinOpGTE          = 11
	BinOpHypot        = 23
	BinOpHypotApx     = 24
	BinOpLCM          = 17
	BinOpLT           = 8
	BinOpLTE          = 10
	BinOpMax          = 13
	BinOpModulo       = 5
	BinOpMin          = 12
	BinOpMul          = 2
	BinOpPow          = 25
	BinOpRing1        = 30
	BinOpRing2        = 31
	BinOpRing3        = 32
	BinOpRing4        = 33
	BinOpRound        = 19
	BinOpScaleneg     = 41
	BinOpSqrdif       = 37
	BinOpSqrsum       = 36
	BinOpSumsqr       = 35
	BinOpThresh       = 39
	BinOpTrunc        = 21
	BinOpWrap2        = 45
	UnaryOpAbs        = 5
	UnaryOpAcos       = 32
	UnaryOpAmpDb      = 22
	UnaryOpAsin       = 31
	UnaryOpAtan       = 33
	UnaryOpBilinrand  = 40
	UnaryOpCeil       = 8
	UnaryOpCoin       = 44
	UnaryOpCos        = 29
	UnaryOpCosh       = 35
	UnaryOpCpsmidi    = 18
	UnaryOpCpsoct     = 24
	UnaryOpCubed      = 13
	UnaryOpDbAmp      = 21
	UnaryOpDistort    = 42
	UnaryOpExp        = 15
	UnaryOpFloor      = 9
	UnaryOpFrac       = 10
	UnaryOpLinrand    = 39
	UnaryOpLog        = 25
	UnaryOpLog10      = 27
	UnaryOpLog2       = 26
	UnaryOpMidicps    = 17
	UnaryOpMidiratio  = 19
	UnaryOpNeg        = 0
	UnaryOpOctcps     = 23
	UnaryOpRatiomidi  = 20
	UnaryOpRand       = 37
	UnaryOpRand2      = 38
	UnaryOpReciprocal = 16
	UnaryOpSoftClip   = 43
	UnaryOpSign       = 11
	UnaryOpSin        = 28
	UnaryOpSinh       = 34
	UnaryOpSquared    = 12
	UnaryOpSqrt       = 14
	UnaryOpSum3rand   = 41
	UnaryOpTan        = 30
	UnaryOpTanh       = 36
)

// Ugen names.
const (
	BinOpUgenName   = "BinaryOpUGen"
	UnaryOpUgenName = "UnaryOpUGen"
)

// binOpAbsdif returns the absolute value of the difference of two inputs.
func binOpAbsdif(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpAbsdif, numOutputs, x, y)
}

// binOpAdd creates a BinaryOpUGen that represents addition.
func binOpAdd(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpAdd, numOutputs, x, y)
}

// binOpAmclip returns 0 when b <= 0, a*b when b > 0.
func binOpAmclip(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpAmclip, numOutputs, x, y)
}

// binOpAtan2 returns the arctangent of y/x.
func binOpAtan2(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpAtan2, numOutputs, x, y)
}

// binOpClip2 clips input wave a to +/- b
func binOpClip2(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpClip2, numOutputs, x, y)
}

// binOpDifsqr computes a difference of squares.
func binOpDifsqr(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpDifsqr, numOutputs, x, y)
}

// binOpDiv divides one input by another.
func binOpDiv(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpDiv, numOutputs, x, y)
}

// binOpExcess returns the difference of the original signal and its clipped form: (a - clip2(a,b)).
func binOpExcess(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpExcess, numOutputs, x, y)
}

// binOpExpon creates a BinaryOpUgen that raises one Input to the power of another.
func binOpExpon(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpExpon, numOutputs, x, y)
}

// binOpFold2 folds input wave a to +/- b
func binOpFold2(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpFold2, numOutputs, x, y)
}

// binOpGCD computes the gcd of one Input and another.
func binOpGCD(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpGCD, numOutputs, x, y)
}

// binOpGT computes x > y.
func binOpGT(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpGT, numOutputs, x, y)
}

// binOpGTE computes x >= y.
func binOpGTE(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpGTE, numOutputs, x, y)
}

// binOpHypot returns the square root of the sum of the squares of x and y.
// Or equivalently, the distance from the origin to the point (x, y).
func binOpHypot(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpHypot, numOutputs, x, y)
}

// binOpHypotApx returns an approximation of the square root of the sum of the squares of x and y.
func binOpHypotApx(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpHypotApx, numOutputs, x, y)
}

// binOpLCM computes the lcm of one Input and another.
func binOpLCM(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpLCM, numOutputs, x, y)
}

// binOpLT computes x < y.
func binOpLT(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpLT, numOutputs, x, y)
}

// binOpLTE computes x <= y.
func binOpLTE(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpLTE, numOutputs, x, y)
}

// binOpMax creates a BinaryOpUgen that represents the maximum of two signals.
func binOpMax(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpMax, numOutputs, x, y)
}

// binOpMin creates a BinaryOpUgen that represents the minimum of two signals.
func binOpMin(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpMin, numOutputs, x, y)
}

// binOpModulo computes the modulo of one Input and another.
func binOpModulo(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpModulo, numOutputs, x, y)
}

// binOpMul creates a BinaryOpUGen that represents multiplication.
func binOpMul(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpMul, numOutputs, x, y)
}

// binOpPow creates a BinaryOpUgen that raises one Input to the power of another.
func binOpPow(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpPow, numOutputs, x, y)
}

// binOpRing1 creates a BinaryOpUgen that is ring modulation plus first source.
func binOpRing1(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpRing1, numOutputs, x, y)
}

// binOpRing2 creates a BinaryOpUgen that is ring modulation plus both sources.
func binOpRing2(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpRing2, numOutputs, x, y)
}

// binOpRing3 creates a BinaryOpUgen that returns the value of (a*a *b).
func binOpRing3(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpRing3, numOutputs, x, y)
}

// binOpRing4 creates a BinaryOpUgen that returns the value of ((a*a *b) - (a*b*b)).
func binOpRing4(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpRing4, numOutputs, x, y)
}

// binOpRound performs quantization by rounding. Rounds a to the nearest multiple of b.
func binOpRound(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpRound, numOutputs, x, y)
}

// binOpScaleneg returns a*b when a < 0, otherwise a.
func binOpScaleneg(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpScaleneg, numOutputs, x, y)
}

// binOpSqrdif computes the square of the difference of two inputs.
func binOpSqrdif(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpSqrdif, numOutputs, x, y)
}

// binOpSqrsum computes the square of the sum of two inputs.
func binOpSqrsum(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpSqrsum, numOutputs, x, y)
}

// binOpSumsqr computes a sum of squares.
func binOpSumsqr(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpSumsqr, numOutputs, x, y)
}

// binOpThresh performs thresholding. This returns 0 when x < y, otherwise x.
func binOpThresh(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpThresh, numOutputs, x, y)
}

// binOpTrunc performs quantization by truncation. Truncate a to a multiple of b.
func binOpTrunc(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpTrunc, numOutputs, x, y)
}

// binOpWrap2 wraps input wave to +/-b
func binOpWrap2(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(BinOpUgenName, rate, BinOpWrap2, numOutputs, x, y)
}

// mulAdd creates a MulAdd ugen.
func mulAdd(rate int8, in, mul, add Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput("MulAdd", rate, 0, numOutputs, in, mul, add)
}

// unaryOpAbs computes the absolute value of a signal.
func unaryOpAbs(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpAbs, numOutputs, in)
}

// unaryOpAcos computes the arccosine of a signal.
func unaryOpAcos(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpAcos, numOutputs, in)
}

// unaryOpAmpDb converts linear amplitude to decibels.
func unaryOpAmpDb(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpAmpDb, numOutputs, in)
}

// unaryOpAsin computes the arcsine of a signal.
func unaryOpAsin(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpAsin, numOutputs, in)
}

// unaryOpAtan computes the arctangent of a signal.
func unaryOpAtan(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpAtan, numOutputs, in)
}

// unaryOpBilinrand returns a linearly distributed random value between [+in ... -in].
func unaryOpBilinrand(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpBilinrand, numOutputs, in)
}

// unaryOpCeil computes the ceiling of a signal.
func unaryOpCeil(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpCeil, numOutputs, in)
}

// unaryOpCoin returns one or zero with the probability given by the argument.
func unaryOpCoin(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpCoin, numOutputs, in)
}

// unaryOpCos returns the cosine of an input.
func unaryOpCos(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpCos, numOutputs, in)
}

// unaryOpCosh returns the hyperbolic cosine of an input.
func unaryOpCosh(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpCosh, numOutputs, in)
}

// unaryOpCpsmidi converts frequency in Hz to midi note values.
func unaryOpCpsmidi(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpCpsmidi, numOutputs, in)
}

// unaryOpCpsoct converts cycles per second to decimal octaves.
func unaryOpCpsoct(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpCpsoct, numOutputs, in)
}

// unaryOpCubed computes the cube of a signal.
func unaryOpCubed(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpCubed, numOutputs, in)
}

// unaryOpDbAmp converts decibels to linear amplitude.
func unaryOpDbAmp(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpDbAmp, numOutputs, in)
}

// unaryOpDistort performs non-linear distortion on a signal.
func unaryOpDistort(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpDistort, numOutputs, in)
}

// unaryOpExp computes the exponential of a signal.
func unaryOpExp(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpExp, numOutputs, in)
}

// unaryOpFloor computes the floor (next lowest integer) of a signal.
func unaryOpFloor(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpFloor, numOutputs, in)
}

// unaryOpFrac returns the fractional part of a signal.
func unaryOpFrac(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpFrac, numOutputs, in)
}

// unaryOpLinrand returns a linearly distributed random value between in and zero.
func unaryOpLinrand(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpLinrand, numOutputs, in)
}

// unaryOpLog computes a natural logarithm.
func unaryOpLog(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpLog, numOutputs, in)
}

// unaryOpLog10 computes a base 10 logarithm.
func unaryOpLog10(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpLog10, numOutputs, in)
}

// unaryOpLog2 computes a base 2 logarithm.
func unaryOpLog2(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpLog2, numOutputs, in)
}

// unaryOpMidicps converts MIDI note numbers to cycles per second.
func unaryOpMidicps(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpMidicps, numOutputs, in)
}

// unaryOpMidiratio converts an interval in MIDI note numbers to a frequency ratio. TODO
func unaryOpMidiratio(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpMidiratio, numOutputs, in)
}

// unaryOpNeg multiplies a signal by -1.
func unaryOpNeg(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpNeg, numOutputs, in)
}

// unaryOpOctcps converts decimal octaves to cycles per second.
func unaryOpOctcps(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpOctcps, numOutputs, in)
}

// unaryOpRand returns an evenly distributed random value between in and zero.
func unaryOpRand(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpRand, numOutputs, in)
}

// unaryOpRand2 returns an evenly distributed random value between [+in ... -in].
func unaryOpRand2(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpRand2, numOutputs, in)
}

// unaryOpRatiomidi converts a frequency ratio to an interval in MIDI notes.
func unaryOpRatiomidi(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpRatiomidi, numOutputs, in)
}

// unaryOpReciprocal returns the reciprocal of an input signal.
func unaryOpReciprocal(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpReciprocal, numOutputs, in)
}

// unaryOpSign computes the sign of a signal.
// This returns -1 when a < 0, +1 when a > 0, 0 when a is 0.
func unaryOpSign(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpSign, numOutputs, in)
}

// unaryOpSin returns the sine of an input.
func unaryOpSin(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpSin, numOutputs, in)
}

// unaryOpSinh returns the hyperbolic sine of an input.
func unaryOpSinh(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpSinh, numOutputs, in)
}

// unaryOpSoftClip adds distortion to a ugen.
func unaryOpSoftClip(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpSoftClip, numOutputs, in)
}

// unaryOpSqrt computes the square root of a signal.
func unaryOpSqrt(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpSqrt, numOutputs, in)
}

// unaryOpSquared computes the square of a signal.
func unaryOpSquared(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpSquared, numOutputs, in)
}

// unaryOpSum3rand
func unaryOpSum3rand(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpSum3rand, numOutputs, in)
}

// unaryOpTan returns the tangent of an input.
func unaryOpTan(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpTan, numOutputs, in)
}

// unaryOpTanh returns the hyperbolic tangent of an input.
func unaryOpTanh(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(UnaryOpUgenName, rate, UnaryOpTanh, numOutputs, in)
}
