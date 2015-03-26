package ugens

import . "github.com/briansorahan/sc/types"

// UgenInput creates a ugen suitable for use as an input to other ugens.
// It will return either a single-channel ugen or a multi-channel ugen.
func UgenInput(name string, rate int8, specialIndex int16, inputs ...Input) Input {
	expanded := expand(inputs...)
	l := len(expanded)
	if l == 1 {
		return NewNode(name, rate, specialIndex, inputs...)
	}
	// return MultiNode
	a := make([]*Node, len(expanded))
	for i := range a {
		a[i] = NewNode(name, rate, specialIndex, expanded[i]...)
	}
	return NewMultiNode(a...)
}

func expand(inputs ...Input) [][]Input {
	// first pass to determine how large each array needs to be
	// this could probably be more efficient but it doesn't matter
	sz := 0
	for _, in := range inputs {
		if multi, isMulti := in.(MultiInput); isMulti {
			ins := multi.Inputs()
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
		brr := make([]Input, n)

		for j := range brr {
			in := inputs[j]

			if multi, isMulti := in.(MultiInput); isMulti {
				ins := multi.Inputs()
				brr[j] = ins[j % len(ins)]
			} else {
				brr[j] = in
			}
		}

		arr[i] = brr
	}

	return arr
}
