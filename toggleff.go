package sc

// ToggleFF toggles between 0 and 1 upon receiving a trigger.
type ToggleFF struct {
	Trig Input
}

func (t *ToggleFF) defaults() {
	if t.Trig == nil {
		t.Trig = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (t ToggleFF) Rate(rate int8) Input {
	CheckRate(rate)
	(&t).defaults()
	return NewInput("ToggleFF", rate, 0, 1, t.Trig)
}
