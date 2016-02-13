package sc

// MouseX allpass delay with cubic interpolation
type MouseX struct {
	// Min is the value of this ugen's output when the
	// mouse is at the left edge of the screen
	Min Input
	// Max is the value of this ugen's output when the
	// mouse is at the right edge of the screen
	Max Input
	// Warp is the mapping curve. 0 is linear, 1 is exponential
	Warp Input
	// Lag factor to dezipper cursor movements
	Lag Input
}

func (m *MouseX) defaults() {
	if m.Min == nil {
		m.Min = C(0)
	}
	if m.Max == nil {
		m.Max = C(1)
	}
	if m.Warp == nil {
		m.Warp = C(0)
	}
	if m.Lag == nil {
		m.Lag = C(0.2)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (m MouseX) Rate(rate int8) Input {
	CheckRate(rate)
	(&m).defaults()
	return UgenInput("MouseX", rate, 0, 1, m.Min, m.Max, m.Warp, m.Lag)
}
