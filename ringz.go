package sc

// Ringz is the same as Resonz, except that it is a constant skirt gain filter,
// meaning that the peak gain depends on the value of Q.
// Also, instead of the resonance parameter in Resonz,
// the bandwidth is specified in a 60dB ring decay time.
// One Ringz is equivalent to one component of the Klank UGen.
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
