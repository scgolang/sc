package ugens

// 0, 3, -99, -99, -- starting level, num segments, releaseNode, loopNode
// 1, 0.1, 5, 4, -- first segment: level, time, curve type, curvature
// 0.5, 1, 5, -4, -- second segment: level, time, curve type, curvature
// 0, 0.2, 5, 4 -- and so on

import (
	"fmt"
	. "github.com/briansorahan/sc/types"
	"sort"
)

const (
	CurveStep    = C(0)
	CurveLinear  = C(1)
	CurveExp     = C(2)
	CurveSine    = C(3)
	CurveWelch   = C(4)
	CurveCustom  = C(5)
	CurveSquared = C(6)
	CurveCubed   = C(7)
)

type Env struct {
	Levels      []Input
	Times       []Input
	CurveTypes  []Input
	Curvature   Input
	ReleaseNode Input
	LoopNode    Input
}

func (self *Env) defaults() {
	if self.Levels == nil {
		self.Levels = []Input{C(0), C(1), C(0)}
	}
	if self.Times == nil {
		self.Times = []Input{C(1), C(1)}
	}
	if self.CurveTypes == nil {
		self.CurveTypes = []Input{CurveLinear, CurveLinear}
	}
}

func (self Env) Inputs() []Input {
	lc, lt := len(self.CurveTypes), len(self.Times)
	if lc != lt {
		panic(fmt.Errorf("%d curve types != %d times", lc, lt))
	}
	(&self).defaults()
	numSegments := len(self.Levels) - 1
	arr := make([]Input, 4*(numSegments+1))
	arr[0] = self.Levels[0]
	arr[1] = C(numSegments)
	arr[2] = self.ReleaseNode
	arr[3] = self.LoopNode
	for i, t := range self.Times {
		arr[(4*i)+4] = self.Levels[i+1]
		arr[(4*i)+5] = t
		arr[(4*i)+6] = self.CurveTypes[i]
		arr[(4*i)+7] = self.Curvature
	}
	return arr
}

// EnvLinen creates a new envelope which has a trapezoidal shape
type EnvLinen struct {
	Attack, Sustain, Release, Level, CurveType Input
}

func (self *EnvLinen) defaults() {
	if self.Attack == nil {
		self.Attack = C(0.01)
	}
	if self.Sustain == nil {
		self.Sustain = C(1)
	}
	if self.Release == nil {
		self.Release = C(1)
	}
	if self.Level == nil {
		self.Level = C(1)
	}
	if self.CurveType == nil {
		self.CurveType = C(1)
	}
}

func (self EnvLinen) Inputs() []Input {
	(&self).defaults()
	levels := []Input{C(0), self.Level, self.Level, C(0)}
	times := []Input{self.Attack, self.Sustain, self.Release}
	ct := self.CurveType
	cts := []Input{ct, ct, ct}
	e := Env{levels, times, cts, C(0), C(-99), C(-99)}
	return e.Inputs()
}

// EnvTriangle creates a new envelope that has a triangle shape
type EnvTriangle struct {
	Dur, Level Input
}

func (self *EnvTriangle) defaults() {
	if self.Dur == nil {
		self.Dur = C(1)
	}
	if self.Level == nil {
		self.Level = C(1)
	}
}

func (self EnvTriangle) Inputs() []Input {
	(&self).defaults()
	levels := []Input{C(0), self.Level, C(0)}
	d := self.Dur.Mul(C(0.5))
	times := []Input{d, d}
	cts := []Input{CurveLinear, CurveLinear}
	e := Env{levels, times, cts, C(0), C(-99), C(-99)}
	return e.Inputs()
}

// EnvSine creates a new envelope which has a hanning window shape
type EnvSine struct {
	Dur, Level Input
}

func (self *EnvSine) defaults() {
	if self.Dur == nil {
		self.Dur = C(1)
	}
	if self.Level == nil {
		self.Level = C(1)
	}
}

func (self EnvSine) Inputs() []Input {
	(&self).defaults()
	levels := []Input{C(0), self.Level, C(0)}
	d := self.Dur.Mul(C(0.5))
	times := []Input{d, d}
	cts := []Input{CurveSine, CurveSine}
	e := Env{levels, times, cts, C(0), C(-99), C(-99)}
	return e.Inputs()
}

// EnvPerc creates a new envelope that has a percussive shape
type EnvPerc struct {
	Attack, Release, Level, Curvature Input
}

func (self *EnvPerc) defaults() {
	if self.Attack == nil {
		self.Attack = C(0.01)
	}
	if self.Release == nil {
		self.Release = C(1)
	}
	if self.Level == nil {
		self.Level = C(1)
	}
	if self.Curvature == nil {
		self.Curvature = C(-4)
	}
}

func (self EnvPerc) Inputs() []Input {
	(&self).defaults()
	levels := []Input{C(0), self.Level, C(0)}
	times := []Input{self.Attack, self.Release}
	cts := []Input{CurveCustom, CurveCustom}
	crv := self.Curvature
	e := Env{levels, times, cts, crv, C(-99), C(-99)}
	return e.Inputs()
}

// Pairs are pairs of floats: the first float is time,
// the second is level.
// They get sorted by time.
type Pairs [][2]float32

func (self Pairs) Len() int {
	return len(self)
}

func (self Pairs) Less(i, j int) bool {
	return self[i][0] < self[j][0]
}

func (self Pairs) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

// EnvPairs creates a new envelope from coordinates/pairs
type EnvPairs struct {
	Pairs     Pairs
	CurveType C
}

