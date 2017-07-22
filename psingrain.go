package sc

// PSinGrain is a very fast sine grain with a parabolic envelope.
type PSinGrain struct {
	// Frequency in Hz.
	Freq Input

	// Grain duration in seconds.
	Dur Input

	// Grain amplitude.
	Amp Input
}

func (p *PSinGrain) defaults() {
	if p.Freq == nil {
		p.Freq = C(440)
	}
	if p.Dur == nil {
		p.Dur = C(0.2)
	}
	if p.Amp == nil {
		p.Amp = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is a supported value this method will cause a runtime panic.
func (p PSinGrain) Rate(rate int8) Input {
	CheckRate(rate)
	(&p).defaults()
	return NewInput("PSinGrain", rate, 0, 1, p.Freq, p.Dur, p.Amp)
}
