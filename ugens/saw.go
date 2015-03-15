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

func (self Saw) Rate(rate int8) *Node {
	checkRate(rate)
	(&self).defaults()
	return NewNode("Saw", rate, 0, self.Freq)
}
