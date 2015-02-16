package ugens

import (
	"github.com/briansorahan/sc"
)

var Out = newOut()

// ugen implementation
type out struct {
}

func (self *out) Ar(args ...interface{}) sc.UgenNode {
	return newOutNode(2, args...)
}

func (self *out) Kr(args ...interface{}) sc.UgenNode {
	return newOutNode(1, args...)
}

func (self *out) Ir(args ...interface{}) sc.UgenNode {
	return newOutNode(0, args...)
}

func newOut() *out {
	o := out{}
	return &o
}

func newOutNode(rate int8, args ...interface{}) sc.UgenNode {
	node := newNode("Out", rate)
	// parse arguments
	return node
}
