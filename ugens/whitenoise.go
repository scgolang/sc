package ugens

import . "github.com/briansorahan/sc/types"

// WhiteNoise
type WhiteNoise struct {
}

func (self WhiteNoise) Rate(rate int8) Input {
	checkRate(rate)
	return UgenInput("WhiteNoise", rate, 0)
}
