package ugens

import . "github.com/briansorahan/sc/types"

// Balance2 equal power panner
type Balance2 struct {
	// L is the left input signal
	L Input
	// R is the right input signal
	R Input
	// Pos stereo position where -1 is hard left and +1 is hard right
	Pos Input
	// Level gain [0, 1]
	Level Input
}

func (self *Balance2) defaults() {
	if self.Pos == nil {
		self.Pos = C(0)
	}
	if self.Level == nil {
		self.Level = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (self Balance2) Rate(rate int8) Input {
	if self.L == nil || self.R == nil {
		panic("Balance2 expects L and R to not be nil")
	}
	checkRate(rate)
	(&self).defaults()
	return UgenInput("Balance2", rate, 0, self.L, self.R, self.Pos, self.Level)
}
