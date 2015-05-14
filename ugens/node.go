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
type Node struct {
	name         string
	rate         int8
	specialIndex int16
	inputs       []Input
	outputs      []Output
}

func (self *Node) Name() string {
	return self.name
}

func (self *Node) Rate() int8 {
	return self.rate
}

func (self *Node) SpecialIndex() int16 {
	return self.specialIndex
}

func (self *Node) Inputs() []Input {
	return self.inputs
}

func (self *Node) Outputs() []Output {
	return self.outputs
}

func (self *Node) IsOutput() {
	if len(self.outputs) == 0 {
		self.outputs = append(self.outputs, output(self.rate))
	}
}

func (self *Node) Mul(val Input) Input {
	return BinOpMul(self.rate, self, val)
}

func (self *Node) Add(val Input) Input {
	return BinOpAdd(self.rate, self, val)
}

func (self *Node) MulAdd(mul, add Input) Input {
	return MulAdd(self.rate, self, mul, add)
}

// NewNode is a factory function for creating new Node instances
func NewNode(name string, rate int8, specialIndex int16, inputs ...Input) *Node {
	n := new(Node)
	n.name = name
	n.rate = rate
	n.specialIndex = specialIndex
	n.inputs = make([]Input, len(inputs))
	n.outputs = make([]Output, 0)

	// If any inputs are multi inputs, then this node
	// should get promoted to a multi node
	for i, input := range inputs {
		if node, isNode := input.(*Node); isNode {
			node.IsOutput()
		}
		// add outputs to any nodes in a MultiInput
		if multi, isMulti := input.(MultiInput); isMulti {
			for _, in := range multi.InputArray() {
				if n, isn := in.(*Node); isn {
					n.IsOutput()
				}
			}
		}
		n.inputs[i] = input
	}

	return n
}
