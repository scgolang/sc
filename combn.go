package sc

// CombN is a delay line with no interpolation.
type CombN struct {
	// In the input signal
	In Input
	// MaxDelayTime maximum delay time in seconds, which is used
	// to initialize the delay buffer size.
	MaxDelayTime Input
	// DelayTime delay time in seconds
	DelayTime Input
	// DecayTime time for the echoes to decay by 60dB.
	// If this time is negative then the feedback coefficient will be
	// negative, thus emphasizing only odd harmonics an octave lower.
	DecayTime Input
}

func (combn *CombN) defaults() {
	if combn.MaxDelayTime == nil {
		combn.MaxDelayTime = C(0.2)
	}
	if combn.DelayTime == nil {
		combn.DelayTime = C(0.2)
	}
	if combn.DecayTime == nil {
		combn.DecayTime = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If an In signal is not provided this method will
// trigger a runtime panic.
func (combn CombN) Rate(rate int8) Input {
	CheckRate(rate)
	if combn.In == nil {
		panic("CombN expects In to not be nil")
	}
	(&combn).defaults()
	return UgenInput("CombN", rate, 0, 1, combn.In, combn.MaxDelayTime, combn.DelayTime, combn.DecayTime)
}
