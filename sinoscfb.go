package sc

// SinOscFB is a sine oscillator that has phase modulation feedback: its output
// plugs back into the phase input.
// This allows a modulation between a sine wave and a sawtooth like wave.
// Overmodulation causes chaotic oscillation.
// It may be useful if you want to simulate feedback FM synths.
type SinOscFB struct {
	Freq     Input
	Feedback Input
}

func (s *SinOscFB) defaults() {
	if s.Freq == nil {
		s.Freq = C(440)
	}
	if s.Feedback == nil {
		s.Feedback = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (s SinOscFB) Rate(rate int8) Input {
	CheckRate(rate)
	(&s).defaults()
	return NewInput("SinOscFB", rate, 0, 1, s.Freq, s.Feedback)
}
