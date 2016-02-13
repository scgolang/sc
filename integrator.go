package sc

// Integrator integrates an input signal with a leak.
// The formula used is
//
//     out(0) = in(0) + (coef * out(-1))
type Integrator struct {
	// In is the input signal
	In Input
	// Coef is the leak coefficient
	Coef Input
}

func (integrator *Integrator) defaults() {
	if integrator.Coef == nil {
		integrator.Coef = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If an In signal is not provided this method will
// trigger a runtime panic.
func (integrator Integrator) Rate(rate int8) Input {
	CheckRate(rate)
	if integrator.In == nil {
		panic("Integrator expects In to not be nil")
	}
	(&integrator).defaults()
	return UgenInput("Integrator", rate, 0, 1, integrator.In, integrator.Coef)
}
