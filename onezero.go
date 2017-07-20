package sc

// OneZero is a one zero filter. Implements the formula:
//     out(i) = ((1 - abs(coef)) * in(i)) + (coef * in(i-1))
type OneZero struct {
	// In is the input signal.
	In Input

	// Coeff is the feed forward coefficient.
	// +0.5 makes a two point averaging filter (see also LPZ1 ).
	// -0.5 makes a differentiator (see also HPZ1 ).
	// +1 makes a single sample delay (see also Delay1 ).
	// -1 makes an inverted single sample delay.
	Coeff Input
}

func (o *OneZero) defaults() {
	if o.Coeff == nil {
		o.Coeff = C(0.5)
	}
	if o.In == nil {
		o.In = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (o OneZero) Rate(rate int8) Input {
	CheckRate(rate)
	(&o).defaults()
	return NewInput("OneZero", rate, 0, 1, o.In, o.Coeff)
}
