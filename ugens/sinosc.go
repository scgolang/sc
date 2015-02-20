package ugens

var SinOsc = newUgen("SinOsc", func(node *BaseNode, args ...interface{}) {
	nargs := len(args)
	// default argument values
	defaultFreq := float32(440)
	defaultPhase := float32(0)
	// parse arguments
	switch (nargs) {
	case 0:
		node.addConstantInput(defaultFreq)
		node.addConstantInput(defaultPhase)
	case 1:
		getInputs(node, args...)
		node.addConstantInput(defaultPhase)
	default:
		getInputs(node, args...)
	}
})
