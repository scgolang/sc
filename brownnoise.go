package sc

// BrownNoise generates noise whose spectrum falls off in power by 6 dB per octave
type BrownNoise struct{}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (brown BrownNoise) Rate(rate int8) Input {
	CheckRate(rate)
	return UgenInput("BrownNoise", rate, 0, 1)
}
