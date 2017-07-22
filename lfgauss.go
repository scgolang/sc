package sc

// LFGauss is a non-band-limited gaussian function oscillator.
// Output ranges from minval to 1.
// LFGauss implements the formula:
//     f(x) = exp(squared(x - iphase) / (-2.0 * squared(width)))
// where x is to vary in the range -1 to 1 over the period dur.
// minval is the initial value at -1.
type LFGauss struct {
	Duration Input
	Width    Input
	IPhase   Input
	Loop     Input
	Done     int
}

func (l *LFGauss) defaults() {
	if l.Duration == nil {
		l.Duration = C(1)
	}
	if l.Width == nil {
		l.Width = C(0.1)
	}
	if l.IPhase == nil {
		l.IPhase = C(0)
	}
	if l.Loop == nil {
		l.Loop = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (l LFGauss) Rate(rate int8) Input {
	CheckRate(rate)
	(&l).defaults()
	done := C(float32(l.Done))
	return NewInput("LFGauss", rate, 0, 1, l.Duration, l.Width, l.IPhase, l.Loop, done)
}
