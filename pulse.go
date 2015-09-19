package sc


// Pulse band-limited pulse wave generator with pulse width modulation.
type Pulse struct {
	// Freq in Hz
	Freq Input
	// Width pulse width duty cycle from 0 to 1
	Width Input
}

func (self *Pulse) defaults() {
	if self.Freq == nil {
		self.Freq = C(440)
	}
	if self.Width == nil {
		self.Width = C(0.5)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (self Pulse) Rate(rate int8) Input {
	CheckRate(rate)
	(&self).defaults()
	return UgenInput("Pulse", rate, 0, 1, self.Freq, self.Width)
}
