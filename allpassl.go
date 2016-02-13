package sc

// AllpassL allpass delay with linear interpolation
type AllpassL struct {
	// In is the input signal
	In Input
	// MaxDelay is maximum delay time in seconds.
	// This is used to initialize the delay buffer.
	MaxDelay Input
	// Delay time in seconds
	Delay Input
	// Decay time for the echoes to decay by 60dB.
	// If this is negative then the feedback coefficient will
	// be negative, thus emphasizing only odd harmonics
	// at a lower octave.
	Decay Input
}

func (apl *AllpassL) defaults() {
	if apl.MaxDelay == nil {
		apl.MaxDelay = C(0.2)
	}
	if apl.Delay == nil {
		apl.Delay = C(0.2)
	}
	if apl.Decay == nil {
		apl.Decay = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (apl AllpassL) Rate(rate int8) Input {
	if apl.In == nil {
		panic("AllpassL expects In to not be nil")
	}
	CheckRate(rate)
	(&apl).defaults()
	return UgenInput("AllpassL", rate, 0, 1, apl.In, apl.MaxDelay, apl.Delay, apl.Decay)
}
