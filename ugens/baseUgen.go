package ugens

import (
	. "github.com/briansorahan/sc/types"
)

const (
	InitializationRate = 0
	ControlRate        = 1
	AudioRate          = 2
)

// parseNode gives ugens a way to parse the arguments they are given
// to populate the inputs array of a base node
type parseNode func(node *BaseNode, args ...interface{})

type baseUgen struct {
	name string
	pn   parseNode
}

func (self *baseUgen) Ar(args ...interface{}) UgenNode {
	node := newNode(self.name, AudioRate, 0)
	self.pn(node, args...)
	return node
}

func (self *baseUgen) Kr(args ...interface{}) UgenNode {
	node := newNode(self.name, ControlRate, 0)
	self.pn(node, args...)
	return node
}

func (self *baseUgen) Ir(args ...interface{}) UgenNode {
	node := newNode(self.name, InitializationRate, 0)
	self.pn(node, args...)
	return node
}

func newUgen(name string, pn parseNode) *baseUgen {
	base := baseUgen{name, pn}
	return &base
}
