package ugens

// getInput either returns a constant input or a ugen input
// by running some type assertions on the provided arg
// if the arg is neither of these, then it returns nil
func getInput(arg interface{}) interface{} {
	if nv, isNode := arg.(*BaseNode); isNode {
		nv.IsOutput()
	}
	return arg
}

// getInputs iterates through ugen arguments and
// adds inputs to a ugen node
func getInputs(node *BaseNode, args ...interface{}) {
	for _, arg := range args {
		node.addInput(getInput(arg))
	}
}
