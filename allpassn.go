package sc

// AllpassN allpass delay with no interpolation
type AllpassN struct {
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

func (apn *AllpassN) defaults() {
	if apn.MaxDelay == nil {
		apn.MaxDelay = C(0.2)
	}
	if apn.Delay == nil {
		apn.Delay = C(0.2)
	}
	if apn.Decay == nil {
		apn.Decay = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (apn AllpassN) Rate(rate int8) Input {
	if apn.In == nil {
		panic("AllpassN expects In to not be nil")
	}
	CheckRate(rate)
	(&apn).defaults()
	return UgenInput("AllpassN", rate, 0, 1, apn.In, apn.MaxDelay, apn.Delay, apn.Decay)
}
