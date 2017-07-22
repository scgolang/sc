package sc

// LFDClipNoise is like LFClipNoise, except it generates the values -1 or +1
// at a rate given by the freq argument, with two differences:
//     no time quantization
//     fast recovery from low freq values1
// If you don't need very high or very low freqs, or use fixed freqs,
// LFDClipNoise is more efficient.
type LFDClipNoise struct {
	Freq Input
}

func (l *LFDClipNoise) defaults() {
	if l.Freq == nil {
		l.Freq = C(500)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (l LFDClipNoise) Rate(rate int8) Input {
	CheckRate(rate)
	(&l).defaults()
	return NewInput("LFDClipNoise", rate, 0, 1, l.Freq)
}
