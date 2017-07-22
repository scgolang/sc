package sc

// Probability distributions for Gendy1.
var (
	DistLinear    = C(0)
	DistCauchy    = C(1)
	DistLogist    = C(2)
	DistHyperbCos = C(3)
	DistArcSine   = C(4)
	DistExpon     = C(5)
	DistSinus     = C(6)
)

// Gendy1 is an implementation of the dynamic stochastic synthesis generator
// conceived by Iannis Xenakis and described in Formalized Music
// (1992, Stuyvesant, NY: Pendragon Press) chapter 9 (pp 246-254)
// and chapters 13 and 14 (pp 289-322).
// The BASIC program in the book was written by Marie-Helene Serra
// so I think it helpful to credit her too.
// The program code has been adapted to avoid infinities in the probability distribution functions.
// The distributions are hard-coded in C but there is an option to have
// new amplitude or time breakpoints sampled from a continuous controller input.
// All parameters can be modulated at control rate except for initCPs which is used
// only at initialisation.
type Gendy1 struct {
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

	// Minimum allowed frequency of oscillation for the Gendy1 oscillator,
	// so gives the largest period the duration is allowed to take on.
	MinFreq Input

	// Maximum allowed frequency of oscillation for the Gendy1 oscillator,
	// so gives the smallest period the duration is allowed to take on.
	MaxFreq Input

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

	rate int8
}

func (g *Gendy1) defaults() {
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
	switch g.rate {
	case AR:
		if g.MinFreq == nil {
			g.MinFreq = C(440)
		}
		if g.MaxFreq == nil {
			g.MaxFreq = C(660)
		}
	case KR:
		if g.MinFreq == nil {
			g.MinFreq = C(20)
		}
		if g.MaxFreq == nil {
			g.MaxFreq = C(1000)
		}
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
func (g Gendy1) Rate(rate int8) Input {
	CheckRate(rate)
	g.rate = rate
	(&g).defaults()
	return NewInput("Gendy1", rate, 0, 1, g.AmpDist, g.DurDist, g.ADParam, g.DDParam, g.MinFreq, g.MaxFreq, g.AmpScale, g.DurScale, g.InitCPs, g.KNum)
}
