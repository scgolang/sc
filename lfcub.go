package sc

// LFCub is a sine-like shape made of two cubic pieces.
// It is smoother than LFPar.
type LFCub struct {
	// Freq frequency in Hz
	Freq Input
	// Iphase initial phase offset
	Iphase Input
}

func (lfcub *LFCub) defaults() {
	if lfcub.Freq == nil {
		lfcub.Freq = C(440)
	}
	if lfcub.Iphase == nil {
		lfcub.Iphase = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (lfcub LFCub) Rate(rate int8) Input {
	CheckRate(rate)
	(&lfcub).defaults()
	return UgenInput("LFCub", rate, 0, 1, lfcub.Freq, lfcub.Iphase)
}
