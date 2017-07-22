package sc

// LFClipNoise randomly generates the values -1 or +1 at a rate
// given by the nearest integer division of the sample rate by the freqargument.
// It is probably pretty hard on your speakers!
type LFClipNoise struct {
	Freq Input
}

func (l *LFClipNoise) defaults() {
	if l.Freq == nil {
		l.Freq = C(500)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (l LFClipNoise) Rate(rate int8) Input {
	CheckRate(rate)
	(&l).defaults()
	return NewInput("LFClipNoise", rate, 0, 1, l.Freq)
}
