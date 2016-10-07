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

// shapeNames contains the names of the envelope curves.
// This mirrors the IdentityDictionary in SCClassLibrary/Common/Audio/Env.sc.
var shapeNames = map[string]int{
	"step":        0,
	"lin":         1,
	"linear":      1,
	"exp":         2,
	"exponential": 2,
	"sin":         3,
	"sine":        3,
	"wel":         4,
	"welch":       4,
	"sqr":         6,
	"squared":     6,
	"cub":         7,
	"cubed":       7,
}

// Env is a specification for a breakpoint envelope
type Env struct {
	// Levels is the array of levels
	Levels []Input
	// Times is the array of durations (in seconds).
	// The length of this array should be one less than the
	// Levels array.
	Times []Input
	// Curve determines the shape of each envelope segment.
	// This could be a string, a float, a slice of strings, or a slice of floats.
	Curve       interface{}
	ReleaseNode Input
	LoopNode    Input
}

func (env *Env) defaults() {
	if env.Levels == nil {
		env.Levels = []Input{C(0), C(1), C(0)}
	}
	if env.Times == nil {
		env.Times = []Input{C(1), C(1)}
	}
	if env.ReleaseNode == nil {
		env.ReleaseNode = C(-99)
	}
	if env.LoopNode == nil {
		env.LoopNode = C(-99)
	}
	if env.Curve == nil {
		env.Curve = "linear"
	}
}

// Inputs returns the array of inputs that defines the Env.
func (env Env) Inputs() []Input {
	(&env).defaults()

	// This is how the inputs array is constructed:
	// 0, 3, -99, -99, -- starting level, num segments, releaseNode, loopNode
	// 1, 0.1, 5, 4, -- first segment: level, time, curve type, curvature
	// 0.5, 1, 5, -4, -- second segment: level, time, curve type, curvature
	// 0, 0.2, 5, 4 -- and so on
	var (
		curvesArray = env.curvesArray()
		lc          = len(curvesArray)
		lt          = len(env.Times)
	)
	if lc != lt {
		panic(fmt.Errorf("%d curve types != %d times", lc, lt))
	}

	var (
		numSegments = len(env.Levels) - 1
		arr         = make([]Input, 4*(numSegments+1))
	)
	arr[0] = env.Levels[0]
	arr[1] = C(numSegments)
	arr[2] = env.ReleaseNode
	arr[3] = env.LoopNode
	for i, t := range env.Times {
		arr[(4*i)+4] = env.Levels[i+1]
		arr[(4*i)+5] = t
		arr[(4*i)+6] = shapeNumber(curvesArray[i])
		arr[(4*i)+7] = curveValue(curvesArray[i])
	}
	return arr
}

// curvesArray returns the Curve as an array.
func (env Env) curvesArray() []interface{} {
	switch val := env.Curve.(type) {
	case int, string, float64, Input:
		return arrayFromScalar(val, len(env.Times))
	case []int:
		return intsToEmpties(val)
	case []float64:
		return floatsToEmpties(val)
	case []string:
		return stringsToEmpties(val)
	case []Input:
		return inputsToEmpties(val)
	case []interface{}:
		return val
	default:
		panic(fmt.Sprintf("unsupported type for envelope curve: %T", env.Curve))
	}
}

// intsToEmpties converts a int slice to a slice of the empty interface.
func intsToEmpties(arr []int) []interface{} {
	ret := make([]interface{}, len(arr))
	for i, ii := range arr {
		ret[i] = ii
	}
	return ret
}

// floatsToEmpties converts a float slice to a slice of the empty interface.
func floatsToEmpties(arr []float64) []interface{} {
	ret := make([]interface{}, len(arr))
	for i, f := range arr {
		ret[i] = f
	}
	return ret
}

// stringsToEmpties converts a string slice to a slice of the empty interface.
func stringsToEmpties(arr []string) []interface{} {
	ret := make([]interface{}, len(arr))
	for i, s := range arr {
		ret[i] = s
	}
	return ret
}

// inputsToEmpties converts an Input slice to a slice of the empty interface.
func inputsToEmpties(arr []Input) []interface{} {
	ret := make([]interface{}, len(arr))
	for i, in := range arr {
		ret[i] = in
	}
	return ret
}

// arrayFromScalar converts a scalar value to a slice of the empty interface.
func arrayFromScalar(val interface{}, length int) []interface{} {
	ret := make([]interface{}, length)
	for i := range ret {
		ret[i] = val
	}
	return ret
}

