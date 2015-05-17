package ugens

import . "github.com/scgolang/sc/types"

// PinkNoise
type PinkNoise struct {
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (self PinkNoise) Rate(rate int8) Input {
	checkRate(rate)
	return UgenInput("PinkNoise", rate, 0, 1)
}
