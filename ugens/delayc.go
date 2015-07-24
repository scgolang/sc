package ugens

import . "github.com/scgolang/sc/types"

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

func (self *DelayC) defaults() {
	if self.MaxDelayTime == nil {
		self.MaxDelayTime = C(0.2)
	}
	if self.DelayTime == nil {
		self.DelayTime = C(0.2)
	}
}

// Rate creates a new ugen at a specific rate.
// If an In signal is not provided this method will
// trigger a runtime panic.
func (self DelayC) Rate(rate int8) Input {
	checkRate(rate)
	if self.In == nil {
		panic("DelayC expects In to not be nil")
	}
	(&self).defaults()
	return UgenInput("DelayC", rate, 0, 1, self.In, self.MaxDelayTime, self.DelayTime)
}
