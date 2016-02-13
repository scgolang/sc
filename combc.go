package sc

// CombC is a delay line with cubic interpolation.
type CombC struct {
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

func (combc *CombC) defaults() {
	if combc.MaxDelayTime == nil {
		combc.MaxDelayTime = C(0.2)
	}
	if combc.DelayTime == nil {
		combc.DelayTime = C(0.2)
	}
	if combc.DecayTime == nil {
		combc.DecayTime = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If an In signal is not provided this method will
// trigger a runtime panic.
func (combc CombC) Rate(rate int8) Input {
	CheckRate(rate)
	if combc.In == nil {
		panic("CombC expects In to not be nil")
	}
	(&combc).defaults()
	return UgenInput("CombC", rate, 0, 1, combc.In, combc.MaxDelayTime, combc.DelayTime, combc.DecayTime)
}
