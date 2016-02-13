package sc

// CombL is a delay line with linear interpolation.
type CombL struct {
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

func (combl *CombL) defaults() {
	if combl.MaxDelayTime == nil {
		combl.MaxDelayTime = C(0.2)
	}
	if combl.DelayTime == nil {
		combl.DelayTime = C(0.2)
	}
	if combl.DecayTime == nil {
		combl.DecayTime = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If an In signal is not provided this method will
// trigger a runtime panic.
func (combl CombL) Rate(rate int8) Input {
	CheckRate(rate)
	if combl.In == nil {
		panic("CombL expects In to not be nil")
	}
	(&combl).defaults()
	return UgenInput("CombL", rate, 0, 1, combl.In, combl.MaxDelayTime, combl.DelayTime, combl.DecayTime)
}
