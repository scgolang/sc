package ugens

func applyDefaults(defaults []float32, args ...interface{}) []interface{} {
	var numVals int
	numArgs := len(args)
	numDefaults := len(defaults)
	if numArgs >= numDefaults {
		numVals = numArgs
	} else {
		numVals = numDefaults
	}
	vals := make([]interface{}, numVals)
	for i := 0; i < numVals; i++ {
		if i < numArgs {
			vals[i] = args[i]
		} else {
			vals[i] = defaults[i]
		}
	}
	return vals
}
