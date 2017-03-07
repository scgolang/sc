package sc

// Rand generates a single random float value in uniform distribution from Lo to Hi.
// It generates this when the SynthDef first starts playing,
// and remains fixed for the duration of the synth's existence.
type Rand struct {
	Lo, Hi Input
}

func (rand *Rand) defaults() {
	if rand.Lo == nil {
		rand.Lo = C(0)
	}
	if rand.Hi == nil {
		rand.Hi = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (rand Rand) Rate(rate int8) Input {
	CheckRate(rate)
	(&rand).defaults()
	return UgenInput("Rand", rate, 0, 1, rand.Lo, rand.Hi)
}
