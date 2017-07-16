package sc

// PulseDivider outputs one impulse each time it receives
// a certain number of triggers at its input.
type PulseDivider struct {
	// Trig can be any signal.
	// A trigger happens when the signal changes from non-positive to positive.
	Trig Input

	// Div is the number of triggers to count before outputting an impulse.
	Div Input

	// Start is the starting value for the trigger count.
	// This lets you start somewhere in the middle of a count.
	// If start is negative it adds that many counts to the first time the output is triggered.
	Start Input
}

func (pd *PulseDivider) defaults() {
	if pd.Trig == nil {
		pd.Trig = C(0)
	}
	if pd.Div == nil {
		pd.Div = C(2)
	}
	if pd.Start == nil {
		pd.Start = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (pd PulseDivider) Rate(rate int8) Input {
	CheckRate(rate)
	(&pd).defaults()
	return NewInput("PulseDivider", rate, 0, 1, pd.Trig, pd.Div, pd.Start)
}
