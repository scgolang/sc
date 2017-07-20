package sc

// Shaper performs waveshaping on the input signal by indexing into the table.
type Shaper struct {
	// BufNum is the number of a buffer filled in wavetable format
	// containing the transfer function.
	BufNum Input

	// In is the input signal.
	In Input
}

func (s *Shaper) defaults() {
	if s.In == nil {
		s.In = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
// If BufNum is nil this method will panic.
func (s Shaper) Rate(rate int8) Input {
	CheckRate(rate)
	if s.BufNum == nil {
		panic("Shaper requires a buffer number")
	}
	(&s).defaults()
	return NewInput("Shaper", rate, 0, 1, s.BufNum, s.In)
}
