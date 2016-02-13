package sc

// Line generates a line from the start value to the end value
type Line struct {
	Start Input
	End   Input
	Dur   Input
	Done  int
}

func (line *Line) defaults() {
	if line.Start == nil {
		line.Start = C(0)
	}
	if line.End == nil {
		line.End = C(1)
	}
	if line.Dur == nil {
		line.Dur = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (line Line) Rate(rate int8) Input {
	CheckRate(rate)
	(&line).defaults()
	return UgenInput("Line", rate, 0, 1, line.Start, line.End, line.Dur, C(float32(line.Done)))
}
