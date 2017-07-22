package sc

// Spring models the force of a resonating spring.
type Spring struct {
	// Modulated input force.
	In Input

	// Spring constant (incl. mass)
	Spring Input

	// Damping.
	Damp Input
}

func (b *Spring) defaults() {
	if b.In == nil {
		b.In = C(0)
	}
	if b.Spring == nil {
		b.Spring = C(1)
	}
	if b.Damp == nil {
		b.Damp = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (b Spring) Rate(rate int8) Input {
	CheckRate(rate)
	(&b).defaults()
	return NewInput("Spring", rate, 0, 1, b.In, b.Spring, b.Damp)
}
