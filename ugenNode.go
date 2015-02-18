package sc

import . "github.com/briansorahan/sc/types"

type ugenNode struct {
	name string
	rate int8
	inputs []Input
	outputs []Output
}

func (self *ugenNode) Name() string {
	return self.name
}

func (self *ugenNode) Rate() int8 {
	return self.rate
}

func newNode(name string, rate int8) *ugenNode {
	node := ugenNode{name, rate, make([]Input, 0), make([]Output, 0)}
	return &node
}
