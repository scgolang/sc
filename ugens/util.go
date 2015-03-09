package ugens

// getInputs iterates through ugen arguments and
// adds inputs to a ugen node
func getInputs(node *BaseNode, args ...interface{}) {
	for _, arg := range args {
		node.addInput(arg)
	}
}
