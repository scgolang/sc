package ugens

import (
	"fmt"
)

// Out write a signal to a bus
var Out = newUgen("Out", func(node *baseNode, args ...interface{}) {
	nargs := len(args)
	// parse arguments
	if nargs < 2 {
		panic(fmt.Errorf("Out expects at least 2 arguments, but was given %d", nargs))
	}
	if bus, isInt := args[0].(int); isInt {
		node.addConstantInput(float32(bus))
	} else {
		panic(fmt.Errorf("Out expects first argument to be int"))
	}
	for i := 1; i < nargs; i++ {
		arg := args[i]
		if in, isNode := arg.(*baseNode); isNode {
			node.addInput(in)
		} else {
			panic(fmt.Errorf("Out expects ugen arguments"))
		}
	}
})
