package ugens

import . "github.com/briansorahan/sc/types"

func BinOpMul(rate int8, x, y Input) *Node {
	checkRate(rate)
	return NewNode("BinaryOpUGen", rate, 2, x, y)
}

func BinOpAdd(rate int8, x, y Input) *Node {
	checkRate(rate)
	return NewNode("BinaryOpUGen", rate, 0, x, y)
}
