package sc

// PulseCount counts pulses.
// Each trigger increments a counter which is output as a signal.
type PulseCount struct {
	// Trig is the trigger signal.
	// A trigger happens when the signal changes from non-positive to positive.
	Trig Input

	// Reset resets the counter to zero when triggered.
	Reset Input
}

func (pc *PulseCount) defaults() {
	if pc.Trig == nil {
		pc.Trig = C(0)
	}
	if pc.Reset == nil {
		pc.Reset = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (pc PulseCount) Rate(rate int8) Input {
	CheckRate(rate)
	(&pc).defaults()
	return NewInput("PulseCount", rate, 0, 1, pc.Trig, pc.Reset)
}
