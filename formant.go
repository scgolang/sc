package sc

// Formant generates a set of harmonics around a formant frequency
// at a given fundamental frequency.
// The frequency inputs are read at control rate only,
// so if you use an audio rate UGen as an input,
// it will only be sampled at the start of each audio synthesis block.
type Formant struct {
	// FundFreq is the fundamental frequency in Hertz. (control rate)
	FundFreq Input

	// FormantFreq is the formant frequency in Hertz. (control rate)
	FormantFreq Input

	// BWFreq is the pulse width frequency in Hertz.
	// Controls the bandwidth of the formant. (control rate)
	// Must be greater than or equal to fundfreq.
	BWFreq Input
}

func (f *Formant) defaults() {
	if f.FundFreq == nil {
		f.FundFreq = C(440)
	}
	if f.FormantFreq == nil {
		f.FormantFreq = C(1760)
	}
	if f.BWFreq == nil {
		f.BWFreq = C(880)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (f Formant) Rate(rate int8) Input {
	CheckRate(rate)
	(&f).defaults()
	return NewInput("Formant", rate, 0, 1, f.FundFreq, f.FormantFreq, f.BWFreq)
}
