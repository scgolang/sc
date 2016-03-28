package sc

// Pan2 is a table-lookup sinewave oscillator
type Pan2 struct {
	// In is the input signal.
	In Input
	// Pos is the pan position, -1 is hard left, +1 is hard right.
	Pos Input
}

func (pan *Pan2) defaults() {
	if pan.In == nil {
		panic("Pan2 requires an input")
	}
	if pan.Pos == nil {
		pan.Pos = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (pan Pan2) Rate(rate int8) Input {
	CheckRate(rate)
	(&pan).defaults()
	return UgenInput("Pan2", rate, 0, 1, pan.In, pan.Pos)
}
