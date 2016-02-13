package sc

import (
	"fmt"
	"sort"
)

const (
	// CurveStep is a flat envelope segment.
	CurveStep = C(0)
	// CurveLinear is a linear envelope segment.
	CurveLinear = C(1)
	// CurveExp is an exponential envelope segment.
	CurveExp = C(2)
	// CurveSine is a sinusoidal shaped envelope segment.
	CurveSine = C(3)
	// CurveWelch is a sinusoidal segment shaped like the sides of a welch window.
	CurveWelch = C(4)
	// CurveCustom is an undocumented (on doc.sccode.org) envelope segment shape.
	CurveCustom = C(5)
	// CurveSquared is a squared envelope segment.
	CurveSquared = C(6)
	// CurveCubed is a cubed envelope segment.
	CurveCubed = C(7)
)

// Env is a specification for a breakpoint envelope
type Env struct {
	// Levels is the array of levels
	Levels []Input
	// Times is the array of durations (in seconds).
	// The length of this array should be one less than the
	// Levels array.
	Times []Input
	// CurveTypes determines the shape of each envelope segment.
	CurveTypes  []Input
	Curvature   Input
	ReleaseNode Input
	LoopNode    Input
}

func (e *Env) defaults() {
	if e.Levels == nil {
		e.Levels = []Input{C(0), C(1), C(0)}
	}
	if e.Times == nil {
		e.Times = []Input{C(1), C(1)}
	}
	if e.CurveTypes == nil {
		numSegments := len(e.Times)
		e.CurveTypes = make([]Input, numSegments)

		if e.Curvature == nil {
			for i := 0; i < numSegments; i++ {
				e.CurveTypes[i] = CurveLinear
			}
		} else {
			for i := 0; i < numSegments; i++ {
				e.CurveTypes[i] = e.Curvature
			}
		}
		e.Curvature = C(0)
	}
}

// Inputs returns the array of inputs that defines the Env.
func (e Env) Inputs() []Input {
	// This is how the inputs array is constructed:
	// 0, 3, -99, -99, -- starting level, num segments, releaseNode, loopNode
	// 1, 0.1, 5, 4, -- first segment: level, time, curve type, curvature
	// 0.5, 1, 5, -4, -- second segment: level, time, curve type, curvature
	// 0, 0.2, 5, 4 -- and so on
	var (
		lc = len(e.CurveTypes)
		lt = len(e.Times)
	)
	if lc != lt {
		panic(fmt.Errorf("%d curve types != %d times", lc, lt))
	}
	(&e).defaults()

	var (
		numSegments = len(e.Levels) - 1
		arr         = make([]Input, 4*(numSegments+1))
	)
	arr[0] = e.Levels[0]
	arr[1] = C(numSegments)
	arr[2] = e.ReleaseNode
	arr[3] = e.LoopNode
	for i, t := range e.Times {
		arr[(4*i)+4] = e.Levels[i+1]
		arr[(4*i)+5] = t
		arr[(4*i)+6] = e.CurveTypes[i]
		arr[(4*i)+7] = e.Curvature
	}
	return arr
}

// EnvLinen creates a new envelope which has a trapezoidal shape
type EnvLinen struct {
	Attack, Sustain, Release, Level, CurveType Input
}

func (linen *EnvLinen) defaults() {
	if linen.Attack == nil {
		linen.Attack = C(0.01)
	}
	if linen.Sustain == nil {
		linen.Sustain = C(1)
	}
	if linen.Release == nil {
		linen.Release = C(1)
	}
	if linen.Level == nil {
		linen.Level = C(1)
	}
	if linen.CurveType == nil {
		linen.CurveType = C(1)
	}
}

// Inputs returns the array of inputs that defines the Env.
func (linen EnvLinen) Inputs() []Input {
	(&linen).defaults()

	var (
		levels = []Input{C(0), linen.Level, linen.Level, C(0)}
		times  = []Input{linen.Attack, linen.Sustain, linen.Release}
		ct     = linen.CurveType
		cts    = []Input{ct, ct, ct}
	)
	return Env{levels, times, cts, C(0), C(-99), C(-99)}.Inputs()
}

// EnvTriangle creates a new envelope that has a triangle shape
type EnvTriangle struct {
	Dur, Level Input
}

func (tri *EnvTriangle) defaults() {
	if tri.Dur == nil {
		tri.Dur = C(1)
	}
	if tri.Level == nil {
		tri.Level = C(1)
	}
}

