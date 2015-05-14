package ugens

import . "github.com/scgolang/sc/types"

// WhiteNoise
type WhiteNoise struct {
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (self WhiteNoise) Rate(rate int8) Input {
	checkRate(rate)
	return UgenInput("WhiteNoise", rate, 0)
}
