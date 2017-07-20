package sc

// PanAz is a four channel equal power panner.
// Outputs are in order LeftFront, RightFront, LeftBack, RightBack.
type PanAz struct {
	NumChans int

	// In is the input signal.
	In Input

	// Pos is the pan position.
	// Channels are evenly spaced over a cyclic period of 2.0
	// in pos with 0.0 equal to channel zero and 2.0/numChans equal to
	// channel 1, 4.0/numChans equal to channel 2, etc.
	Pos Input

	// Level is a control rate level input.
	Level Input

	// Width is the width of the panning envelope.
	// Nominally this is 2.0 which pans between pairs of adjacent speakers.
	// Width values greater than two will spread the pan over greater numbers of speakers.
	// Width values less than one will leave silent gaps between speakers.
	Width Input

	// Orientation should be zero if the front is a vertex of the polygon.
	// The first speaker will be directly in front.
	// Should be 0.5 if the front bisects a side of the polygon.
	// Then the first speaker will be the one left of center.
	Orientation Input
}

func (pan *PanAz) defaults() {
	if pan.Pos == nil {
		pan.Pos = C(0)
	}
	if pan.Level == nil {
		pan.Level = C(1)
	}
	if pan.Width == nil {
		pan.Width = C(2)
	}
	if pan.Orientation == nil {
		pan.Orientation = C(0.5)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
// If the input signal is nil this method will panic.
func (pan PanAz) Rate(rate int8) Input {
	CheckRate(rate)
	if pan.In == nil {
		panic("PanAz requires an input")
	}
	(&pan).defaults()
	return NewInput("PanAz", rate, 0, pan.NumChans, pan.In, pan.Pos, pan.Level, pan.Width, pan.Orientation)

}
