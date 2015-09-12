package ugens

import . "github.com/scgolang/sc/types"

// Gate
type Gate struct {
	// In is the input signal.
	In Input
	// Trig is the rigger signal. The output is held fixed when this is non-positive.
	Trig Input
}

func (self *Gate) defaults() {
	if self.Trig == nil {
		self.Trig = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If an In signal is not provided this method will
// trigger a runtime panic.
func (self Gate) Rate(rate int8) Input {
	CheckRate(rate)
	if self.In == nil {
		panic("Gate expects In to not be nil")
	}
	(&self).defaults()
	return UgenInput("Gate", rate, 0, 1, self.In, self.Trig)
}
