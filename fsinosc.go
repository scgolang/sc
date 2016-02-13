package sc

// FSinOsc is a very fast sine wave generator implemented using a
// ringing filter. This generates a much cleaner sine wave than a
// table lookup oscillator and is a lot faster. However, the
// amplitude of the wave will vary with frequency. Generally the
// amplitude will go down as you raise the frequency and go up
// as you lower the frequency.
type FSinOsc struct {
	// Freq is frequency in Hz
	Freq Input
	// Phase is the initial phase offset
	Phase Input
}

func (fso *FSinOsc) defaults() {
	if fso.Freq == nil {
		fso.Freq = C(440)
	}
	if fso.Phase == nil {
		fso.Phase = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (fso FSinOsc) Rate(rate int8) Input {
	CheckRate(rate)
	(&fso).defaults()
	return UgenInput("FSinOsc", rate, 0, 1, fso.Freq, fso.Phase)
}
