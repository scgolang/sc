package ugens

import . "github.com/briansorahan/sc/types"

// Saw
type Saw struct {
	Freq Input
}

func (self *Saw) defaults() {
	if self.Freq == nil {
		self.Freq = C(440)
	}
}

func (self Saw) Rate(rate int8) Input {
	checkRate(rate)
	(&self).defaults()
	return UgenInput("Saw", rate, 0, self.Freq)
}
