package ugens

import (
	. "github.com/briansorahan/sc/types"
)

const (
	initializationRate = 0
	controlRate        = 1
	audioRate          = 2
)

type baseUgen struct {
	name     string
	defaults []float32
}

func (self *baseUgen) Ar(args ...interface{}) UgenNode {
	return self.atRate(audioRate, args...)
}

func (self *baseUgen) Kr(args ...interface{}) UgenNode {
	return self.atRate(controlRate, args...)
}

func (self *baseUgen) Ir(args ...interface{}) UgenNode {
	return self.atRate(initializationRate, args...)
}

func (self *baseUgen) atRate(rate int8, args ...interface{}) UgenNode {
	node := newNode(self.name, rate, 0)
	withDefaults := applyDefaults(self.defaults, args...)
	getInputs(node, withDefaults...)
	// getInputs(node, args...)
	return node
}

func newUgen(name string, defaults []float32) *baseUgen {
	base := baseUgen{name, defaults}
	return &base
}
