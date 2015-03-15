package ugens

import . "github.com/briansorahan/sc/types"

type MultiInput struct {
	inputs []Input
}

func (self *MultiInput) Add(val Input) Input {
	l := len(self.inputs)
	ia := make([]Input, l)
	for i, in := range self.inputs {
		ia[i] = in.Add(val)
	}
	return &MultiInput{ia}
}

func (self *MultiInput) Mul(val Input) Input {
	l := len(self.inputs)
	ia := make([]Input, l)
	for i, in := range self.inputs {
		ia[i] = in.Mul(val)
	}
	return &MultiInput{ia}
}

func (self *MultiInput) IsMulti() bool {
	return true
}

func Multi(inputs ...Input) *MultiInput {
	return &MultiInput{inputs}
}
