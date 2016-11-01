package sc

// Limiter limits the input amplitude to the given level.
// Limiter will not overshoot like Compander will,
// but it needs to look ahead in the audio.
// Thus there is a delay equal to twice the value of the dur parameter.
// Limiter, unlike Compander, is completely transparent for an in range signal.
// The Rate method of Limiter will panic if In is nil.
type Limiter struct {
	In    Input // The signal to be processed.
	Level Input // The peak output amplitude level to which to normalize the input.
	Dur   Input // The buffer delay time. Shorter times will produce smaller delays and quicker transient response times, but may introduce amplitude modulation artifacts.
}

func (limiter *Limiter) defaults() {
	if limiter.In == nil {
		panic("Limiter needs an input")
	}
	if limiter.Level == nil {
		limiter.Level = C(1)
	}
	if limiter.Dur == nil {
		limiter.Dur = C(0.01)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (limiter Limiter) Rate(rate int8) Input {
	CheckRate(rate)
	(&limiter).defaults()
	return UgenInput("Limiter", rate, 0, 1, limiter.In, limiter.Level, limiter.Dur)
}
