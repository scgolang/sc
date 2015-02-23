package sc

import . "github.com/briansorahan/sc/types"

type ugenGraph struct {
	root UgenNode
}

func (self *ugenGraph) Root() UgenNode {
	return self.root
}

func newGraph(root UgenNode) *ugenGraph {
	return &ugenGraph{root}
}

func flatten(root UgenNode, def *synthdef) *input {
	stack := NewStack()
	inputs := root.Inputs()
	// iterate through ugen inputs in reverse order
	for i := len(inputs)-1; i >= 0; i-- {
		input := inputs[i]
		if input.IsConstant() {
			// push a float32
			stack.Push(input.(ConstantInput).Value())
		} else {
			ugenNode := input.(UgenInput).Value()
			// recurse with the next ugen as root
			// and push an *input
			stack.Push(flatten(ugenNode, def))
		}
	}

	// add inputs to root
	var in *input
	u := cloneUgen(root)
	for val := stack.Pop(); val != nil; val = stack.Pop() {
		if floatVal, isFloat := val.(float32); isFloat {
			in = def.AddConstant(floatVal)
		} else if inputVal, isInput := val.(*input); isInput {
			in = inputVal
		} else {
			panic("input was neither a float nor")
		}
		u.AppendInput(in)
	}

	return def.AddUgen(u)
}