// shapeNumber returns the mysterious shape number for the Env inputs array.
//
// The relevant sclang code (from Env.sc) is:
//
//     *shapeNumber { arg shapeName;
//             ^shapeName.asArray.collect { |name|
//                     var shape;
//                     if(name.isValidUGenInput) { 5 } {
//                             shape = shapeNames.at(name);
//                             if(shape.isNil) { Error("Env shape not defined.").throw };
//                             shape
//                     }
//             }.unbubble
//     }
//
// Note: we don't handle the case where shapeValue is an array since we force
// this func to be called with either string, float64, or Input.
//
func shapeNumber(shapeValue interface{}) Input {
	if _, ok := isValidUgenInput(shapeValue); ok {
		return C(5)
	}
	switch val := shapeValue.(type) {
	case int:
		return C(float32(val))
	case float64:
		return C(float32(val))
	case float32:
		return C(val)
	case Input:
		return val
	case string:
		shapeNum, ok := shapeNames[val]
		if !ok {
			panic(fmt.Sprintf("invalid curve: %s", val))
		}
		return C(shapeNum)
	default:
		panic(fmt.Sprintf("invalid curve type (must be float, string, or Input): %T", shapeValue))
	}
}

// isValidUgenInput returns false if val is not a valid ugen input and true otherwise.
func isValidUgenInput(val interface{}) (Input, bool) {
	switch x := val.(type) {
	case int:
		return C(float32(x)), true
	case float64:
		return C(float32(x)), true
	case float32:
		return C(x), true
	case Input:
		return x, true
	default:
		return nil, false
	}
}

// curveValue returns the mysterious curve value for the Env inputs array.
//
// The relevant sclang code (from Env.sc) is:
//
// curveValue { arg curve;
//         ^if(curve.isSequenceableCollection) {
//                 curve.collect { |x|
//                         if(x.isValidUGenInput) { x } { 0 }
//                 }
//         } {
//                 if(curve.isValidUGenInput) { curve } { 0 }
//         }
// }
func curveValue(curve interface{}) Input {
	if input, ok := isValidUgenInput(curve); ok {
		return input
	}
	return C(0)
}

// EnvLinen creates a new envelope which has a trapezoidal shape
type EnvLinen struct {
	Attack, Sustain, Release, Level Input
	Curve                           interface{}
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
	if linen.Curve == nil {
		linen.Curve = "lin"
	}
}

// Inputs returns the array of inputs that defines the Env.
func (linen EnvLinen) Inputs() []Input {
	(&linen).defaults()

	return Env{
		Levels: []Input{C(0), linen.Level, linen.Level, C(0)},
		Times:  []Input{linen.Attack, linen.Sustain, linen.Release},
		Curve:  linen.Curve,
	}.Inputs()
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

	d := tri.Dur.Mul(C(0.5))

	return Env{
		Levels: []Input{C(0), tri.Level, C(0)},
		Times:  []Input{d, d},
	}.Inputs()
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

	d := sine.Dur.Mul(C(0.5))

	return Env{
		Levels: []Input{C(0), sine.Level, C(0)},
		Times:  []Input{d, d},
		Curve:  "sine",
	}.Inputs()
}

// EnvPerc creates a new envelope that has a percussive shape
type EnvPerc struct {
	Attack, Release, Level, Curve Input
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
	if perc.Curve == nil {
		perc.Curve = C(-4)
	}
}

// Inputs returns the array of inputs that defines the Env.
func (perc EnvPerc) Inputs() []Input {
	(&perc).defaults()

	return Env{
		Levels: []Input{C(0), perc.Level, C(0)},
		Times:  []Input{perc.Attack, perc.Release},
		Curve:  []Input{perc.Curve, perc.Curve},
	}.Inputs()
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
	Pairs Pairs
	Curve interface{}
}

// Inputs returns the array of inputs that defines the Env.
func (pairs EnvPairs) Inputs() []Input {
	sort.Sort(&(pairs.Pairs))

	var (
		lp     = len(pairs.Pairs)
		levels = make([]Input, lp)
		times  = make([]Input, lp-1)
	)
	for i, p := range pairs.Pairs {
		levels[i] = C(p[1])
		if i > 0 {
			times[i-1] = C(p[0] - pairs.Pairs[i-1][0])
		}
	}
	return Env{
		Levels: levels,
		Times:  times,
		Curve:  pairs.Curve,
	}.Inputs()
}

// TLC (time, level, curve) triplet
type TLC struct {
	Time, Level float32
	Curve       interface{}
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
		curves = make([]interface{}, lp-1)
	)
	for i, t := range tlc {
		levels[i] = C(t.Level)
		if i > 0 {
			times[i-1] = C(t.Time - tlc[i-1].Time)
			curves[i-1] = tlc[i-1].Curve
		}
	}
	return Env{
		Levels: levels,
		Times:  times,
		Curve:  curves,
	}.Inputs()
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
	return Env{
		Levels: levels,
		Times:  times,
		Curve:  cts,
	}.Inputs()
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
	return Env{
		Levels:      levels,
		Times:       times,
		Curve:       cts,
		ReleaseNode: C(3),
	}.Inputs()
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
	return Env{
		Levels:      levels,
		Times:       times,
		Curve:       cts,
		ReleaseNode: C(1),
	}.Inputs()
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
	return Env{
		Levels:      levels,
		Times:       times,
		Curve:       cts,
		ReleaseNode: C(0),
	}.Inputs()
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
