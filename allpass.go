package sc

import "fmt"

// Allpass is an allpass delay.
type Allpass struct {
	// Interpolation determines the type of interpolation (e.g. none, linear, cubic).
	Interpolation Interpolation

	// In is the input signal
	In Input

	// MaxDelayTime is maximum delay time in seconds.
	// This is used to initialize the delay buffer.
	MaxDelayTime Input

	// DelayTime time in seconds
	DelayTime Input

	// DecayTime time for the echoes to decay by 60dB.
	// If this is negative then the feedback coefficient will
	// be negative, thus emphasizing only odd harmonics
	// at a lower octave.
	DecayTime Input
}

func (apc *Allpass) defaults() {
	if apc.MaxDelayTime == nil {
		apc.MaxDelayTime = C(0.2)
	}
	if apc.DelayTime == nil {
		apc.DelayTime = C(0.2)
	}
	if apc.DecayTime == nil {
		apc.DecayTime = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
// If an In signal is not provided this method will panic.
// If Interpolation is set to an unsupported value this method will panic.
func (apc Allpass) Rate(rate int8) Input {
	if apc.In == nil {
		panic("Allpass expects In to not be nil")
	}
	CheckRate(rate)

	(&apc).defaults()

	switch apc.Interpolation {
	case InterpolationNone:
		return NewInput("AllpassN", rate, 0, 1, apc.In, apc.MaxDelayTime, apc.DelayTime, apc.DecayTime)
	case InterpolationLinear:
		return NewInput("AllpassL", rate, 0, 1, apc.In, apc.MaxDelayTime, apc.DelayTime, apc.DecayTime)
	case InterpolationCubic:
		return NewInput("AllpassC", rate, 0, 1, apc.In, apc.MaxDelayTime, apc.DelayTime, apc.DecayTime)
	default:
		panic(fmt.Errorf("invalid interpolation: %d", apc.Interpolation))
	}
}
