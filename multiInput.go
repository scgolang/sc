package sc

type Inputs []Input

func (ins Inputs) Add(val Input) Input {
	l := len(ins)
	ia := make([]Input, l)
	for i, in := range ins {
		ia[i] = in.Add(val)
	}
	return Inputs(ia)
}

func (ins Inputs) Mul(val Input) Input {
	l := len(ins)
	ia := make([]Input, l)
	for i, in := range ins {
		ia[i] = in.Mul(val)
	}
	return Inputs(ia)
}

func (ins Inputs) MulAdd(mul, add Input) Input {
	l := len(ins)
	ia := make([]Input, l)
	for i, in := range ins {
		ia[i] = in.MulAdd(mul, add)
	}
	return Inputs(ia)
}

func (ins Inputs) InputArray() []Input {
	return ins
}

func Multi(inputs ...Input) Inputs {
	return Inputs(inputs)
}
