package sc

import "fmt"

// Comb is a delay line with cubic interpolation.
type Comb struct {
	// Interpolation determines the type of interpolation.
	Interpolation Interpolation
	// In the input signal
	In Input
	// MaxDelayTime maximum delay time in seconds, which is used
	// to initialize the delay buffer size.
	MaxDelayTime Input
	// DelayTime delay time in seconds
	DelayTime Input
	// DecayTime time for the echoes to decay by 60dB.
	// If this time is negative then the feedback coefficient will be
	// negative, thus emphasizing only odd harmonics an octave lower.
	DecayTime Input
}

// Rate creates a new ugen at a specific rate.
// If an In signal is not provided this method will
// trigger a runtime panic.
// If Interpolation is an unsupported value, then
// this method will trigger a runtime panic.
func (comb Comb) Rate(rate int8) Input {
	CheckRate(rate)
	if comb.In == nil {
		panic("Comb expects In to not be nil")
	}
	if comb.MaxDelayTime == nil {
		comb.MaxDelayTime = C(0.2)
	}
	if comb.DelayTime == nil {
		comb.DelayTime = C(0.2)
	}
	if comb.DecayTime == nil {
		comb.DecayTime = C(1)
	}
	switch comb.Interpolation {
	case InterpolationNone:
		return UgenInput("CombN", rate, 0, 1, comb.In, comb.MaxDelayTime, comb.DelayTime, comb.DecayTime)
	case InterpolationLinear:
		return UgenInput("CombL", rate, 0, 1, comb.In, comb.MaxDelayTime, comb.DelayTime, comb.DecayTime)
	case InterpolationCubic:
		return UgenInput("Comb", rate, 0, 1, comb.In, comb.MaxDelayTime, comb.DelayTime, comb.DecayTime)
	default:
		panic(fmt.Errorf("invalid interpolation: %d", comb.Interpolation))
	}
}
