package ugens

import . "github.com/scgolang/sc/types"

// Dust generates random impulses from -1 to +1
type Dust2 struct {
	// Density is the average number of impulses per second
	Density Input
}

func (self *Dust2) defaults() {
	if self.Density == nil {
		self.Density = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (self Dust2) Rate(rate int8) Input {
	CheckRate(rate)
	(&self).defaults()
	return UgenInput("Dust2", rate, 0, 1, self.Density)
}
