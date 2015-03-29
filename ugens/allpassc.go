package ugens

import . "github.com/briansorahan/sc/types"

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

func (self *AllpassC) defaults() {
	if self.MaxDelay == nil {
		self.MaxDelay = C(0.2)
	}
	if self.Delay == nil {
		self.Delay = C(0.2)
	}
	if self.Decay == nil {
		self.Decay = C(1)
	}
}

func (self AllpassC) Rate(rate int8) Input {
	if self.In == nil {
		panic("AllpassC expects In to not be nil")
	}
	checkRate(rate)
	(&self).defaults()
	return UgenInput("AllpassC", rate, 0, self.In, self.MaxDelay, self.Delay, self.Decay)
}
