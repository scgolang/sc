package sc

// UGen done actions, see http://doc.sccode.org/Reference/UGen-doneActions.html
const (
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

// UgenNode represents a node in a ugen graph.
type UgenNode struct {
	name         string
	rate         int8
	specialIndex int16
	numOutputs   int
	inputs       []Input
	outputs      []Output
}

// Name returns the name of the ugen node.
func (un *UgenNode) Name() string {
	return un.name
}

// Rate returns the rate of the ugen node.
func (un *UgenNode) Rate() int8 {
	return un.rate
}

// SpecialIndex returns the special index of the ugen node.
func (un *UgenNode) SpecialIndex() int16 {
	return un.specialIndex
}

// Inputs returns the inputs of the ugen node.
func (un *UgenNode) Inputs() []Input {
	return un.inputs
}

// Outputs returns the outputs of the ugen node.
func (un *UgenNode) Outputs() []Output {
	return un.outputs
}

// IsOutput initializes the outputs array of the ugen node.
func (un *UgenNode) IsOutput() {
	if un.outputs == nil {
		un.outputs = make([]Output, un.numOutputs)
		for i := range un.outputs {
			un.outputs[i] = Output(un.rate)
		}
	}
}

// Mul multiplies the ugen node by an input.
func (un *UgenNode) Mul(val Input) Input {
	return BinOpMul(un.rate, un, val, un.numOutputs)
}

// Add adds an input to a ugen node.
func (un *UgenNode) Add(val Input) Input {
	return BinOpAdd(un.rate, un, val, un.numOutputs)
}

// MulAdd multiplies and adds inputs to a ugen node.
func (un *UgenNode) MulAdd(mul, add Input) Input {
	return MulAdd(un.rate, un, mul, add, un.numOutputs)
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
