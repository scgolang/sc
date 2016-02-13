package sc

// RLPF a resonant low pass filter
type RLPF struct {
	// In is the input signal
	In Input
	// Freq cutoff in Hz
	Freq Input
	// RQ Reciprocal of Q
	RQ Input
}

func (rlpf *RLPF) defaults() {
	if rlpf.Freq == nil {
		rlpf.Freq = C(440)
	}
	if rlpf.RQ == nil {
		rlpf.RQ = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (rlpf RLPF) Rate(rate int8) Input {
	if rlpf.In == nil {
		panic("RLPF expects In to not be nil")
	}
	CheckRate(rate)
	(&rlpf).defaults()
	return UgenInput("RLPF", rate, 0, 1, rlpf.In, rlpf.Freq, rlpf.RQ)
}
