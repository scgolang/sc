package ugens

import . "github.com/briansorahan/sc/types"

func BinOpMul(rate int8, x, y Input) *BaseNode {
	n := newNode("BinaryOpUGen", rate, 2)
	n.addInput(x)
	n.addInput(y)
	return n
}

func BinOpAdd(rate int8, x, y Input) *BaseNode {
	n := newNode("BinaryOpUGen", rate, 0)
	n.addInput(x)
	n.addInput(y)
	return n
}
