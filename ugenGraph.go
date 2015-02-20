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

func flatten(root UgenNode, def *synthdef) *ugen {
	constants := NewStack()
	inputs := root.Inputs()

	// iterate through ugen inputs in reverse order
	for i := len(inputs)-1; i >= 0; i-- {
		input := inputs[i]
		if input.IsConstant() {
			constants.Push(input.(ConstantInput).Value())
		} else {
			// drill down into the next ugen after ensuring that
			// it has an output that is used as the input to
			// this one
			un := input.(UgenInput).Value()
			u := cloneUgen(un)
			u.AddOutput()
			def.AddUgen(flatten(un, def))
		}
	}

	u := cloneUgen(root)

	// need to add inputs to u

	for val := constants.Pop(); val != nil; val = constants.Pop() {
		def.AddConstant(val.(float32))
	}

	def.AddUgen(u)

	return u
}
