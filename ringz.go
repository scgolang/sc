package sc

// This is the same as Ringz , except that it has a constant gain at 0 dB
// instead of being constant skirt.
// It is a two pole resonant filter with zeroes at z = +-1.
type Ringz struct {
	// In is the input signal.
	In Input

	// Freq is the resonant frequency in Hertz.
	// WARNING: due to the nature of its implementation frequency
	// values close to 0 may cause glitches and/or extremely loud audio artifacts!
	Freq Input

	// DecayTime is the 60 dB decay time of the filter.
	DecayTime Input
}

func (r *Ringz) defaults() {
	if r.In == nil {
		r.In = C(0)
	}
	if r.Freq == nil {
		r.Freq = C(440)
	}
	if r.DecayTime == nil {
		r.DecayTime = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (r Ringz) Rate(rate int8) Input {
	CheckRate(rate)
	(&r).defaults()
	return NewInput("Ringz", rate, 0, 1, r.In, r.Freq, r.DecayTime)
}
