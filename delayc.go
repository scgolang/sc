package sc

// DelayC
// Simple delay line with cubic interpolation.
// DelayC is more computationally intensive than DelayL, but more accurate.
// Note that DelayC needs at least 4 samples of delay buffer.
type DelayC struct {
	// In the input signal
	In Input
	// MaxDelayTime maximum delay time in seconds, which is used
	// to initialize the delay buffer size.
	MaxDelayTime Input
	// DelayTime delay time in seconds
	DelayTime Input
}

func (delayc *DelayC) defaults() {
	if delayc.MaxDelayTime == nil {
		delayc.MaxDelayTime = C(0.2)
	}
	if delayc.DelayTime == nil {
		delayc.DelayTime = C(0.2)
	}
}

// Rate creates a new ugen at a specific rate.
// If an In signal is not provided this method will
// trigger a runtime panic.
func (delayc DelayC) Rate(rate int8) Input {
	CheckRate(rate)
	if delayc.In == nil {
		panic("DelayC expects In to not be nil")
	}
	(&delayc).defaults()
	return UgenInput("DelayC", rate, 0, 1, delayc.In, delayc.MaxDelayTime, delayc.DelayTime)
}
