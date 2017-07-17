package sc

// LocalBuf allocates a buffer local to a synth node.
type LocalBuf struct {
	NumChannels Input
	NumFrames   Input
}

func (lb *LocalBuf) defaults() {
	if lb.NumChannels == nil {
		lb.NumChannels = C(1)
	}
	if lb.NumFrames == nil {
		lb.NumFrames = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (lb LocalBuf) Rate(rate int8) Input {
	CheckRate(rate)
	(&lb).defaults()
	mlb := maxLocalBufs{}.Rate(rate)
	return NewInput("LocalBuf", rate, 0, 1, lb.NumChannels, lb.NumFrames, mlb)
}
