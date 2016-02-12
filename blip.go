package sc

// Blip band-limited impulse generator
type Blip struct {
	// Freq frequency in Hz
	Freq Input
	// Harm the number of harmonics
	Harm Input
}

func (self *Blip) defaults() {
	if self.Freq == nil {
		self.Freq = C(440)
	}
	if self.Harm == nil {
		self.Harm = C(200)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (self Blip) Rate(rate int8) Input {
	CheckRate(rate)
	(&self).defaults()
	return UgenInput("Blip", rate, 0, 1, self.Freq, self.Harm)
}
