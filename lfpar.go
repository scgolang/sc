package sc

// LFPar is a sine-like shape made of two parabolas and
// the integral of a triangular wave.
// It has audible odd harmonics and is non-band-limited.
// Output ranges from -1 to +1.
type LFPar struct {
	// Freq is frequency in Hz.
	Freq Input

	// IPhase is the initial phase (0..1).
	IPhase Input
}

func (l *LFPar) defaults() {
	if l.Freq == nil {
		l.Freq = C(440)
	}
	if l.IPhase == nil {
		l.IPhase = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (l LFPar) Rate(rate int8) Input {
	CheckRate(rate)
	(&l).defaults()
	return NewInput("LFPar", rate, 0, 1, l.Freq, l.IPhase)
}
