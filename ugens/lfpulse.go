package ugens

import . "github.com/briansorahan/sc/types"

// LFPulse a resonant low pass filter
type LFPulse struct {
	// Freq in Hz
	Freq Input
	// Iphase initial phase offset in cycles (0..1)
	Iphase Input
	// Width pulse width duty cycle from 0 to 1
	Width Input
}

func (self *LFPulse) defaults() {
	if self.Freq == nil {
		self.Freq = C(440)
	}
	if self.Iphase == nil {
		self.Iphase = C(0)
	}
	if self.Width == nil {
		self.Width = C(0.5)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (self LFPulse) Rate(rate int8) Input {
	checkRate(rate)
	(&self).defaults()
	return UgenInput("LFPulse", rate, 0, self.Freq, self.Iphase, self.Width)
}
