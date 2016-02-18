package sc

import "fmt"

// Delay is a simple delay line.
type Delay struct {
	// Interpolation determines the type of interpolation.
	Interpolation Interpolation
	// In the input signal
	In Input
	// MaxDelayTime maximum delay time in seconds, which is used
	// to initialize the delay buffer size.
	MaxDelayTime Input
	// DelayTime delay time in seconds
	DelayTime Input
}

// Rate creates a new ugen at a specific rate.
// If an In signal is not provided this method will
// trigger a runtime panic.
func (delay Delay) Rate(rate int8) Input {
	CheckRate(rate)
	if delay.In == nil {
		panic("Delay expects In to not be nil")
	}
	if delay.MaxDelayTime == nil {
		delay.MaxDelayTime = C(0.2)
	}
	if delay.DelayTime == nil {
		delay.DelayTime = C(0.2)
	}
	switch delay.Interpolation {
	case InterpolationNone:
		return UgenInput("DelayN", rate, 0, 1, delay.In, delay.MaxDelayTime, delay.DelayTime)
	case InterpolationLinear:
		return UgenInput("DelayL", rate, 0, 1, delay.In, delay.MaxDelayTime, delay.DelayTime)
	case InterpolationCubic:
		return UgenInput("DelayC", rate, 0, 1, delay.In, delay.MaxDelayTime, delay.DelayTime)
	default:
		panic(fmt.Errorf("invalid interpolation %d", delay.Interpolation))
	}
}
