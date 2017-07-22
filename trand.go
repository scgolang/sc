package sc

// TRand generates a random float value in uniform distribution from
// Lo to Hi each time Trig changes from nonpositive to positive values.
type TRand struct {
	Lo   Input
	Hi   Input
	Trig Input
}

func (t *TRand) defaults() {
	if t.Lo == nil {
		t.Lo = C(0)
	}
	if t.Hi == nil {
		t.Hi = C(1)
	}
	if t.Trig == nil {
		t.Trig = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (t TRand) Rate(rate int8) Input {
	CheckRate(rate)
	(&t).defaults()
	return NewInput("TRand", rate, 0, 1, t.Lo, t.Hi, t.Trig)
}
