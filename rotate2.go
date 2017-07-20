package sc

// Rotate2 can be used for rotating an ambisonic B-format sound field around an axis.
// Rotate2 does an equal power rotation so it also works well on stereo sounds.
// It takes two audio inputs (x, y) and an angle control (pos).
// It outputs two channels (x, y). It computes this:
//     xout = cos(angle) * xin + sin(angle) * yin;
//     yout = cos(angle) * yin - sin(angle) * xin;
// where angle = pos * pi, so that -1 becomes -pi and +1 becomes +pi.
// This allows you to use an LFSaw to do continuous rotation around a circle.
type Rotate2 struct {
	// X and Y are the input signals.
	X, Y Input

	// Pos is the angle to rotate around the circle from -1 to +1.
	// -1 is 180 degrees, -0.5 is left, 0 is forward, +0.5 is right, +1 is behind.
	Pos Input
}

func (r *Rotate2) defaults() {
	if r.Pos == nil {
		r.Pos = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
// If X or Y is nil this method will panic.
// TODO: we have to be able to access individual outputs of a ugen to be able to implement
// the Rotate2 example from the SuperCollider help files.
func (r Rotate2) Rate(rate int8) Input {
	CheckRate(rate)
	if r.X == nil || r.Y == nil {
		panic("Rotate2 expects both X and Y signals to not be nil")
	}
	(&r).defaults()
	return NewInput("Rotate2", rate, 0, 1)
}
