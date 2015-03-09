package ugens

import (
	. "github.com/briansorahan/sc/types"
)

var EnvGen = newEnvGen()

type envGen struct {
	name     string
	defaults []float32
}

func (self *envGen) Ar(args ...interface{}) UgenNode {
	return self.atRate(audioRate, args...)
}

func (self *envGen) Kr(args ...interface{}) UgenNode {
	return self.atRate(controlRate, args...)
}

func (self *envGen) Ir(args ...interface{}) UgenNode {
	return self.atRate(initializationRate, args...)
}

func (self *envGen) atRate(rate int8, args ...interface{}) UgenNode {
	var envArray []interface{}
	// sclang throws if the first arg is not an Env, so we do the same
	if len(args) == 0 {
		panic("EnvGen needs an Envelope as the first argument")
	}
	if envelope, isEnvelope := args[0].(envelope); isEnvelope {
		envArray = envelope.InputsArray()
	} else {
		panic("EnvGen needs an Envelope as the first argument")
	}
	node := newNode(self.name, rate, 0)
	withDefaults := applyDefaults(self.defaults, args[1:]...)
	inputs := append(withDefaults, envArray...)
	getInputs(node, inputs...)
	return node
}

func newEnvGen() *envGen {
	eg := envGen{"EnvGen", []float32{1, 1, 0, 1, 2}}
	return &eg
}
