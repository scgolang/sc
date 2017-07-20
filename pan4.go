package sc

// Pan4 is a four channel equal power panner.
// Outputs are in order LeftFront, RightFront, LeftBack, RightBack.
type Pan4 struct {
	// In is the input signal.
	In Input

	// XPos is the X pan position from -1 to +1 (left to right).
	XPos Input

	// YPos is the Y pan position from -1 to +1 (back to front).
	YPos Input

	// Level is a control rate level input.
	Level Input
}

func (pan *Pan4) defaults() {
	if pan.XPos == nil {
		pan.XPos = C(0)
	}
	if pan.YPos == nil {
		pan.YPos = C(0)
	}
	if pan.Level == nil {
		pan.Level = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
// If the input signal is nil this method will panic.
func (pan Pan4) Rate(rate int8) Input {
	CheckRate(rate)
	if pan.In == nil {
		panic("Pan4 requires an input")
	}
	(&pan).defaults()
	in := NewInput("Pan4", rate, 0, 1, pan.In, pan.XPos, pan.YPos, pan.Level)
	return Multi(in, in, in, in)

}
