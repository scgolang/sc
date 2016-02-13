package sc

// WhiteNoise generates noise whose spectrum has equal power at all frequencies.
type WhiteNoise struct {
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (white WhiteNoise) Rate(rate int8) Input {
	CheckRate(rate)
	return UgenInput("WhiteNoise", rate, 0, 1)
}
