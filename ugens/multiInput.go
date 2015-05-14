package ugens

import . "github.com/scgolang/sc/types"

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

func (self *Inputs) MulAdd(mul, add Input) Input {
	l := len(self.inputs)
	ia := make([]Input, l)
	for i, in := range self.inputs {
		ia[i] = in.MulAdd(mul, add)
	}
	return &Inputs{ia}
}

func (self *Inputs) InputArray() []Input {
	return self.inputs
}

func Multi(inputs ...Input) *Inputs {
	return &Inputs{inputs}
}
