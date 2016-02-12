package sc

// MouseY allpass delay with cubic interpolation
type MouseY struct {
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

func (self *MouseY) defaults() {
	if self.Min == nil {
		self.Min = C(0)
	}
	if self.Max == nil {
		self.Max = C(1)
	}
	if self.Warp == nil {
		self.Warp = C(0)
	}
	if self.Lag == nil {
		self.Lag = C(0.2)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (self MouseY) Rate(rate int8) Input {
	CheckRate(rate)
	(&self).defaults()
	return UgenInput("MouseY", rate, 0, 1, self.Min, self.Max, self.Warp, self.Lag)
}
