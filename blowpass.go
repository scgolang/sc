package sc

// BLowPass is a lowpass filter based on the Second Order Section
// biquad UGen.
type BLowPass struct {
	// In is the input signal.
	In Input
	// Freq is frequency in Hz.
	Freq Input
	// RQ is the reciprocal of Q, bandwidth / cutoff.
	RQ Input
}

func (blp *BLowPass) defaults() {
	if blp.In == nil {
		panic("BLowPass needs an input")
	}
	if blp.Freq == nil {
		blp.Freq = C(1200)
	}
	if blp.RQ == nil {
		blp.RQ = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (blp BLowPass) Rate(rate int8) Input {
	CheckRate(rate)
	(&blp).defaults()
	return UgenInput("BLowPass", rate, 0, 1, blp.In, blp.Freq, blp.RQ)
}
