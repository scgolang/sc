package ugens

import (
	"fmt"
	"github.com/briansorahan/sc"
)

var SinOsc = newUgen("SinOsc", func(node *baseNode, args ...interface{}) {
	var in sc.Input
	nargs := len(args)
	defaultFreq := float32(440)
	defaultPhase := float32(0)
	switch (nargs) {
	case 0:
		node.addConstantInput(defaultFreq)
		node.addConstantInput(defaultPhase)
	case 1:
		in = getInput(args[0])
		if in == nil {
			panic(fmt.Errorf("SinOsc.Ar argument %d neither a constant nor a ugen", 1))
		}
		node.addInput(in)
		node.addConstantInput(defaultPhase)
	case 2:
		in = getInput(args[0])
		if in == nil {
			panic(fmt.Errorf("SinOsc.Ar argument %d was neither a constant nor a ugen", 1))
		}
		node.addInput(in)
	}
})
