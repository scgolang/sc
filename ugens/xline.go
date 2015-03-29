package ugens

import . "github.com/briansorahan/sc/types"

type XLine struct {
	Start Input
	End   Input
	Dur   Input
	Done  int
}

func (self *XLine) defaults() {
	if self.Start == nil {
		self.Start = C(1)
	}
	if self.End == nil {
		self.End = C(2)
	}
	if self.Dur == nil {
		self.Dur = C(1)
	}
}

// Rate ugen implementation
func (self XLine) Rate(rate int8) Input {
	checkRate(rate)
	(&self).defaults()
	return UgenInput("XLine", rate, 0, self.Start, self.End, self.Dur, C(float32(self.Done)))
}
