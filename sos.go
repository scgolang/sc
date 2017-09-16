package sc

// SOS is a standard second order filter section.
// Filter coefficients are given directly rather than calculated for you. Formula is equivalent to:
//     out(i) = (a0 * in(i)) + (a1 * in(i-1)) + (a2 * in(i-2)) + (b1 * out(i-1)) + (b2 * out(i-2))
type SOS struct {
	// In is the input signal.
	In Input

	// Coefficients.
	A0, A1, A2 Input
	B1, B2     Input
}

func (s *SOS) defaults() {
	if s.In == nil {
		panic("SOS needs an input")
	}
	if s.A0 == nil {
		s.A0 = C(0)
	}
	if s.A1 == nil {
		s.A1 = C(0)
	}
	if s.A2 == nil {
		s.A2 = C(0)
	}
	if s.B1 == nil {
		s.B1 = C(0)
	}
	if s.B2 == nil {
		s.B2 = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (s SOS) Rate(rate int8) Input {
	CheckRate(rate)
	(&s).defaults()
	return NewUgenInput("SOS", rate, 0, 1, s.In, s.A0, s.A1, s.A2, s.B1, s.B2)
}
