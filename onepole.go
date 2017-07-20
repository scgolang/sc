package sc

// OnePole is a one pole filter. Implements the formula:
//     out(i) = ((1 - abs(coef)) * in(i)) + (coef * out(i-1))
type OnePole struct {
	// In is the input signal.
	In Input

	// Coeff is the feedback coefficient. Should be between -1 and +1
	Coeff Input
}

func (o *OnePole) defaults() {
	if o.Coeff == nil {
		o.Coeff = C(0.5)
	}
	if o.In == nil {
		o.In = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (o OnePole) Rate(rate int8) Input {
	CheckRate(rate)
	(&o).defaults()
	return NewInput("OnePole", rate, 0, 1, o.In, o.Coeff)
}
