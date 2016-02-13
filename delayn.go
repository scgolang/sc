package sc

// DelayN is a simple delay line with no interpolation.
type DelayN struct {
	// In the input signal
	In Input
	// MaxDelayTime maximum delay time in seconds, which is used
	// to initialize the delay buffer size.
	MaxDelayTime Input
	// DelayTime delay time in seconds
	DelayTime Input
}

func (delayn *DelayN) defaults() {
	if delayn.MaxDelayTime == nil {
		delayn.MaxDelayTime = C(0.2)
	}
	if delayn.DelayTime == nil {
		delayn.DelayTime = C(0.2)
	}
}

// Rate creates a new ugen at a specific rate.
// If an In signal is not provided this method will
// trigger a runtime panic.
func (delayn DelayN) Rate(rate int8) Input {
	CheckRate(rate)
	if delayn.In == nil {
		panic("DelayN expects In to not be nil")
	}
	(&delayn).defaults()
	return UgenInput("DelayN", rate, 0, 1, delayn.In, delayn.MaxDelayTime, delayn.DelayTime)
}
