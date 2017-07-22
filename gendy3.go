package sc

// Gendy3 is a dynamic stochastic synthesis generator.
// See Gendy1 for background.
// This variant of GENDYN normalises the durations in each period
// to force oscillation at the desired pitch.
// The breakpoints still get perturbed as in Gendy1.
// There is some glitching in the oscillator caused by the stochastic effects:
// control points as they vary cause big local jumps of amplitude.
// Put ampscale and durscalelow to minimise the rate of this.
type Gendy3 struct {
	// Choice of probability distribution for the next perturbation
	// of the amplitude of a control point.
	AmpDist Input

	// Choice of distribution for the perturbation of the current inter control point duration.
	DurDist Input

	// A parameter for the shape of the amplitude probability distribution,
	// requires values in the range 0.0001 to 1 (there are safety checks
	// in the code so don't worry too much if you want to modulate!).
	ADParam Input

	// A parameter for the shape of the duration probability distribution,
	// requires values in the range 0.0001 to 1.
	DDParam Input

	// Oscillation frquency.
	Freq Input

	// Normally 0.0 to 1.0, multiplier for the distribution's delta value for amplitude.
	// An ampscale of 1.0 allows the full range of -1 to 1 for a change of amplitude.
	AmpScale Input

	// Normally 0.0 to 1.0, multiplier for the distribution's delta value for duration.
	// An ampscale of 1.0 allows the full range of -1 to 1 for a change of duration.
	DurScale Input

	// Initialise the number of control points in the memory.
	// Xenakis specifies 12.
	// There would be this number of control points per cycle of the oscillator,
	// though the oscillator's period will constantly change due to the duration distribution.
	InitCPs Input

	// Current number of utilised control points, allows modulation.
	KNum Input
}

func (g *Gendy3) defaults() {
	if g.AmpDist == nil {
		g.AmpDist = DistCauchy
	}
	if g.DurDist == nil {
		g.DurDist = DistCauchy
	}
	if g.ADParam == nil {
		g.ADParam = C(1)
	}
	if g.DDParam == nil {
		g.DDParam = C(1)
	}
	if g.Freq == nil {
		g.Freq = C(440)
	}
	if g.AmpScale == nil {
		g.AmpScale = C(0.5)
	}
	if g.DurScale == nil {
		g.DurScale = C(0.5)
	}
	if g.InitCPs == nil {
		g.InitCPs = C(12)
	}
	if g.KNum == nil {
		g.KNum = C(12)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (g Gendy3) Rate(rate int8) Input {
	CheckRate(rate)
	(&g).defaults()
	return NewInput("Gendy3", rate, 0, 1, g.AmpDist, g.DurDist, g.ADParam, g.DDParam, g.Freq, g.AmpScale, g.DurScale, g.InitCPs, g.KNum)
}
