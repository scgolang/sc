package sc

// C wraps a float32 and implements the Input interface.
type C float32

// Mul multiplies the constant by another input.
func (c C) Mul(val Input) Input {
	if v, ok := val.(C); ok {
		return C(float32(v) * float32(c))
	}
	return val.Mul(c)
}

// Add adds another input to the constant.
func (c C) Add(val Input) Input {
	if v, ok := val.(C); ok {
		return C(float32(v) + float32(c))
	}
	return val.Add(c)
}

// MulAdd multiplies and adds at the same time.
func (c C) MulAdd(mul, add Input) Input {
	if m, mok := mul.(C); mok {
		if a, aok := add.(C); aok {
			return C((float32(m) * float32(c)) + float32(a))
		}
		return add.MulAdd(c, mul)
	}
	return mul.MulAdd(c, add)
}
