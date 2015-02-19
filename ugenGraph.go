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
			// drill down into the next ugen after ensuring that
			// it has an output that is used as the input to
			// this one
			u := input.(UgenInput).Value()
			u.EnsureOutput()
			flatten(u, def)
		}
	}

	for val := constants.Pop(); val != nil; val = constants.Pop() {
		def.AppendConstant(val.(float32))
	}

	def.AppendUgen(cloneUgen(root))
}
