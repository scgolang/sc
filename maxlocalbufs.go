package sc

// maxLocalBufs is used internally by LocalBuf,
// sets the maximum number of local buffers in a synth.
type maxLocalBufs struct {
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (maxLocalBufs) Rate(rate int8) Input {
	CheckRate(rate)
	return NewInput("MaxLocalBufs", rate, 0, 1, C(1))
}
