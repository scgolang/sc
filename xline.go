package sc

// XLine generates an exponential curve from the start value to
// the end value. Both the start and end values must be non-zero
// and have the same sign.
type XLine struct {
	Start Input
	End   Input
	Dur   Input
	Done  int
}

func (xline *XLine) defaults() {
	if xline.Start == nil {
		xline.Start = C(1)
	}
	if xline.End == nil {
		xline.End = C(2)
	}
	if xline.Dur == nil {
		xline.Dur = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (xline XLine) Rate(rate int8) Input {
	CheckRate(rate)
	(&xline).defaults()
	return UgenInput("XLine", rate, 0, 1, xline.Start, xline.End, xline.Dur, C(float32(xline.Done)))
}
