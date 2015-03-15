package ugens

import . "github.com/briansorahan/sc/types"

type MultiNode struct {
	nodes []*Node
}

func (self *MultiNode) Nodes() []UgenNode {
	l := len(self.nodes)
	nodes := make([]UgenNode, l)
	for i, n := range self.nodes {
		nodes[i] = n
	}
	return nodes
}

// Input implementation
func (self *MultiNode) Add(val Input) Input {
	l := len(self.nodes)
	a := make([]*Node, l)
	for i, n := range self.nodes {
		a[i] = BinOpAdd(n.Rate(), n, val)
	}
	return &MultiNode{a}
}

// Input implementation
func (self *MultiNode) Mul(val Input) Input {
	l := len(self.nodes)
	a := make([]*Node, l)
	for i, n := range self.nodes {
		a[i] = BinOpMul(n.Rate(), n, val)
	}
	return &MultiNode{a}
}

func (self *MultiNode) IsMulti() bool {
	return true
}

func NewMultiNode(nodes ...*Node) *MultiNode {
	return &MultiNode{nodes}
}
