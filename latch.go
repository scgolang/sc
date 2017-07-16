package sc

// Latch is a sample and hold.
// It holds input signal value when triggered.
// Latch will output 0 until it receives its first trigger.
type Latch struct {
	// In is the input signal.
	In Input

	// Trig can be any signal.
	// A trigger happens when the signal changes from non-positive to positive.
	Trig Input
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
// If the input signal is nil this method will panic.
// If the trig signal is nil this method will panic.
func (l Latch) Rate(rate int8) Input {
	CheckRate(rate)
	if l.In == nil {
		panic("Latch requires an input signal")
	}
	if l.Trig == nil {
		panic("Latch requires an trig signal")
	}
	return NewInput("Latch", rate, 0, 1, l.In, l.Trig)
}
