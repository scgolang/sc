package sc

// Operator constants.
const (
	BinOpAdd          = 0
	UnaryOpNeg        = 0
	BinOpMul          = 2
	BinOpMax          = 13
	UnaryOpMidicps    = 17
	UnaryOpReciprocal = 16
	UnaryOpSoftClip   = 43
)

// binOpMax creates a BinaryOpUgen that represents the maximum of two signals.
func binOpMax(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput("BinaryOpUGen", rate, BinOpMax, numOutputs, x, y)
}

// binOpMul creates a BinaryOpUGen that represents multiplication.
func binOpMul(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput("BinaryOpUGen", rate, BinOpMul, numOutputs, x, y)
}

// binOpAdd creates a BinaryOpUGen that represents addition.
func binOpAdd(rate int8, x, y Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput("BinaryOpUGen", rate, BinOpAdd, numOutputs, x, y)
}

// mulAdd creates a MulAdd ugen.
func mulAdd(rate int8, in, mul, add Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput("MulAdd", rate, 0, numOutputs, in, mul, add)
}

// unaryOpMidicps converts MIDI note numbers to cycles per second.
func unaryOpMidicps(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput("UnaryOpUGen", rate, UnaryOpMidicps, numOutputs, in)
}

// unaryOpNeg multiplies a signal by -1.
func unaryOpNeg(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput("UnaryOpUGen", rate, UnaryOpNeg, numOutputs, in)
}

// unaryOpReciprocal returns the reciprocal of an input signal.
func unaryOpReciprocal(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput("UnaryOpUGen", rate, UnaryOpReciprocal, numOutputs, in)
}

// unaryOpSoftClip adds distortion to a ugen.
func unaryOpSoftClip(rate int8, in Input, numOutputs int) Input {
	CheckRate(rate)
	return NewInput("UnaryOpUGen", rate, UnaryOpSoftClip, numOutputs, in)
}
