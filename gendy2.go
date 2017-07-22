package sc

// Gendy2 is a dynamic stochastic synthesis generator.
// See Gendy1 for background.
// This variant of GENDYN is closer to that presented in Hoffmann,
// Peter. (2000) The New GENDYN Program. Computer Music Journal 24:2, pp 31-38.
type Gendy2 struct {
	Gendy1

	// Parameter for Lehmer random number generator perturbed by Xenakis as in ((old*a)+c)%1.0
	A Input

	// Parameter for Lehmer random number generator perturbed by Xenakis.
	C Input
}

func (g *Gendy2) defaults() {
	(&g.Gendy1).defaults()

	if g.A == nil {
		g.A = C(1.17)
	}
	if g.C == nil {
		g.C = C(0.31)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (g Gendy2) Rate(rate int8) Input {
	CheckRate(rate)
	g.rate = rate
	(&g).defaults()
	return NewInput("Gendy2", rate, 0, 1, g.AmpDist, g.DurDist, g.ADParam, g.DDParam, g.MinFreq, g.MaxFreq, g.AmpScale, g.DurScale, g.InitCPs, g.KNum, g.A, g.C)
}
