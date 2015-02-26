package ugens

import (
	. "github.com/briansorahan/sc/types"
)

// ugen node base type
type BaseNode struct {
	name string
	rate int8
	specialIndex int16
	inputs []interface{}
	outputs []Output
}

func (self *BaseNode) Name() string {
	return self.name
}

func (self *BaseNode) Rate() int8 {
	return self.rate
}

func (self *BaseNode) SpecialIndex() int16 {
	return self.specialIndex
}

func (self *BaseNode) Inputs() []interface{} {
	return self.inputs
}

func (self *BaseNode) Outputs() []Output {
	return self.outputs
}

func (self *BaseNode) IsOutput() {
	if len(self.outputs) == 0 {
		self.outputs = append(self.outputs, output(self.rate))
	}
}

func (self *BaseNode) Mul(f float32) UgenNode {
	if f == float32(1) {
		return self
	}
	node := newNode("BinaryOpUGen", self.rate, 2)
	node.addInput(self)
	node.addInput(f)
	self.IsOutput()
	return node
}

func (self *BaseNode) Add(f float32) UgenNode {
	if f == float32(0) {
		return self
	}
	node := newNode("BinaryOpUGen", self.rate, 0)
	node.addInput(self)
	node.addInput(f)
	self.IsOutput()
	return node
}

// addInput appends an Input to this node's list of inputs
func (self *BaseNode) addInput(in interface{}) {
	self.inputs = append(self.inputs, in)
}

// newNode is a factory function for creating new BaseNode instances
func newNode(name string, rate int8, specialIndex int16) *BaseNode {
	node := BaseNode{
		name,
		rate,
		specialIndex,
		make([]interface{}, 0),
		make([]Output, 0),
	}
	return &node
}
