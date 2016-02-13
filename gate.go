package sc

// Gate allows an input signal value to pass when gate is
// positive, otherwise holds last value.
type Gate struct {
	// In is the input signal.
	In Input
	// Trig is the rigger signal. The output is held fixed when this is non-positive.
	Trig Input
}

func (gate *Gate) defaults() {
	if gate.Trig == nil {
		gate.Trig = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If an In signal is not provided this method will
// trigger a runtime panic.
func (gate Gate) Rate(rate int8) Input {
	CheckRate(rate)
	if gate.In == nil {
		panic("Gate expects In to not be nil")
	}
	(&gate).defaults()
	return UgenInput("Gate", rate, 0, 1, gate.In, gate.Trig)
}
