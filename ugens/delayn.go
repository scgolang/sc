package ugens

import . "github.com/scgolang/sc/types"

// DelayN
// Simple delay line with no interpolation.
type DelayN struct {
	// In the input signal
	In Input
	// MaxDelayTime maximum delay time in seconds, which is used
	// to initialize the delay buffer size.
	MaxDelayTime Input
	// DelayTime delay time in seconds
	DelayTime Input
}

func (self *DelayN) defaults() {
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
func (self DelayN) Rate(rate int8) Input {
	CheckRate(rate)
	if self.In == nil {
		panic("DelayN expects In to not be nil")
	}
	(&self).defaults()
	return UgenInput("DelayN", rate, 0, 1, self.In, self.MaxDelayTime, self.DelayTime)
}
