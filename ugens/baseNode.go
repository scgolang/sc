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

// addInput appends an Input to this node's list of inputs
func (self *baseNode) addInput(in sc.Input) {
	self.inputs = append(self.inputs, in)
}

// addConstantInput is a helper that wraps a float32 with
// the constantInput type (which implements the Input interface)
func (self *baseNode) addConstantInput(val float32) {
	self.inputs = append(self.inputs, constantInput(val))
}

// ensureOutput ensures that a ugen node has an output with the
// given rate
func (self *baseNode) ensureOutput() {
	numOutputs := len(self.outputs)
	if numOutputs == 0 {
		self.outputs = append(self.outputs, output(self.Rate()))
	}
}

// newNode is a factory function for creating new baseNode instances
func newNode(name string, rate int8) *baseNode {
	node := baseNode{
		name,
		rate,
		make([]sc.Input, 0),
		make([]sc.Output, 0),
	}
	return &node
}
