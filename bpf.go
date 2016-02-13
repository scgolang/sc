package sc

// BPF a resonant low pass filter
type BPF struct {
	// In is the input signal
	In Input
	// Freq cutoff in Hz
	Freq Input
	// RQ Reciprocal of Q
	RQ Input
}

func (bpf *BPF) defaults() {
	if bpf.Freq == nil {
		bpf.Freq = C(440)
	}
	if bpf.RQ == nil {
		bpf.RQ = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (bpf BPF) Rate(rate int8) Input {
	if bpf.In == nil {
		panic("BPF expects In to not be nil")
	}
	CheckRate(rate)
	(&bpf).defaults()
	return UgenInput("BPF", rate, 0, 1, bpf.In, bpf.Freq, bpf.RQ)
}
