package sc


// Saw
type Saw struct {
	Freq Input
}

func (self *Saw) defaults() {
	if self.Freq == nil {
		self.Freq = C(440)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (self Saw) Rate(rate int8) Input {
	CheckRate(rate)
	(&self).defaults()
	return UgenInput("Saw", rate, 0, 1, self.Freq)
}
