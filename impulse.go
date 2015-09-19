package sc


// Impulse non-band-limited single-sample impulses
type Impulse struct {
	// Freq frequency in Hz
	Freq Input
	// Phase offset in cycles [0, 1]
	Phase Input
}

func (self *Impulse) defaults() {
	if self.Freq == nil {
		self.Freq = C(440)
	}
	if self.Phase == nil {
		self.Phase = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (self Impulse) Rate(rate int8) Input {
	CheckRate(rate)
	(&self).defaults()
	return UgenInput("Impulse", rate, 0, 1, self.Freq, self.Phase)
}
