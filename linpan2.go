package sc

// LinPan2 is a two channel linear panner.
// The signal is lowered as it pans from left (or right) to center using a straight line
// from 1 (left or right) to 0.5 (center) for a 6dB reduction in the middle.
// A problem inherent to linear panning is that the perceived volume of the signal drops in the middle.
// Pan2 solves this by taking the square root of the linear scaling factor going from 1
// (left or right) to 0.5.sqrt (~=0.707) in the center, which is about 3dB reduction.
// This is equal power panning. LinPan2 sounds more like the Rhodes tremolo than Pan2.
type LinPan2 struct {
	In    Input
	Pos   Input
	Level Input
}

func (l *LinPan2) defaults() {
	if l.Pos == nil {
		l.Pos = C(0)
	}
	if l.Level == nil {
		l.Level = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
// If the input signal is nil this method will panic.
func (l LinPan2) Rate(rate int8) Input {
	CheckRate(rate)
	if l.In == nil {
		panic("LinPan2 requires an input signal")
	}
	(&l).defaults()
	in := NewInput("LinPan2", rate, 0, 1, l.In, l.Pos, l.Level)
	return Multi(in, in)
}
