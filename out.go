package sc

// Out writes a signal to a bus.
type Out struct {
	Bus      Input
	Channels Input
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (out Out) Rate(rate int8) Ugen {
	CheckRate(rate)
	u := NewUgen("Out", rate, 0, 1, out.Bus, out.Channels)
	return *u
}
