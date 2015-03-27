package ugens

import . "github.com/briansorahan/sc/types"

// PinkNoise
type PinkNoise struct {
}

func (self PinkNoise) Rate(rate int8) Input {
	checkRate(rate)
	return UgenInput("PinkNoise", rate, 0)
}
