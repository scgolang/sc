package ugens

import (
	. "github.com/briansorahan/sc/types"
)

const (
	// UGen done actions, see http://doc.sccode.org/Reference/UGen-doneActions.html
	DoNothing             = 0
	Pause                 = 1
	FreeEnclosing         = 2
	FreePreceding         = 3
	FreeFollowing         = 4
	FreePrecedingGroup    = 5
	FreeFollowingGroup    = 6
	FreeAllPreceding      = 7
	FreeAllFollowing      = 8
	FreeAndPausePreceding = 9
	FreeAndPauseFollowing = 10
	DeepFreePreceding     = 11
	DeepFreeFollowing     = 12
	FreeAllInGroup        = 13
	// I do not understand the difference between the last and
	// next-to-last options [bps]
)

// ugen node base type
type BaseNode struct {
	name         string
	rate         int8
	specialIndex int16
	inputs       []Input
	outputs      []Output
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

func (self *BaseNode) Inputs() []Input {
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

func (self *BaseNode) Mul(val Input) Input {
	return BinOpMul(self.rate, self, val)
}

func (self *BaseNode) Add(val Input) Input {
	return BinOpAdd(self.rate, self, val)
}

// addInput appends an Input to this node's list of inputs
func (self *BaseNode) addInput(in Input) {
	if node, isNode := in.(*BaseNode); isNode {
		node.IsOutput()
	}
	self.inputs = append(self.inputs, in)
}

// addInputs appends some Inputs to this node's list of inputs
func (self *BaseNode) addInputs(ins ...Input) {
	for _, in := range ins {
		self.addInput(in)
	}
}

// newNode is a factory function for creating new BaseNode instances
func newNode(name string, rate int8, specialIndex int16) *BaseNode {
	node := BaseNode{
		name,
		rate,
		specialIndex,
		make([]Input, 0),
		make([]Output, 0),
	}
	return &node
}
