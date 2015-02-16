package ugens

import (
	"github.com/briansorahan/sc"
)

// parseNode gives ugens a way to parse the arguments they are given
// to populate the inputs array of a base node
type parseNode func(node *baseNode, args ...interface{})

type baseUgen struct {
	name string
	pn parseNode
}

func (self *baseUgen) Ar(args ...interface{}) sc.UgenNode {
	node := newNode(self.name, 2)
	self.pn(node, args...)
	return node
}

func (self *baseUgen) Kr(args ...interface{}) sc.UgenNode {
	node := newNode(self.name, 1)
	self.pn(node, args...)
	return node
}

func (self *baseUgen) Ir(args ...interface{}) sc.UgenNode {
	node := newNode(self.name, 0)
	self.pn(node, args...)
	return node
}

func newUgen(name string, pn parseNode) *baseUgen {
	base := baseUgen{name, pn}
	return &base
}
