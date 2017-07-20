package sc

// Lag is essentially the same as OnePole except that instead of
// supplying the coefficient directly, it is calculated from a 60 dB lag time.
// This is the time required for the filter to converge to within 0.01% of a value.
// This is useful for smoothing out control signals.
type Lag struct {
	// In is the input signal.
	In Input

	// LagTime is the 60 dB lag time in seconds.
	LagTime Input
}

func (l *Lag) defaults() {
	if l.LagTime == nil {
		l.LagTime = C(0.1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
// If the input signal is nil this method will panic.
func (l Lag) Rate(rate int8) Input {
	CheckRate(rate)
	if l.In == nil {
		panic("Lag requires an input signal")
	}
	return NewInput("Lag", rate, 0, 1, l.In, l.LagTime)
}
