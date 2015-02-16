package ugens

import (
	"github.com/briansorahan/sc"
)

// ugen node base type
type baseNode struct {
	name string
	rate int8
	inputs []sc.Input
	outputs []sc.Output
}

// public interface

func (self *baseNode) Name() string {
	return self.name
}

func (self *baseNode) Rate() int8 {
	return self.rate
}

func (self *baseNode) Inputs() []sc.Input {
	return self.inputs
}

func (self *baseNode) Outputs() []sc.Output {
	return self.outputs
}

func (self *baseNode) IsConstant() bool {
	return false
}

func (self *baseNode) Value() interface{} {
	return *self
}

// private methods

func (self *baseNode) addConstantInput(val float32) {
	self.inputs = append(self.inputs, newConstantInput(val))
}

// factory

func newNode(name string, rate int8) *baseNode {
	node := baseNode{
		name,
		rate,
		make([]sc.Input, 0),
		make([]sc.Output, 0),
	}
	return &node
}
