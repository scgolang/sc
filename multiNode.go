package sc

type MultiNode struct {
	nodes []*UgenNode
}

func (mn *MultiNode) InputArray() []Input {
	l := len(mn.nodes)
	inputs := make([]Input, l)
	for i, n := range mn.nodes {
		inputs[i] = n
	}
	return inputs
}

func (mn *MultiNode) Nodes() []*UgenNode {
	return mn.nodes
}

// Input implementation
func (mn *MultiNode) Mul(val Input) Input {
	l := len(mn.nodes)
	a := make([]*UgenNode, l)
	for i, n := range mn.nodes {
		a[i] = BinOpMul(n.Rate(), n, val, n.numOutputs)
	}
	return &MultiNode{a}
}

// Input implementation
func (mn *MultiNode) Add(val Input) Input {
	l := len(mn.nodes)
	a := make([]*UgenNode, l)
	for i, n := range mn.nodes {
		a[i] = BinOpAdd(n.Rate(), n, val, n.numOutputs)
	}
	return &MultiNode{a}
}

func (mn *MultiNode) MulAdd(mul, add Input) Input {
	l := len(mn.nodes)
	a := make([]*UgenNode, l)
	for i, n := range mn.nodes {
		a[i] = MulAdd(n.Rate(), n, mul, add, n.numOutputs)
	}
	return &MultiNode{a}
}

func NewMultiNode(nodes ...*UgenNode) *MultiNode {
	return &MultiNode{nodes}
}
