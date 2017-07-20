package sc

// Slew limits the slope of an input signal.
// The slope is expressed in units per second.
// For smoothing out control signals, take a look at Lag and VarLag.
type Slew struct {
	// In is the input signal
	In Input

	// Up is the maximum upward slope in units per second.
	Up Input

	// Dn is the maximum downward slope in units per second.
	Dn Input
}

func (s *Slew) defaults() {
	if s.In == nil {
		s.In = C(0)
	}
	if s.Up == nil {
		s.Up = C(1)
	}
	if s.Dn == nil {
		s.Dn = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (s Slew) Rate(rate int8) Input {
	CheckRate(rate)
	(&s).defaults()
	return NewInput("Slew", rate, 0, 1, s.In, s.Up, s.Dn)
}
