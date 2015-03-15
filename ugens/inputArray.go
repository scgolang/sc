package ugens

import . "github.com/briansorahan/sc/types"

type InputArray []Input

func (self InputArray) Add(val Input) Input {
	l := len(self)
	ia := make([]Input, l)
	for i, in := range self {
		ia[i] = in.Add(val)
	}
	return InputArray(ia)
}

func (self InputArray) Mul(val Input) Input {
	l := len(self)
	ia := make([]Input, l)
	for i, in := range self {
		ia[i] = in.Mul(val)
	}
	return InputArray(ia)
}
