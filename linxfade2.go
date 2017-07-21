package sc

// LinXFade2 is a two channel linear crossfader.
type LinXFade2 struct {
	A     Input
	B     Input
	Pan   Input // Cross fade position from -1 to +1.
	Level Input
}

func (x *LinXFade2) defaults() {
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
func (x LinXFade2) Rate(rate int8) Input {
	CheckRate(rate)
	(&x).defaults()
	// Not sure why Level doesn't show up in sclang's synthdefs that use LinXFade2
	return NewInput("LinXFade2", rate, 0, 1, x.A, x.B, x.Pan)
}
