package sc

// Dust2 generates random impulses from -1 to +1
type Dust2 struct {
	// Density is the average number of impulses per second
	Density Input
}

func (dust2 *Dust2) defaults() {
	if dust2.Density == nil {
		dust2.Density = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (dust2 Dust2) Rate(rate int8) Input {
	CheckRate(rate)
	(&dust2).defaults()
	return UgenInput("Dust2", rate, 0, 1, dust2.Density)
}
