package sc

// moddif returns the smaller of the great circle distances between the two points.
// This func panics if x is nil or if rate is not one of the supported rates (IR, KR, AR).
func moddif(rate int8, x, y, mod Input, numOutputs int) Input {
	CheckRate(rate)

	if x == nil {
		panic("Moddif expects the first argument to not be nil")
	}
	if y == nil {
		y = C(0)
	}
	if mod == nil {
		mod = C(1)
	}
	return NewInput("ModDif", rate, 0, numOutputs, x, y, mod)
}
