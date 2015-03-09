package ugens

import (
	. "github.com/briansorahan/sc/types"
)

// ugen node base type
type baseNode struct {
	name string
	rate int8
	specialIndex int16
	inputs []interface{}
	outputs []Output
}

func (self *baseNode) Name() string {
	return self.name
}

func (self *baseNode) Rate() int8 {
	return self.rate
}

func (self *baseNode) SpecialIndex() int16 {
	return self.specialIndex
}

func (self *baseNode) Inputs() []interface{} {
	return self.inputs
}

func (self *baseNode) Outputs() []Output {
	return self.outputs
}

func (self *baseNode) IsOutput() {
	if len(self.outputs) == 0 {
		self.outputs = append(self.outputs, output(self.rate))
	}
}

func (self *baseNode) Mul(val interface{}) UgenNode {
	node := newNode("BinaryOpUGen", self.rate, 2)
	node.addInput(self)
	node.addInput(val)
	return node
}

func (self *baseNode) Add(val interface{}) UgenNode {
	node := newNode("BinaryOpUGen", self.rate, 0)
	node.addInput(self)
	node.addInput(val)
	return node
}

// addInput appends an Input to this node's list of inputs
func (self *baseNode) addInput(in interface{}) {
	if node, isNode := in.(*baseNode); isNode {
		node.IsOutput()
	}
	self.inputs = append(self.inputs, in)
}

// newNode is a factory function for creating new baseNode instances
func newNode(name string, rate int8, specialIndex int16) *baseNode {
	node := baseNode{
		name,
		rate,
		specialIndex,
		make([]interface{}, 0),
		make([]Output, 0),
	}
	return &node
}
