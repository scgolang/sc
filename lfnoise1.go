package sc

// LFNoise1 generates linearly interpolated random values at a
// rate which is the nearest integer division of the sample
// rate by the Freq parameter.
type LFNoise1 struct {
	// Freq is the approximate rate at which to
	// generate random values
	Freq Input
}

func (lfn *LFNoise1) defaults() {
	if lfn.Freq == nil {
		lfn.Freq = C(500)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (lfn LFNoise1) Rate(rate int8) Input {
	CheckRate(rate)
	(&lfn).defaults()
	return UgenInput("LFNoise1", rate, 0, 1, lfn.Freq)
}
