package sc

// Saw is a band-limited sawtooth wave generator.
type Saw struct {
	Freq Input
}

func (saw *Saw) defaults() {
	if saw.Freq == nil {
		saw.Freq = C(440)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (saw Saw) Rate(rate int8) Input {
	CheckRate(rate)
	(&saw).defaults()
	return UgenInput("Saw", rate, 0, 1, saw.Freq)
}
