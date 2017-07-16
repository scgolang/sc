package sc

// LPF is a second order low pass filter.
type LPF struct {
	// In is the input signal
	In Input
	// Freq cutoff in Hz
	Freq Input
}

func (lpf *LPF) defaults() {
	if lpf.Freq == nil {
		lpf.Freq = C(440)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (lpf LPF) Rate(rate int8) Input {
	if lpf.In == nil {
		panic("LPF expects In to not be nil")
	}
	CheckRate(rate)
	(&lpf).defaults()
	return NewInput("LPF", rate, 0, 1, lpf.In, lpf.Freq)
}
