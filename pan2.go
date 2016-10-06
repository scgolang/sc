package sc

// Pan2 is a two-channel equal power panner.
// It takes the square root of the linear scaling factor going from 1 (left or right)
// to 0.5.sqrt (~=0.707) in the center, which is about 3dB reduction.
// With linear panning (LinPan2) the signal is lowered as it approaches center
// using a straight line from 1 (left or right) to 0.5 (center) for a 6dB reduction
// in the middle. A problem inherent to linear panning is that the perceived volume
// of the signal drops in the middle. Pan2 solves this.
type Pan2 struct {
	// In is the input signal.
	In Input

	// Pos is the pan position, -1 is hard left, +1 is hard right.
	Pos Input

	// Level is a control rate level input.
	Level Input
}

func (pan *Pan2) defaults() {
	if pan.In == nil {
		panic("Pan2 requires an input")
	}
	if pan.Pos == nil {
		pan.Pos = C(0)
	}
	if pan.Level == nil {
		pan.Level = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (pan Pan2) Rate(rate int8) Input {
	CheckRate(rate)
	(&pan).defaults()
	return UgenInput("Pan2", rate, 0, 2, pan.In, pan.Pos, pan.Level)
}
