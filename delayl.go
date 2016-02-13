package sc

// DelayL
// Simple delay line with linear interpolation.
// DelayL is more computationally intensive than DelayN, but more accurate.
type DelayL struct {
	// In the input signal
	In Input
	// MaxDelayTime maximum delay time in seconds, which is used
	// to initialize the delay buffer size.
	MaxDelayTime Input
	// DelayTime delay time in seconds
	DelayTime Input
}

func (delayl *DelayL) defaults() {
	if delayl.MaxDelayTime == nil {
		delayl.MaxDelayTime = C(0.2)
	}
	if delayl.DelayTime == nil {
		delayl.DelayTime = C(0.2)
	}
}

// Rate creates a new ugen at a specific rate.
// If an In signal is not provided this method will
// trigger a runtime panic.
func (delayl DelayL) Rate(rate int8) Input {
	CheckRate(rate)
	if delayl.In == nil {
		panic("DelayL expects In to not be nil")
	}
	(&delayl).defaults()
	return UgenInput("DelayL", rate, 0, 1, delayl.In, delayl.MaxDelayTime, delayl.DelayTime)
}
