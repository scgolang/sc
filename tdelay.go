package sc

// TDelay delays a trigger by a given time.
// Any triggers which arrive in the time between an input trigger and its
// delayed output, are ignored.
type TDelay struct {
	// In is the input trigger signal.
	In Input

	// Dur is the delay time in seconds.
	Dur Input
}

func (t *TDelay) defaults() {
	if t.In == nil {
		t.In = C(0)
	}
	if t.Dur == nil {
		t.Dur = C(0.1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (t TDelay) Rate(rate int8) Input {
	CheckRate(rate)
	(&t).defaults()
	return NewInput("TDelay", rate, 0, 1, t.In, t.Dur)
}
