package sc

// Slope measures the rate of change per second of a signal.
// The formula implemented is:
//     out[i] = (in[i] - in[i-1]) * sampling_rate
type Slope struct {
	// Input signal.
	In Input
}

func (s *Slope) defaults() {
	if s.In == nil {
		s.In = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (s Slope) Rate(rate int8) Input {
	CheckRate(rate)
	(&s).defaults()
	return NewInput("Slope", rate, 0, 1, s.In)
}
