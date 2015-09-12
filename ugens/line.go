package ugens

import . "github.com/scgolang/sc/types"

// Line generates a line from the start value to the end value
type Line struct {
	Start Input
	End   Input
	Dur   Input
	Done  int
}

func (self *Line) defaults() {
	if self.Start == nil {
		self.Start = C(0)
	}
	if self.End == nil {
		self.End = C(1)
	}
	if self.Dur == nil {
		self.Dur = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (self Line) Rate(rate int8) Input {
	CheckRate(rate)
	(&self).defaults()
	return UgenInput("Line", rate, 0, 1, self.Start, self.End, self.Dur, C(float32(self.Done)))
}
