package sc

// PinkNoise
type PinkNoise struct {
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (pink PinkNoise) Rate(rate int8) Input {
	CheckRate(rate)
	return UgenInput("PinkNoise", rate, 0, 1)
}
