package ugens

import . "github.com/scgolang/sc/types"

// LFTri is a non-band-limited triangle oscillator.
// Output is [-1, 1].
type LFTri struct {
	// Freq is the approximate rate at which to
	// generate random values
	Freq Input
	// Iphase initial phase offset in the range [0, 4]
	Iphase Input
}

func (self *LFTri) defaults() {
	if self.Freq == nil {
		self.Freq = C(500)
	}
	if self.Iphase == nil {
		self.Iphase = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (self LFTri) Rate(rate int8) Input {
	checkRate(rate)
	(&self).defaults()
	return UgenInput("LFTri", rate, 0, self.Freq, self.Iphase)
}
