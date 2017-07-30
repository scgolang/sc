package sc

// Operator constants.
const (
	BinOpAdd          = 0
	BinOpDiv          = 4
	BinOpExpon        = 25
	BinOpGCD          = 18
	BinOpGT           = 9
	BinOpGTE          = 11
	BinOpHypot        = 23
	BinOpLCM          = 17
	BinOpLT           = 8
	BinOpLTE          = 10
	BinOpMax          = 13
	BinOpModulo       = 5
	BinOpMul          = 2
	BinOpPow          = 25
	BinOpRound        = 19
	BinOpTrunc        = 21
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

const (
	binopUgenName   = "BinaryOpUGen"
	unaryOpUgenName = "UnaryOpUGen"
)

// binOpAdd creates a BinaryOpUGen that represents addition.
func binOpAdd(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(binopUgenName, rate, BinOpAdd, numOutputs, x, y)
}

// binOpDiv divides one input by another.
func binOpDiv(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(binopUgenName, rate, BinOpDiv, numOutputs, x, y)
}

// binOpExpon creates a BinaryOpUgen that raises one Input to the power of another.
func binOpExpon(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(binopUgenName, rate, BinOpExpon, numOutputs, x, y)
}

// binOpGCD computes the gcd of one Input and another.
func binOpGCD(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(binopUgenName, rate, BinOpGCD, numOutputs, x, y)
}

// binOpGT computes x > y.
func binOpGT(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(binopUgenName, rate, BinOpGT, numOutputs, x, y)
}

// binOpGTE computes x >= y.
func binOpGTE(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(binopUgenName, rate, BinOpGTE, numOutputs, x, y)
}

// binOpHypot returns the square root of the sum of the squares of x and y.
// Or equivalently, the distance from the origin to the point (x, y).
func binOpHypot(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(binopUgenName, rate, BinOpHypot, numOutputs, x, y)
}

// binOpLCM computes the lcm of one Input and another.
func binOpLCM(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(binopUgenName, rate, BinOpLCM, numOutputs, x, y)
}

// binOpLT computes x < y.
func binOpLT(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(binopUgenName, rate, BinOpLT, numOutputs, x, y)
}

// binOpLTE computes x <= y.
func binOpLTE(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(binopUgenName, rate, BinOpLTE, numOutputs, x, y)
}

// binOpMax creates a BinaryOpUgen that represents the maximum of two signals.
func binOpMax(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(binopUgenName, rate, BinOpMax, numOutputs, x, y)
}

// binOpModulo computes the modulo of one Input and another.
func binOpModulo(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(binopUgenName, rate, BinOpModulo, numOutputs, x, y)
}

// binOpMul creates a BinaryOpUGen that represents multiplication.
func binOpMul(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(binopUgenName, rate, BinOpMul, numOutputs, x, y)
}

// binOpPow creates a BinaryOpUgen that raises one Input to the power of another.
func binOpPow(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(binopUgenName, rate, BinOpPow, numOutputs, x, y)
}

// binOpRound performs quantization by rounding. Rounds a to the nearest multiple of b.
func binOpRound(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(binopUgenName, rate, BinOpRound, numOutputs, x, y)
}

// binOpTrunc performs quantization by truncation. Truncate a to a multiple of b.
func binOpTrunc(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(binopUgenName, rate, BinOpTrunc, numOutputs, x, y)
}

// mulAdd creates a MulAdd ugen.
func mulAdd(rate int8, in, mul, add Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput("MulAdd", rate, 0, numOutputs, in, mul, add)
}

// unaryOpAbs computes the absolute value of a signal.
func unaryOpAbs(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpAbs, numOutputs, in)
}

// unaryOpAcos computes the arccosine of a signal.
func unaryOpAcos(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpAcos, numOutputs, in)
}

// unaryOpAmpDb converts linear amplitude to decibels.
func unaryOpAmpDb(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpAmpDb, numOutputs, in)
}

// unaryOpAsin computes the arcsine of a signal.
func unaryOpAsin(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpAsin, numOutputs, in)
}

// unaryOpAtan computes the arctangent of a signal.
func unaryOpAtan(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpAtan, numOutputs, in)
}

// unaryOpBilinrand returns a linearly distributed random value between [+in ... -in].
func unaryOpBilinrand(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpBilinrand, numOutputs, in)
}

