package sc

import (
	"fmt"
)

// NoiseInterpolation defines different types of interpolation for LFNoise ugens.
type NoiseInterpolation int

// Noise interpolation types.
const (
	NoiseStep NoiseInterpolation = iota
	NoiseLinear
	NoiseQuadratic
)

// LFNoise generates random values at a rate given by
// the nearest integer division of the sample rate by the Freq input.
// The values will be interpolated according to the Interpolation input.
type LFNoise struct {
	Interpolation NoiseInterpolation
	Freq          Input
}

func (lfnoise *LFNoise) defaults() {
	if lfnoise.Freq == nil {
		lfnoise.Freq = C(500)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (lfnoise LFNoise) Rate(rate int8) Input {
	CheckRate(rate)
	(&lfnoise).defaults()
	switch lfnoise.Interpolation {
	case NoiseStep:
		return UgenInput("LFNoise0", rate, 0, 1, lfnoise.Freq)
	case NoiseLinear:
		return UgenInput("LFNoise1", rate, 0, 1, lfnoise.Freq)
	case NoiseQuadratic:
		return UgenInput("LFNoise2", rate, 0, 1, lfnoise.Freq)
	default:
		panic(fmt.Sprintf("Unknown Interpolation value for LFNoise: %d", lfnoise.Interpolation))
	}
}
