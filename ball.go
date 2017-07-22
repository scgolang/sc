package sc

// Ball models the path of a bouncing object that is reflected by a vibrating surface.
type Ball struct {
	In       Input
	Gravity  Input
	Damp     Input
	Friction Input
}

func (b *Ball) defaults() {
	if b.In == nil {
		b.In = C(0)
	}
	if b.Gravity == nil {
		b.Gravity = C(1)
	}
	if b.Damp == nil {
		b.Damp = C(0)
	}
	if b.Friction == nil {
		b.Friction = C(0.01)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (b Ball) Rate(rate int8) Input {
	CheckRate(rate)
	(&b).defaults()
	return NewInput("Ball", rate, 0, 1, b.In, b.Gravity, b.Damp, b.Friction)
}
