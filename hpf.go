package sc

// HPF is a 2nd order Butterworth highpass filter.
type HPF struct {
	// In is the input signal
	In Input
	// Freq cutoff in Hz
	Freq Input
}

func (hpf *HPF) defaults() {
	if hpf.Freq == nil {
		hpf.Freq = C(440)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (hpf HPF) Rate(rate int8) Input {
	if hpf.In == nil {
		panic("HPF expects In to not be nil")
	}
	CheckRate(rate)
	(&hpf).defaults()
	return NewInput("HPF", rate, 0, 1, hpf.In, hpf.Freq)
}
