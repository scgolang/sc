package ugens

import . "github.com/scgolang/sc/types"

// LFSaw a non-band-limited sawtooth oscillator
// output ranges from -1 to +1
type LFSaw struct {
	// Freq frequency in Hz
	Freq Input
	// Iphase initial phase offset in cycles:
	// for efficiency this is in the rage [0, 2]
	Iphase Input
}

func (self *LFSaw) defaults() {
	if self.Freq == nil {
		self.Freq = C(440)
	}
	if self.Iphase == nil {
		self.Iphase = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (self LFSaw) Rate(rate int8) Input {
	CheckRate(rate)
	(&self).defaults()
	return UgenInput("LFSaw", rate, 0, 1, self.Freq, self.Iphase)
}
