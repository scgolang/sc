package ugens

import . "github.com/scgolang/sc/types"

// SinOsc is a table-lookup sinewave oscillator
type SinOsc struct {
	// Freq is frequency in Hz
	Freq Input
	// Phase is the initial phase offset
	Phase Input
}

func (self *SinOsc) defaults() {
	if self.Freq == nil {
		self.Freq = C(440)
	}
	if self.Phase == nil {
		self.Phase = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (self SinOsc) Rate(rate int8) Input {
	checkRate(rate)
	(&self).defaults()
	return UgenInput("SinOsc", rate, 0, 1, self.Freq, self.Phase)
}
