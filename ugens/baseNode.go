package ugens

import (
	. "github.com/briansorahan/sc/types"
)

// ugen node base type
type BaseNode struct {
	name string
	rate int8
	inputs []Input
}

func (self *BaseNode) Name() string {
	return self.name
}

func (self *BaseNode) Rate() int8 {
	return self.rate
}

func (self *BaseNode) Inputs() []Input {
	return self.inputs
}

func (self *BaseNode) IsConstant() bool {
	return false
}

func (self *BaseNode) Value() UgenNode {
	return self
}

// addInput appends an Input to this node's list of inputs
func (self *BaseNode) addInput(in Input) {
	self.inputs = append(self.inputs, in)
}

// addConstantInput is a helper that wraps a float32 with
// the constantInput type (which implements the Input interface)
func (self *BaseNode) addConstantInput(val float32) {
	self.inputs = append(self.inputs, constantInput(val))
}

// newNode is a factory function for creating new BaseNode instances
func newNode(name string, rate int8) *BaseNode {
	node := BaseNode{
		name,
		rate,
		make([]Input, 0),
	}
	return &node
}
