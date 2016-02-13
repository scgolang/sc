package sc

// LFTri is a non-band-limited triangle oscillator.
// Output is [-1, 1].
type LFTri struct {
	// Freq is the approximate rate at which to
	// generate random values
	Freq Input
	// Iphase initial phase offset in the range [0, 4]
	Iphase Input
}

func (lftri *LFTri) defaults() {
	if lftri.Freq == nil {
		lftri.Freq = C(500)
	}
	if lftri.Iphase == nil {
		lftri.Iphase = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (lftri LFTri) Rate(rate int8) Input {
	CheckRate(rate)
	(&lftri).defaults()
	return UgenInput("LFTri", rate, 0, 1, lftri.Freq, lftri.Iphase)
}
