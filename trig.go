package sc

// Trig outputs the level of the triggering input for the specified duration
// when a nonpositive to positive transition occurs at the input.
// Otherwise it outputs zero.
type Trig struct {
	// In is the input trigger signal.
	// A trigger happens when the signal changes from non-positive to positive.
	In Input

	// Dur is the duration of the trigger output.
	Dur Input
}

func (t *Trig) defaults() {
	if t.In == nil {
		t.In = C(0)
	}
	if t.Dur == nil {
		t.Dur = C(0.1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (t Trig) Rate(rate int8) Input {
	CheckRate(rate)
	(&t).defaults()
	return NewInput("Trig", rate, 0, 1, t.In, t.Dur)
}
