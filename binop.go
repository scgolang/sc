package sc

// Operator constants.
const (
	BinOpAdd          = 0
	BinOpMax          = 13
	BinOpMul          = 2
	UnaryOpAbs        = 5
	UnaryOpAmpDb      = 22
	UnaryOpBilinrand  = 40
	UnaryOpCeil       = 8
	UnaryOpCoin       = 44
	UnaryOpCos        = 29
	UnaryOpCpsmidi    = 18
	UnaryOpCpsoct     = 24
	UnaryOpCubed      = 13
	UnaryOpDbAmp      = 21
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
	UnaryOpSquared    = 12
	UnaryOpSqrt       = 14
	UnaryOpSum3rand   = 41
	UnaryOpTan        = 30
)

const (
	binopUgenName   = "BinaryOpUGen"
	unaryOpUgenName = "UnaryOpUGen"
)

// binOpMax creates a BinaryOpUgen that represents the maximum of two signals.
func binOpMax(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(binopUgenName, rate, BinOpMax, numOutputs, x, y)
}

// binOpMul creates a BinaryOpUGen that represents multiplication.
func binOpMul(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(binopUgenName, rate, BinOpMul, numOutputs, x, y)
}

// binOpAdd creates a BinaryOpUGen that represents addition.
func binOpAdd(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(binopUgenName, rate, BinOpAdd, numOutputs, x, y)
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

// unaryOpAmpDb converts linear amplitude to decibels.
func unaryOpAmpDb(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpAmpDb, numOutputs, in)
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
