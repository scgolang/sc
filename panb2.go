package sc

// PanB2 encodes a mono signal to 2-dimensional ambisonic B-format.
type PanB2 struct {
	// In is the input signal.
	In Input

	// Azimuth is the position around the circle from -1 to +1. -1 is behind,
	// -0.5 is left, 0 is forward, +0.5 is right, +1 is behind.
	Azimuth Input

	// Gain controls the amplitude of the output signal.
	Gain Input
}

func (pan *PanB2) defaults() {
	if pan.Azimuth == nil {
		pan.Azimuth = C(0)
	}
	if pan.Gain == nil {
		pan.Gain = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
// If the input signal is nil this method will panic.
func (pan PanB2) Rate(rate int8) Input {
	CheckRate(rate)
	if pan.In == nil {
		panic("PanB2 requires an input")
	}
	(&pan).defaults()
	return NewInput("PanB2", rate, 0, 3, pan.In, pan.Azimuth, pan.Gain)

}
