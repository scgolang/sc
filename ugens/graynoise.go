package ugens

import . "github.com/scgolang/sc/types"

// GrayNoise
type GrayNoise struct {
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (self GrayNoise) Rate(rate int8) Input {
	CheckRate(rate)
	return UgenInput("GrayNoise", rate, 0, 1)
}
