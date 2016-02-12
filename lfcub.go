package sc

// LFCub is a sine-like shape made of two cubic pieces.
// It is smoother than LFPar.
type LFCub struct {
	// Freq frequency in Hz
	Freq Input
	// Iphase initial phase offset
	Iphase Input
}

func (self *LFCub) defaults() {
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
func (self LFCub) Rate(rate int8) Input {
	CheckRate(rate)
	(&self).defaults()
	return UgenInput("LFCub", rate, 0, 1, self.Freq, self.Iphase)
}
