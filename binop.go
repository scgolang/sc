package sc

// BinOpMul creates a BinaryOpUGen that represents multiplication.
func BinOpMul(rate int8, x, y Input, numOutputs int) *UgenNode {
	CheckRate(rate)
	return NewUgenNode("BinaryOpUGen", rate, 2, numOutputs, x, y)
}

// BinOpAdd creates a BinaryOpUGen that represents addition.
func BinOpAdd(rate int8, x, y Input, numOutputs int) *UgenNode {
	CheckRate(rate)
	return NewUgenNode("BinaryOpUGen", rate, 0, numOutputs, x, y)
}

// MulAdd creates a MulAdd ugen.
func MulAdd(rate int8, in, mul, add Input, numOutputs int) *UgenNode {
	return NewUgenNode("MulAdd", rate, 0, numOutputs, in, mul, add)
}
