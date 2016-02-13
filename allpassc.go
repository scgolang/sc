package sc

// AllpassC allpass delay with cubic interpolation
type AllpassC struct {
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

func (apc *AllpassC) defaults() {
	if apc.MaxDelay == nil {
		apc.MaxDelay = C(0.2)
	}
	if apc.Delay == nil {
		apc.Delay = C(0.2)
	}
	if apc.Decay == nil {
		apc.Decay = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (apc AllpassC) Rate(rate int8) Input {
	if apc.In == nil {
		panic("AllpassC expects In to not be nil")
	}
	CheckRate(rate)
	(&apc).defaults()
	return UgenInput("AllpassC", rate, 0, 1, apc.In, apc.MaxDelay, apc.Delay, apc.Decay)
}
