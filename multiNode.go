package sc


type MultiNode struct {
	nodes []*UgenNode
}

func (self *MultiNode) InputArray() []Input {
	l := len(self.nodes)
	inputs := make([]Input, l)
	for i, n := range self.nodes {
		inputs[i] = n
	}
	return inputs
}

func (self *MultiNode) Nodes() []*UgenNode {
	return self.nodes
}

// Input implementation
func (self *MultiNode) Mul(val Input) Input {
	l := len(self.nodes)
	a := make([]*UgenNode, l)
	for i, n := range self.nodes {
		a[i] = BinOpMul(n.Rate(), n, val, n.numOutputs)
	}
	return &MultiNode{a}
}

// Input implementation
func (self *MultiNode) Add(val Input) Input {
	l := len(self.nodes)
	a := make([]*UgenNode, l)
	for i, n := range self.nodes {
		a[i] = BinOpAdd(n.Rate(), n, val, n.numOutputs)
	}
	return &MultiNode{a}
}

func (self *MultiNode) MulAdd(mul, add Input) Input {
	l := len(self.nodes)
	a := make([]*UgenNode, l)
	for i, n := range self.nodes {
		a[i] = MulAdd(n.Rate(), n, mul, add, n.numOutputs)
	}
	return &MultiNode{a}
}

func NewMultiNode(nodes ...*UgenNode) *MultiNode {
	return &MultiNode{nodes}
}
