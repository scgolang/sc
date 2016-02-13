package sc

// Decay is an exponential decay.
type Decay struct {
	// In is the input signal
	In Input
	// Decay 60dB decay time in seconds
	Decay Input
}

func (decay *Decay) defaults() {
	if decay.Decay == nil {
		decay.Decay = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If an In signal is not provided this method will
// trigger a runtime panic.
func (decay Decay) Rate(rate int8) Input {
	CheckRate(rate)
	if decay.In == nil {
		panic("Decay expects In to not be nil")
	}
	(&decay).defaults()
	return UgenInput("Decay", rate, 0, 1, decay.In, decay.Decay)
}
