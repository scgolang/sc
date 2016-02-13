package sc

// Decay2 is just like Decay, except it rounds off the attack
// by subtracting one Decay from another. This fixes the sharp
// attacks and clicks that can sometimes occur with Decay.
type Decay2 struct {
	// In is the input signal
	In Input
	// Attack 60dB attack time in seconds
	Attack Input
	// Decay 60dB decay time in seconds
	Decay Input
}

func (decay2 *Decay2) defaults() {
	if decay2.Attack == nil {
		decay2.Attack = C(0.01)
	}
	if decay2.Decay == nil {
		decay2.Decay = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If an In signal is not provided this method will
// trigger a runtime panic.
func (decay2 Decay2) Rate(rate int8) Input {
	CheckRate(rate)
	if decay2.In == nil {
		panic("Decay2 expects In to not be nil")
	}
	(&decay2).defaults()
	return UgenInput("Decay2", rate, 0, 1, decay2.In, decay2.Attack, decay2.Decay)
}
