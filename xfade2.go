package sc

// XFade2 is a two channel equal power crossfader.
type XFade2 struct {
	A     Input
	B     Input
	Pan   Input
	Level Input
}

func (x *XFade2) defaults() {
	if x.A == nil {
		x.A = C(0)
	}
	if x.B == nil {
		x.B = C(0)
	}
	if x.Pan == nil {
		x.Pan = C(0)
	}
	if x.Level == nil {
		x.Level = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (x XFade2) Rate(rate int8) Input {
	CheckRate(rate)
	(&x).defaults()
	return NewInput("XFade2", rate, 0, 1, x.A, x.B, x.Pan, x.Level)
}
