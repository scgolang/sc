package sc

// Formlet is a resonant filter whose impulse response is like that of a sine wave
// with a Decay2 envelope over it.
// The great advantage to this filter over FOF is that there is no limit
// to the number of overlapping grains since the grain is just the
// impulse response of the filter.
type Formlet struct {
	// In the input signal
	In Input
	// Freq resonant frequency in Hz
	Freq Input
	// AttackTime 60 dB attack time in seconds
	AttackTime Input
	// DecayTime 60 dB decay time in seconds
	DecayTime Input
}

func (formlet *Formlet) defaults() {
	if formlet.Freq == nil {
		formlet.Freq = C(440)
	}
	if formlet.AttackTime == nil {
		formlet.AttackTime = C(1)
	}
	if formlet.DecayTime == nil {
		formlet.DecayTime = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If an In signal is not provided this method will
// trigger a runtime panic.
func (formlet Formlet) Rate(rate int8) Input {
	CheckRate(rate)
	if formlet.In == nil {
		panic("Formlet expects In to not be nil")
	}
	(&formlet).defaults()
	return UgenInput("Formlet", rate, 0, 1, formlet.In, formlet.Freq, formlet.AttackTime, formlet.DecayTime)
}
