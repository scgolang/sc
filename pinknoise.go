package sc

// PinkNoise generates noise whose spectrum falls off
// in power by 3 dB per octave.
// This gives equal power over the span of each octave.
// This version gives 8 octaves of pink noise.
type PinkNoise struct {
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (pink PinkNoise) Rate(rate int8) Input {
	CheckRate(rate)
	return UgenInput("PinkNoise", rate, 0, 1)
}
