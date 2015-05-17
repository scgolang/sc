package ugens

import . "github.com/scgolang/sc/types"

// BPF a resonant low pass filter
type BPF struct {
	// In is the input signal
	In Input
	// Freq cutoff in Hz
	Freq Input
	// RQ Reciprocal of Q
	RQ Input
}

func (self *BPF) defaults() {
	if self.Freq == nil {
		self.Freq = C(440)
	}
	if self.RQ == nil {
		self.RQ = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (self BPF) Rate(rate int8) Input {
	if self.In == nil {
		panic("BPF expects In to not be nil")
	}
	checkRate(rate)
	(&self).defaults()
	return UgenInput("BPF", rate, 0, 1, self.In, self.Freq, self.RQ)
}
