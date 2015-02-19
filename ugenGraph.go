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

func flatten(root UgenNode, def *synthdef) {
	constants := NewStack()
	inputs := root.Inputs()

	for i := len(inputs)-1; i >= 0; i-- {
		input := inputs[i]
		if input.IsConstant() {
			constants.Push(input.(ConstantInput).Value())
		} else {
			// flatten(input.Value().(UgenNode), def)
			flatten(input.(UgenInput).Value(), def)
		}
	}

	for val := constants.Pop(); val != nil; val = constants.Pop() {
		def.AppendConstant(val.(float32))
	}
}
