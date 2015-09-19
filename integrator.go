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

func (self *Integrator) defaults() {
	if self.Coef == nil {
		self.Coef = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If an In signal is not provided this method will
// trigger a runtime panic.
func (self Integrator) Rate(rate int8) Input {
	CheckRate(rate)
	if self.In == nil {
		panic("Integrator expects In to not be nil")
	}
	(&self).defaults()
	return UgenInput("Integrator", rate, 0, 1, self.In, self.Coef)
}
