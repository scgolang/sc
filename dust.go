package sc

// Dust generates random impulses from 0 to +1
type Dust struct {
	// Density is the average number of impulses per second
	Density Input
}

func (self *Dust) defaults() {
	if self.Density == nil {
		self.Density = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (self Dust) Rate(rate int8) Input {
	CheckRate(rate)
	(&self).defaults()
	return UgenInput("Dust", rate, 0, 1, self.Density)
}
