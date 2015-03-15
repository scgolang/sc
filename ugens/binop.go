package ugens

import . "github.com/briansorahan/sc/types"

func BinOpMul(rate int8, x, y Input) *BaseNode {
	return NewNode("BinaryOpUGen", rate, 2, x, y)
	// n.addInput(x)
	// n.addInput(y)
	// return n
}

func BinOpAdd(rate int8, x, y Input) *BaseNode {
	return NewNode("BinaryOpUGen", rate, 0, x, y)
	// n.addInput(x)
	// n.addInput(y)
	// return n
}