func (self EnvPairs) Inputs() []Input {
	sort.Sort(&(self.Pairs))
	lp := len(self.Pairs)
	levels := make([]Input, lp)
	times := make([]Input, lp-1)
	cts := make([]Input, lp-1)
	for i, p := range self.Pairs {
		levels[i] = C(p[1])
		if i > 0 {
			times[i-1] = C(p[0] - self.Pairs[i-1][0])
			cts[i-1] = self.CurveType
		}
	}
	e := Env{levels, times, cts, C(0), C(-99), C(-99)}
	return e.Inputs()
}

// TLC (time, level, curve) triplet
type TLC struct {
	Time, Level float32
	Curve C
}

// EnvTLC creates a new envelope from an array of (time, level, curve) triplets
// This is renamed from Env.xyc.
// The Curve value of the last triplet is ignored.
type EnvTLC []TLC

func (self EnvTLC) Len() int {
	return len(self)
}

func (self EnvTLC) Less(i, j int) bool {
	return self[i].Time < self[j].Time
}

func (self EnvTLC) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self EnvTLC) Inputs() []Input {
	sort.Sort(self)
	lp := len(self)
	levels := make([]Input, lp)
	times := make([]Input, lp-1)
	cts := make([]Input, lp-1)
	for i, tlc := range self {
		levels[i] = C(tlc.Level)
		if i > 0 {
			times[i-1] = C(tlc.Time - self[i-1].Time)
			cts[i-1] = self[i-1].Curve
		}
	}
	e := Env{levels, times, cts, C(0), C(-99), C(-99)}
	return e.Inputs()
}

// EnvADSR represents the ever-popular ADSR envelope
type EnvADSR struct {
	A, D, S, R, Peak, Curve, Bias Input
}

func (self *EnvADSR) defaults() {
	if self.A == nil {
		self.A = C(0.01)
	}
	if self.D == nil {
		self.D = C(0.3)
	}
	if self.S == nil {
		self.S = C(0.5)
	}
	if self.R == nil {
		self.R = C(1)
	}
	if self.Peak == nil {
		self.Peak = C(1)
	}
	if self.Curve == nil {
		self.Curve = C(-4)
	}
	if self.Bias == nil {
		self.Bias = C(0)
	}
}

func (self EnvADSR) Inputs() []Input {
	(&self).defaults()
	levels := []Input{
		C(0).Add(self.Bias),
		self.Peak.Add(self.Bias),
		self.S.Add(self.Bias),
		C(0).Add(self.Bias),
	}
	times := []Input{self.A, self.D, self.R}
	cts := []Input{CurveCustom, CurveCustom, CurveCustom}
	e := Env{levels, times, cts, self.Curve, C(2), C(-99)}
	return e.Inputs()
}

// EnvDADSR is EnvADSR with its onset delayed by D seconds
type EnvDADSR struct {
	Delay, A, D, S, R, Peak, Curve, Bias Input
}

func (self *EnvDADSR) defaults() {
	if self.Delay == nil {
		self.Delay = C(0.1)
	}
	if self.A == nil {
		self.A = C(0.01)
	}
	if self.D == nil {
		self.D = C(0.3)
	}
	if self.S == nil {
		self.S = C(0.5)
	}
	if self.R == nil {
		self.R = C(1)
	}
	if self.Peak == nil {
		self.Peak = C(1)
	}
	if self.Curve == nil {
		self.Curve = C(-4)
	}
	if self.Bias == nil {
		self.Bias = C(0)
	}
}

func (self EnvDADSR) Inputs() []Input {
	(&self).defaults()
	levels := []Input{
		C(0),
		C(0).Add(self.Bias),
		self.Peak.Add(self.Bias),
		self.S.Add(self.Bias),
		C(0).Add(self.Bias),
	}
	times := []Input{self.Delay, self.A, self.D, self.R}
	cts := []Input{CurveCustom, CurveCustom, CurveCustom, CurveCustom}
	e := Env{levels, times, cts, self.Curve, C(3), C(-99)}
	return e.Inputs()
}

// EnvASR is an attack-sustain-release envelope
type EnvASR struct {
	A, S, R, Curve Input
}

func (self *EnvASR) defaults() {
	if self.A == nil {
		self.A = C(0.01)
	}
	if self.S == nil {
		self.S = C(1)
	}
	if self.R == nil {
		self.R = C(1)
	}
	if self.Curve == nil {
		self.Curve = C(-4)
	}
}

func (self EnvASR) Inputs() []Input {
	(&self).defaults()
	levels := []Input{C(0), self.S, C(0)}
	times := []Input{self.A, self.R}
	cts := []Input{CurveCustom, CurveCustom}
	e := Env{levels, times, cts, self.Curve, C(1), C(-99)}
	return e.Inputs()
}

// EnvCutoff creates an envelope with no attack segment.
// It simply sustains at the peak level until released.
type EnvCutoff struct {
	R, Level, CurveType Input
}

func (self *EnvCutoff) defaults() {
	if self.R == nil {
		self.R = C(0.1)
	}
	if self.Level == nil {
		self.Level = C(1)
	}
	if self.CurveType == nil {
		self.CurveType = CurveLinear
	}
}

func (self EnvCutoff) Inputs() []Input {
	(&self).defaults()
	levels := []Input{self.Level, C(0)}
	times := []Input{self.R}
	cts := []Input{self.CurveType}
	e := Env{levels, times, cts, C(0), C(0), C(-99)}
	return e.Inputs()
}

// I don't understand Env.circle [bps]
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
