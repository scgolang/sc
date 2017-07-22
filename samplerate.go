package sc

// SampleRate returns the current sample rate of the server.
type SampleRate struct{}

// Rate creates a new ugen at a specific rate.
// If rate is not IR this method will cause a runtime panic.
func (s SampleRate) Rate(rate int8) Input {
	if rate != IR {
		panic("SampleRate only supports IR")
	}
	return NewInput("SampleRate", rate, 0, 1)
}
