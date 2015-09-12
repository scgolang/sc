package ugens

import (
	. "github.com/scgolang/sc/types"
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
type UgenNode struct {
	name         string
	rate         int8
	specialIndex int16
	numOutputs   int
	inputs       []Input
	outputs      []Output
}

func (self *UgenNode) Name() string {
	return self.name
}

func (self *UgenNode) Rate() int8 {
	return self.rate
}

func (self *UgenNode) SpecialIndex() int16 {
	return self.specialIndex
}

func (self *UgenNode) Inputs() []Input {
	return self.inputs
}

func (self *UgenNode) Outputs() []Output {
	return self.outputs
}

func (self *UgenNode) IsOutput() {
	if self.outputs == nil {
		self.outputs = make([]Output, self.numOutputs)
		for i := range self.outputs {
			self.outputs[i] = output(self.rate)
		}
	}
}

func (self *UgenNode) Mul(val Input) Input {
	return BinOpMul(self.rate, self, val, self.numOutputs)
}

func (self *UgenNode) Add(val Input) Input {
	return BinOpAdd(self.rate, self, val, self.numOutputs)
}

func (self *UgenNode) MulAdd(mul, add Input) Input {
	return MulAdd(self.rate, self, mul, add, self.numOutputs)
}

// NewUgenNode is a factory function for creating new UgenNode instances.
// Panics if rate is not AR, KR, or IR.
// Panics if numOutputs <= 0.
func NewUgenNode(name string, rate int8, specialIndex int16, numOutputs int, inputs ...Input) *UgenNode {
	CheckRate(rate)
	if numOutputs <= 0 {
		panic("numOutputs must be a positive int")
	}
	n := new(UgenNode)
	n.name = name
	n.rate = rate
	n.specialIndex = specialIndex
	n.numOutputs = numOutputs
	n.inputs = make([]Input, len(inputs))

	// If any inputs are multi inputs, then this node
	// should get promoted to a multi node
	for i, input := range inputs {
		if node, isNode := input.(*UgenNode); isNode {
			node.IsOutput()
		}
		// add outputs to any nodes in a MultiInput
		if multi, isMulti := input.(MultiInput); isMulti {
			for _, in := range multi.InputArray() {
				if n, isn := in.(*UgenNode); isn {
					n.IsOutput()
				}
			}
		}
		n.inputs[i] = input
	}

	return n
}
