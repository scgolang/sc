package sc

// SampleDur returns the current sample duration of the server.
// Equivalent to 1/SampleRate.
type SampleDur struct{}

// Rate creates a new ugen at a specific rate.
// If rate is not IR this method will cause a runtime panic.
func (s SampleDur) Rate(rate int8) Input {
	if rate != IR {
		panic("SampleDur only supports IR")
	}
	return NewInput("SampleDur", rate, 0, 1)
}
