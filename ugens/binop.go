package ugens

import . "github.com/scgolang/sc/types"

func BinOpMul(rate int8, x, y Input, numOutputs int) *UgenNode {
	CheckRate(rate)
	return NewUgenNode("BinaryOpUGen", rate, 2, numOutputs, x, y)
}

func BinOpAdd(rate int8, x, y Input, numOutputs int) *UgenNode {
	CheckRate(rate)
	return NewUgenNode("BinaryOpUGen", rate, 0, numOutputs, x, y)
}

func MulAdd(rate int8, in, mul, add Input, numOutputs int) *UgenNode {
	return NewUgenNode("MulAdd", rate, 0, numOutputs, in, mul, add)
}
