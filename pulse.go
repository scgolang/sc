package sc

// Pulse band-limited pulse wave generator with pulse width modulation.
type Pulse struct {
	// Freq in Hz
	Freq Input
	// Width pulse width duty cycle from 0 to 1
	Width Input
}

func (pulse *Pulse) defaults() {
	if pulse.Freq == nil {
		pulse.Freq = C(440)
	}
	if pulse.Width == nil {
		pulse.Width = C(0.5)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (pulse Pulse) Rate(rate int8) Input {
	CheckRate(rate)
	(&pulse).defaults()
	return UgenInput("Pulse", rate, 0, 1, pulse.Freq, pulse.Width)
}
