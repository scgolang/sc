package sc

// Blip band-limited impulse generator
type Blip struct {
	// Freq frequency in Hz
	Freq Input
	// Harm the number of harmonics
	Harm Input
}

func (blip *Blip) defaults() {
	if blip.Freq == nil {
		blip.Freq = C(440)
	}
	if blip.Harm == nil {
		blip.Harm = C(200)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (blip Blip) Rate(rate int8) Input {
	CheckRate(rate)
	(&blip).defaults()
	return UgenInput("Blip", rate, 0, 1, blip.Freq, blip.Harm)
}
