package sc

// RunningSum is a running sum over a user specified number of samples,
// useful for running RMS power windowing.
type RunningSum struct {
	In      Input
	NumSamp Input
}

func (r *RunningSum) defaults() {
	if r.NumSamp == nil {
		r.NumSamp = C(40)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
// If the input signal is nil this method will panic.
func (r RunningSum) Rate(rate int8) Input {
	CheckRate(rate)
	if r.In == nil {
		panic("RunningSum requires an input signal")
	}
	(&r).defaults()
	return NewInput("RunningSum", rate, 0, 1, r.In, r.NumSamp)
}
