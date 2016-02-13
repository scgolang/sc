package sc

// BRF a second order band reject filter
type BRF struct {
	// In is the input signal
	In Input
	// Freq cutoff in Hz
	Freq Input
	// RQ Reciprocal of Q
	RQ Input
}

func (brf *BRF) defaults() {
	if brf.Freq == nil {
		brf.Freq = C(440)
	}
	if brf.RQ == nil {
		brf.RQ = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (brf BRF) Rate(rate int8) Input {
	if brf.In == nil {
		panic("BRF expects In to not be nil")
	}
	CheckRate(rate)
	(&brf).defaults()
	return UgenInput("BRF", rate, 0, 1, brf.In, brf.Freq, brf.RQ)
}
