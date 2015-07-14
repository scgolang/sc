package ugens

import . "github.com/scgolang/sc/types"

func BinOpMul(rate int8, x, y Input) *UgenNode {
	checkRate(rate)
	return NewUgenNode("BinaryOpUGen", rate, 2, 1, x, y)
}

func BinOpAdd(rate int8, x, y Input) *UgenNode {
	checkRate(rate)
	return NewUgenNode("BinaryOpUGen", rate, 0, 1, x, y)
}

func MulAdd(rate int8, in, mul, add Input) *UgenNode {
	return NewUgenNode("MulAdd", rate, 0, 1, in, mul, add)
}
