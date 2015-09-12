package ugens

import . "github.com/scgolang/sc/types"

// Decay2
type Decay2 struct {
	// In is the input signal
	In Input
	// Attack 60dB attack time in seconds
	Attack Input
	// Decay 60dB decay time in seconds
	Decay Input
}

func (self *Decay2) defaults() {
	if self.Attack == nil {
		self.Attack = C(0.01)
	}
	if self.Decay == nil {
		self.Decay = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If an In signal is not provided this method will
// trigger a runtime panic.
func (self Decay2) Rate(rate int8) Input {
	CheckRate(rate)
	if self.In == nil {
		panic("Decay2 expects In to not be nil")
	}
	(&self).defaults()
	return UgenInput("Decay2", rate, 0, 1, self.In, self.Attack, self.Decay)
}
