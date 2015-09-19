package sc


// Decay
type Decay struct {
	// In is the input signal
	In Input
	// Decay 60dB decay time in seconds
	Decay Input
}

func (self *Decay) defaults() {
	if self.Decay == nil {
		self.Decay = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If an In signal is not provided this method will
// trigger a runtime panic.
func (self Decay) Rate(rate int8) Input {
	CheckRate(rate)
	if self.In == nil {
		panic("Decay expects In to not be nil")
	}
	(&self).defaults()
	return UgenInput("Decay", rate, 0, 1, self.In, self.Decay)
}
