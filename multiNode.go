package sc

// MultiNode is a MultiInput that consists of ugens.
type MultiNode struct {
	nodes []*UgenNode
}

// InputArray returns the ugens as a slice of Inputs.
func (mn *MultiNode) InputArray() []Input {
	l := len(mn.nodes)
	inputs := make([]Input, l)
	for i, n := range mn.nodes {
		inputs[i] = n
	}
	return inputs
}

// Nodes returns the slice of Ugen nodes.
func (mn *MultiNode) Nodes() []*UgenNode {
	return mn.nodes
}

// Max returns the maximum of one Input and another.
func (mn *MultiNode) Max(other Input) Input {
	a := make([]*UgenNode, len(mn.nodes))
	for i, n := range mn.nodes {
		a[i] = BinOpMax(n.Rate(), n, other, n.numOutputs)
	}
	return &MultiNode{a}
}

// Mul multiplies all the ugens by an input.
func (mn *MultiNode) Mul(val Input) Input {
	a := make([]*UgenNode, len(mn.nodes))
	for i, n := range mn.nodes {
		a[i] = BinOpMul(n.Rate(), n, val, n.numOutputs)
	}
	return &MultiNode{a}
}

// Add adds an input to all the ugens.
func (mn *MultiNode) Add(val Input) Input {
	a := make([]*UgenNode, len(mn.nodes))
	for i, n := range mn.nodes {
		a[i] = BinOpAdd(n.Rate(), n, val, n.numOutputs)
	}
	return &MultiNode{a}
}

// MulAdd does both multiplication and addition on all the
// ugen nodes.
func (mn *MultiNode) MulAdd(mul, add Input) Input {
	a := make([]*UgenNode, len(mn.nodes))
	for i, n := range mn.nodes {
		a[i] = MulAdd(n.Rate(), n, mul, add, n.numOutputs)
	}
	return &MultiNode{a}
}

// SoftClip adds distortion to the MultiNode.
func (mn *MultiNode) SoftClip() Input {
	a := make([]*UgenNode, len(mn.nodes))
	for i, n := range mn.nodes {
		a[i] = UnaryOpSoftClip(n.Rate(), n, n.numOutputs)
	}
	return &MultiNode{a}
}

// NewMultiNode creates a new MultiNode.
func NewMultiNode(nodes ...*UgenNode) *MultiNode {
	return &MultiNode{nodes}
}