// Inputs returns the array of inputs that defines the Env.
func (tri EnvTriangle) Inputs() []Input {
	(&tri).defaults()

	var (
		levels = []Input{C(0), tri.Level, C(0)}
		d      = tri.Dur.Mul(C(0.5))
		times  = []Input{d, d}
		cts    = []Input{CurveLinear, CurveLinear}
	)
	return Env{levels, times, cts, C(0), C(-99), C(-99)}.Inputs()
}

// EnvSine creates a new envelope which has a hanning window shape
type EnvSine struct {
	Dur, Level Input
}

func (sine *EnvSine) defaults() {
	if sine.Dur == nil {
		sine.Dur = C(1)
	}
	if sine.Level == nil {
		sine.Level = C(1)
	}
}

// Inputs returns the array of inputs that defines the Env.
func (sine EnvSine) Inputs() []Input {
	(&sine).defaults()

	var (
		levels = []Input{C(0), sine.Level, C(0)}
		d      = sine.Dur.Mul(C(0.5))
		times  = []Input{d, d}
		cts    = []Input{CurveSine, CurveSine}
	)
	return Env{levels, times, cts, C(0), C(-99), C(-99)}.Inputs()
}

// EnvPerc creates a new envelope that has a percussive shape
type EnvPerc struct {
	Attack, Release, Level, Curvature Input
}

func (perc *EnvPerc) defaults() {
	if perc.Attack == nil {
		perc.Attack = C(0.01)
	}
	if perc.Release == nil {
		perc.Release = C(1)
	}
	if perc.Level == nil {
		perc.Level = C(1)
	}
	if perc.Curvature == nil {
		perc.Curvature = C(-4)
	}
}

// Inputs returns the array of inputs that defines the Env.
func (perc EnvPerc) Inputs() []Input {
	(&perc).defaults()

	var (
		levels = []Input{C(0), perc.Level, C(0)}
		times  = []Input{perc.Attack, perc.Release}
		cts    = []Input{CurveCustom, CurveCustom}
		crv    = perc.Curvature
	)
	return Env{levels, times, cts, crv, C(-99), C(-99)}.Inputs()
}

// Pairs are pairs of floats: the first float is time,
// the second is level.
// They get sorted by time.
type Pairs [][2]float32

func (p Pairs) Len() int {
	return len(p)
}

func (p Pairs) Less(i, j int) bool {
	return p[i][0] < p[j][0]
}

