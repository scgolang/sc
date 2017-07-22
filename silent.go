package sc

// Silent outputs silence.
type Silent struct {
	// Number of channels to output. Defaults to 1 if not set.
	NumChannels int
}

func (s *Silent) defaults() {
	if s.NumChannels == 0 {
		s.NumChannels = 1
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (s Silent) Rate(rate int8) Input {
	CheckRate(rate)
	(&s).defaults()
	return NewInput("DC", rate, 0, s.NumChannels, C(0))
}
