package ugens

import (
	. "github.com/scgolang/sc/types"
)

type C float32

func (c C) Val() float32 {
	return float32(c)
}

func (c C) Mul(val Input) Input {
	switch v := val.(type) {
	case *UgenNode:
		return v.Mul(c)
	case C:
		return C(v * c)
	default:
		panic("input was neither ugen nor constant")
	}
}

func (c C) Add(val Input) Input {
	switch v := val.(type) {
	case *UgenNode:
		return v.Add(c)
	case C:
		return C(v + c)
	default:
		panic("input was neither ugen nor constant")
	}
}

func (c C) MulAdd(mul, add Input) Input {
	switch v := mul.(type) {
	case *UgenNode:
		return v.MulAdd(c, add)
	case C:
		switch w := add.(type) {
		case *UgenNode:
			// FIXME
			return w.MulAdd(c, mul)
		case C:
			return C(v*c + w)
		default:
			panic("input was neither ugen nor constant")
		}
	default:
		panic("input was neither ugen nor constant")
	}
}

func (c C) Equals(val C) bool {
	return float32(c) == float32(val)
}
