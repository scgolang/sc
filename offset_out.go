package sc

// OffsetOut outputs a signal to a bus.
// Tthe sample offset within the bus is kept exactly; i.e. if the synth is scheduled to be started part way
// through a control cycle, OffsetOut will maintain the correct offset by buffering the output and
// delaying it until the exact time that the synth was scheduled for.
type OffsetOut struct {
	Bus      Input
	Channels Input
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (out OffsetOut) Rate(rate int8) Input {
	CheckRate(rate)
	return NewInput("OffsetOut", rate, 0, 1, out.Bus, out.Channels)
}
