package sc


// CombL
type CombL struct {
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

func (self *CombL) defaults() {
	if self.MaxDelayTime == nil {
		self.MaxDelayTime = C(0.2)
	}
	if self.DelayTime == nil {
		self.DelayTime = C(0.2)
	}
	if self.DecayTime == nil {
		self.DecayTime = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If an In signal is not provided this method will
// trigger a runtime panic.
func (self CombL) Rate(rate int8) Input {
	CheckRate(rate)
	if self.In == nil {
		panic("CombL expects In to not be nil")
	}
	(&self).defaults()
	return UgenInput("CombL", rate, 0, 1, self.In, self.MaxDelayTime, self.DelayTime, self.DecayTime)
}
