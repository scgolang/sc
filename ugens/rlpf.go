package ugens

import . "github.com/briansorahan/sc/types"

// RLPF a resonant low pass filter
type RLPF struct {
	// In is the input signal
	In Input
	// Freq cutoff in Hz
	Freq Input
	// RQ Reciprocal of Q
	RQ Input
}

func (self *RLPF) defaults() {
	if self.Freq == nil {
		self.Freq = C(440)
	}
	if self.RQ == nil {
		self.RQ = C(1)
	}
}

func (self RLPF) Rate(rate int8) *Node {
	if self.In == nil {
		panic("RLPF expects In to not be nil")
	}
	checkRate(rate)
	(&self).defaults()
	n := NewNode("RLPF", rate, 0, self.In, self.Freq, self.RQ)
	return n
}
