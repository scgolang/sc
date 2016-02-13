package sc

// GrayNoise generates noise which results from flipping
// random bits in a word. This type of noise has a high
// RMS level relative to its peak to peak level.
// The spectrum is emphasized towards lower frequencies.
type GrayNoise struct {
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (gn GrayNoise) Rate(rate int8) Input {
	CheckRate(rate)
	return UgenInput("GrayNoise", rate, 0, 1)
}
