package sc

// Crackle is a noise generator based on a chaotic function.
type Crackle struct {
	Chaos Input
}

func (crackle *Crackle) defaults() {
	if crackle.Chaos == nil {
		crackle.Chaos = C(1.5)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (crackle Crackle) Rate(rate int8) Input {
	CheckRate(rate)
	(&crackle).defaults()
	return UgenInput("Crackle", rate, 0, 1, crackle.Chaos)
}