// unaryOpCeil computes the ceiling of a signal.
func unaryOpCeil(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpCeil, numOutputs, in)
}

// unaryOpCoin returns one or zero with the probability given by the argument.
func unaryOpCoin(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpCoin, numOutputs, in)
}

// unaryOpCos returns the cosine of an input.
func unaryOpCos(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpCos, numOutputs, in)
}

// unaryOpCosh returns the hyperbolic cosine of an input.
func unaryOpCosh(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpCosh, numOutputs, in)
}

// unaryOpCpsmidi converts frequency in Hz to midi note values.
func unaryOpCpsmidi(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpCpsmidi, numOutputs, in)
}

// unaryOpCpsoct converts cycles per second to decimal octaves.
func unaryOpCpsoct(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpCpsoct, numOutputs, in)
}

// unaryOpCubed computes the cube of a signal.
func unaryOpCubed(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpCubed, numOutputs, in)
}

// unaryOpDbAmp converts decibels to linear amplitude.
func unaryOpDbAmp(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpDbAmp, numOutputs, in)
}

// unaryOpDistort performs non-linear distortion on a signal.
func unaryOpDistort(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpDistort, numOutputs, in)
}

// unaryOpExp computes the exponential of a signal.
func unaryOpExp(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpExp, numOutputs, in)
}

// unaryOpFloor computes the floor (next lowest integer) of a signal.
func unaryOpFloor(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpFloor, numOutputs, in)
}

// unaryOpFrac returns the fractional part of a signal.
func unaryOpFrac(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpFrac, numOutputs, in)
}

// unaryOpLinrand returns a linearly distributed random value between in and zero.
func unaryOpLinrand(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpLinrand, numOutputs, in)
}

// unaryOpLog computes a natural logarithm.
func unaryOpLog(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpLog, numOutputs, in)
}

// unaryOpLog10 computes a base 10 logarithm.
func unaryOpLog10(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpLog10, numOutputs, in)
}

// unaryOpLog2 computes a base 2 logarithm.
func unaryOpLog2(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpLog2, numOutputs, in)
}

// unaryOpMidicps converts MIDI note numbers to cycles per second.
func unaryOpMidicps(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpMidicps, numOutputs, in)
}

// unaryOpMidiratio converts an interval in MIDI note numbers to a frequency ratio. TODO
func unaryOpMidiratio(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpMidiratio, numOutputs, in)
}

// unaryOpNeg multiplies a signal by -1.
func unaryOpNeg(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpNeg, numOutputs, in)
}

// unaryOpOctcps converts decimal octaves to cycles per second.
func unaryOpOctcps(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpOctcps, numOutputs, in)
}

// unaryOpRand returns an evenly distributed random value between in and zero.
func unaryOpRand(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpRand, numOutputs, in)
}

// unaryOpRand2 returns an evenly distributed random value between [+in ... -in].
func unaryOpRand2(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpRand2, numOutputs, in)
}

// unaryOpRatiomidi converts a frequency ratio to an interval in MIDI notes.
func unaryOpRatiomidi(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpRatiomidi, numOutputs, in)
}

// unaryOpReciprocal returns the reciprocal of an input signal.
func unaryOpReciprocal(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpReciprocal, numOutputs, in)
}

// unaryOpSign computes the sign of a signal.
// This returns -1 when a < 0, +1 when a > 0, 0 when a is 0.
func unaryOpSign(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpSign, numOutputs, in)
}

// unaryOpSin returns the sine of an input.
func unaryOpSin(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpSin, numOutputs, in)
}

// unaryOpSinh returns the hyperbolic sine of an input.
func unaryOpSinh(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpSinh, numOutputs, in)
}

// unaryOpSoftClip adds distortion to a ugen.
func unaryOpSoftClip(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpSoftClip, numOutputs, in)
}

// unaryOpSqrt computes the square root of a signal.
func unaryOpSqrt(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpSqrt, numOutputs, in)
}

// unaryOpSquared computes the square of a signal.
func unaryOpSquared(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpSquared, numOutputs, in)
}

// unaryOpSum3rand
func unaryOpSum3rand(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpSum3rand, numOutputs, in)
}

// unaryOpTan returns the tangent of an input.
func unaryOpTan(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpTan, numOutputs, in)
}

// unaryOpTanh returns the hyperbolic tangent of an input.
func unaryOpTanh(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpTanh, numOutputs, in)
}
