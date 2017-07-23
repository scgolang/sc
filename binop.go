package sc

// Operator constants.
const (
	BinOpAdd          = 0
	BinOpMax          = 13
	BinOpMul          = 2
	UnaryOpAbs        = 5
	UnaryOpAmpDb      = 22
	UnaryOpCeil       = 8
	UnaryOpCpsmidi    = 18
	UnaryOpCpsoct     = 24
	UnaryOpCubed      = 13
	UnaryOpDbAmp      = 21
	UnaryOpExp        = 15
	UnaryOpFloor      = 9
	UnaryOpFrac       = 10
	UnaryOpMidicps    = 17
	UnaryOpMidiratio  = 19
	UnaryOpNeg        = 0
	UnaryOpOctcps     = 23
	UnaryOpRatiomidi  = 20
	UnaryOpReciprocal = 16
	UnaryOpSoftClip   = 43
	UnaryOpSign       = 11
	UnaryOpSquared    = 12
	UnaryOpSqrt       = 14
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

// unaryOpCeil computes the ceiling of a signal.
func unaryOpCeil(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpCeil, numOutputs, in)
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

// unaryOpReciprocal returns the reciprocal of an input signal.
func unaryOpReciprocal(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpReciprocal, numOutputs, in)
}

// unaryOpRatiomidi converts a frequency ratio to an interval in MIDI notes.
func unaryOpRatiomidi(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpRatiomidi, numOutputs, in)
}

// unaryOpSign computes the sign of a signal.
// This returns -1 when a < 0, +1 when a > 0, 0 when a is 0.
func unaryOpSign(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput(unaryOpUgenName, rate, UnaryOpSign, numOutputs, in)
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
