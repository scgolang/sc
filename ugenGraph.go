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
			stack.Push(input.(ConstantInput).Value())
		} else {
			// drill down into the next ugen
			ugenNode := input.(UgenInput).Value()
			stack.Push(ugenNode)
		}
	}

	// add inputs to root
	var in *input
	u := cloneUgen(root)
	for val := stack.Pop(); val != nil; val = stack.Pop() {
		if floatVal, isFloat := val.(float32); isFloat {
			in = def.AddConstant(floatVal)
		} else if nodeVal, isNode := val.(UgenNode); isNode {
			in = flatten(nodeVal, def)
		}
		u.AppendInput(in)
	}

	return def.AddUgen(u)
}
