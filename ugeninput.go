package sc

// UgenInput creates a ugen suitable for use as an input to other ugens.
// It will return either a single-channel ugen or a multi-channel ugen.
func UgenInput(name string, rate int8, specialIndex int16, numOutputs int, inputs ...Input) Input {
	expanded := expandInputs(inputs...)
	l := len(expanded)
	if l == 1 {
		return NewUgenNode(name, rate, specialIndex, numOutputs, inputs...)
	}
	// return MultiNode
	a := make([]*UgenNode, len(expanded))
	for i := range a {
		a[i] = NewUgenNode(name, rate, specialIndex, numOutputs, expanded[i]...)
	}
	return NewMultiNode(a...)
}

// expandInputs turns an array of inputs into a 2-dimensional array
// of inputs where the 1st dimension is the channel and
// the second is the array of inputs for each channel.
func expandInputs(inputs ...Input) [][]Input {
	// first pass to determine how large each array needs to be
	// this could probably be more efficient
	sz := 0
	for _, in := range inputs {
		if multi, isMulti := in.(MultiInput); isMulti {
			ins := multi.InputArray()
			l := len(ins)
			if l > sz {
				sz = l
			}
		}
	}
	if sz == 0 {
		// none were multi-channel inputs
		return [][]Input{inputs}
	}

	n := len(inputs)
	arr := make([][]Input, sz)

	for i := range arr {
		arr[i] = make([]Input, n)

		for j := range arr[i] {
			in := inputs[j]

			if multi, isMulti := in.(MultiInput); isMulti {
				ins := multi.InputArray()
				arr[i][j] = ins[i%len(ins)]
			} else {
				arr[i][j] = in
			}
		}
	}

	return arr
}
