package sc

// LeakDC removes DC offset from a signal.
type LeakDC struct {
	// In is the input signal.
	In Input

	// Coeff is the leak coefficient.
	Coeff Input

	rate int8
}

func (ldc *LeakDC) defaults() {
	if ldc.Coeff == nil {
		switch ldc.rate {
		case AR:
			ldc.Coeff = C(0.995)
		default:
			ldc.Coeff = C(0.9)
		}
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
// If the input signal is nil this method will cause a runtime panic.
func (ldc LeakDC) Rate(rate int8) Input {
	CheckRate(rate)
	if ldc.In == nil {
		panic("LeakDC requires an input signal")
	}
	ldc.rate = rate
	(&ldc).defaults()
	return NewInput("LeakDC", rate, 0, 1, ldc.In, ldc.Coeff)
}
