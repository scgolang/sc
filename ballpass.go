package sc

// BAllPass is an allpass filter based on the Second Order Section (SOS) biquad UGen.
// NOTE: Biquad coefficient calculations imply certain amount of CPU overhead.
// These plugin UGens contain optimizations such that the coefficients get updated
// only when there has been a change to one of the filter's parameters.
// This can cause spikes in CPU performance and should be considered when using
// several of these units.
type BAllPass struct {
	// In is the input signal.
	In Input

	// Freq is the center frequency.
	// WARNING: due to the nature of its implementation frequency values close to 0
	// may cause glitches and/or extremely loud audio artifacts!
	Freq Input

	// RQ is the reciprocal of Q. bandwidth / cutoffFreq.
	RQ Input
}

func (b *BAllPass) defaults() {
	if b.Freq == nil {
		b.Freq = C(1200)
	}
	if b.RQ == nil {
		b.RQ = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
// This method will also panic if In is nil.
func (b BAllPass) Rate(rate int8) Input {
	CheckRate(rate)

	if b.In == nil {
		panic("BAllpass expects In to not be nil")
	}
	(&b).defaults()

	return NewInput("BAllPass", rate, 0, 1, b.In, b.Freq, b.RQ)

}