func (p Pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// EnvPairs creates a new envelope from coordinates/pairs
type EnvPairs struct {
	Pairs     Pairs
	CurveType C
}

// Inputs returns the array of inputs that defines the Env.
func (pairs EnvPairs) Inputs() []Input {
	sort.Sort(&(pairs.Pairs))

	var (
		lp     = len(pairs.Pairs)
		levels = make([]Input, lp)
		times  = make([]Input, lp-1)
		cts    = make([]Input, lp-1)
	)
	for i, p := range pairs.Pairs {
		levels[i] = C(p[1])
		if i > 0 {
			times[i-1] = C(p[0] - pairs.Pairs[i-1][0])
			cts[i-1] = pairs.CurveType
		}
	}
	return Env{levels, times, cts, C(0), C(-99), C(-99)}.Inputs()
}

// TLC (time, level, curve) triplet
type TLC struct {
	Time, Level float32
	Curve       C
}

// EnvTLC creates a new envelope from an array of (time, level, curve) triplets
// This is renamed from Env.xyc.
// The Curve value of the last triplet is ignored.
type EnvTLC []TLC

func (tlc EnvTLC) Len() int {
	return len(tlc)
}

func (tlc EnvTLC) Less(i, j int) bool {
	return tlc[i].Time < tlc[j].Time
}

func (tlc EnvTLC) Swap(i, j int) {
	tlc[i], tlc[j] = tlc[j], tlc[i]
}

// Inputs returns the array of inputs that defines the Env.
func (tlc EnvTLC) Inputs() []Input {
	sort.Sort(tlc)

	var (
		lp     = len(tlc)
		levels = make([]Input, lp)
		times  = make([]Input, lp-1)
		cts    = make([]Input, lp-1)
	)
	for i, t := range tlc {
		levels[i] = C(t.Level)
		if i > 0 {
			times[i-1] = C(t.Time - tlc[i-1].Time)
			cts[i-1] = tlc[i-1].Curve
		}
	}
	return Env{levels, times, cts, C(0), C(-99), C(-99)}.Inputs()
}

// EnvADSR represents the ever-popular ADSR envelope
type EnvADSR struct {
	A, D, S, R, Peak, Curve, Bias Input
}

func (adsr *EnvADSR) defaults() {
	if adsr.A == nil {
		adsr.A = C(0.01)
	}
	if adsr.D == nil {
		adsr.D = C(0.3)
	}
	if adsr.S == nil {
		adsr.S = C(0.5)
	}
	if adsr.R == nil {
		adsr.R = C(1)
	}
	if adsr.Peak == nil {
		adsr.Peak = C(1)
	}
	if adsr.Curve == nil {
		adsr.Curve = C(-4)
	}
	if adsr.Bias == nil {
		adsr.Bias = C(0)
	}
}

// Inputs returns the array of inputs that defines the Env.
func (adsr EnvADSR) Inputs() []Input {
	(&adsr).defaults()

	levels := []Input{
		C(0).Add(adsr.Bias),
		adsr.Peak.Add(adsr.Bias),
		adsr.S.Add(adsr.Bias),
		C(0).Add(adsr.Bias),
	}
	var (
		times = []Input{adsr.A, adsr.D, adsr.R}
		cts   = []Input{CurveCustom, CurveCustom, CurveCustom}
	)
	return Env{levels, times, cts, adsr.Curve, C(2), C(-99)}.Inputs()
}

// EnvDADSR is EnvADSR with its onset delayed by D seconds
type EnvDADSR struct {
	Delay, A, D, S, R, Peak, Curve, Bias Input
}

func (dadsr *EnvDADSR) defaults() {
	if dadsr.Delay == nil {
		dadsr.Delay = C(0.1)
	}
	if dadsr.A == nil {
		dadsr.A = C(0.01)
	}
	if dadsr.D == nil {
		dadsr.D = C(0.3)
	}
	if dadsr.S == nil {
		dadsr.S = C(0.5)
	}
	if dadsr.R == nil {
		dadsr.R = C(1)
	}
	if dadsr.Peak == nil {
		dadsr.Peak = C(1)
	}
	if dadsr.Curve == nil {
		dadsr.Curve = C(-4)
	}
	if dadsr.Bias == nil {
		dadsr.Bias = C(0)
	}
}

// Inputs returns the array of inputs that defines the Env.
func (dadsr EnvDADSR) Inputs() []Input {
	(&dadsr).defaults()

	levels := []Input{
		C(0),
		C(0).Add(dadsr.Bias),
		dadsr.Peak.Add(dadsr.Bias),
		dadsr.S.Add(dadsr.Bias),
		C(0).Add(dadsr.Bias),
	}
	var (
		times = []Input{dadsr.Delay, dadsr.A, dadsr.D, dadsr.R}
		cts   = []Input{CurveCustom, CurveCustom, CurveCustom, CurveCustom}
	)
	return Env{levels, times, cts, dadsr.Curve, C(3), C(-99)}.Inputs()
}

// EnvASR is an attack-sustain-release envelope
type EnvASR struct {
	A, S, R, Curve Input
}

func (asr *EnvASR) defaults() {
	if asr.A == nil {
		asr.A = C(0.01)
	}
	if asr.S == nil {
		asr.S = C(1)
	}
	if asr.R == nil {
		asr.R = C(1)
	}
	if asr.Curve == nil {
		asr.Curve = C(-4)
	}
}

// Inputs returns the array of inputs that defines the Env.
func (asr EnvASR) Inputs() []Input {
	(&asr).defaults()

	var (
		levels = []Input{C(0), asr.S, C(0)}
		times  = []Input{asr.A, asr.R}
		cts    = []Input{CurveCustom, CurveCustom}
	)
	return Env{levels, times, cts, asr.Curve, C(1), C(-99)}.Inputs()
}

// EnvCutoff creates an envelope with no attack segment.
// It simply sustains at the peak level until released.
type EnvCutoff struct {
	R, Level, CurveType Input
}

func (cutoff *EnvCutoff) defaults() {
	if cutoff.R == nil {
		cutoff.R = C(0.1)
	}
	if cutoff.Level == nil {
		cutoff.Level = C(1)
	}
	if cutoff.CurveType == nil {
		cutoff.CurveType = CurveLinear
	}
}

// Inputs returns the array of inputs that defines the Env.
func (cutoff EnvCutoff) Inputs() []Input {
	(&cutoff).defaults()

	var (
		levels = []Input{cutoff.Level, C(0)}
		times  = []Input{cutoff.R}
		cts    = []Input{cutoff.CurveType}
	)
	return Env{levels, times, cts, C(0), C(0), C(-99)}.Inputs()
}

// I don't understand Env.circle [briansorahan]
//
// Env.circle([0, 1, 0], [0.01, 0.5, 0.2]).asArray;
// => [ 0, 2, -99, -99, 1, 0.01, 1, 0, 0, 0.5, 1, 0 ]
//
// Shouldn't loopNode be set to one of the envelope breakpoints?
//
// type EnvCircle struct {
// 	Levels, Times []Input
// 	Curve         Input
// }
