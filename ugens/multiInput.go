package ugens

import . "github.com/briansorahan/sc/types"

type Inputs struct {
	inputs []Input
}

func (self *Inputs) Add(val Input) Input {
	l := len(self.inputs)
	ia := make([]Input, l)
	for i, in := range self.inputs {
		ia[i] = in.Add(val)
	}
	return &Inputs{ia}
}

func (self *Inputs) Mul(val Input) Input {
	l := len(self.inputs)
	ia := make([]Input, l)
	for i, in := range self.inputs {
		ia[i] = in.Mul(val)
	}
	return &Inputs{ia}
}

func Multi(inputs ...Input) *Inputs{
	return &Inputs{inputs}
}
