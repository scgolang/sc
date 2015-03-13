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
	CurveStep        = C(0)
	CurveLinear      = C(1)
	CurveExponential = C(2)
	CurveSine        = C(3)
	CurveWelch       = C(4)
	CurveCustom      = C(5)
	CurveSquared     = C(6)
	CurveCubed       = C(7)
)

// EnvLinen http://doc.sccode.org/Classes/Env.html#linen
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

func (self EnvLinen) InputsArray() []Input {
	(&self).defaults()
	levels := []Input{C(0), self.Level, self.Level, C(0)}
	times := []Input{self.Attack, self.Sustain, self.Release}
	ct := self.CurveType
	cts := []Input{ct, ct, ct}
	e := Env{levels, times, cts, C(0), C(-99), C(-99)}
	return e.InputsArray()
}

// EnvTriangle http://doc.sccode.org/Classes/Env.html#triangle
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

func (self EnvTriangle) InputsArray() []Input {
	levels := []Input{C(0), self.Level, C(0)}
	d := self.Dur.Mul(C(0.5))
	times := []Input{d, d}
	cts := []Input{CurveLinear, CurveLinear}
	e := Env{levels, times, cts, C(0), C(-99), C(-99)}
	return e.InputsArray()
}

// EnvSine http://doc.sccode.org/Classes/Env.html#sine
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

func (self EnvSine) InputsArray() []Input {
	(&self).defaults()
	levels := []Input{C(0), self.Level, C(0)}
	d := self.Dur.Mul(C(0.5))
	times := []Input{d, d}
	cts := []Input{CurveSine, CurveSine}
	e := Env{levels, times, cts, C(0), C(-99), C(-99)}
	return e.InputsArray()
}

// EnvPerc http://doc.sccode.org/Classes/Env.html#perc
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

func (self EnvPerc) InputsArray() []Input {
	(&self).defaults()
	levels := []Input{C(0), self.Level, C(0)}
	times := []Input{self.Attack, self.Release}
	cts := []Input{CurveCustom, CurveCustom}
	crv := self.Curvature
	e := Env{levels, times, cts, crv, C(-99), C(-99)}
	return e.InputsArray()
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

// EnvPairs http://doc.sccode.org/Classes/Env.html#pairs
type EnvPairs struct {
	Pairs     Pairs
	CurveType float32
}

func (self EnvPairs) InputsArray() []Input {
	sort.Sort(self.Pairs)
	lp := len(self.Pairs)
	levels := make([]Input, lp)
	times := make([]Input, lp-1)
	cts := make([]Input, lp-1)
	for i, p := range self.Pairs {
		levels[i] = C(p[1])
		if i > 0 {
			times[i-1] = C(p[0] - p[0])
			cts[i-1] = C(self.CurveType)
		}
	}
	e := Env{levels, times, cts, C(0), C(-99), C(-99)}
	return e.InputsArray()
}

// TLC (time, level, curve) triplet
type TLC struct {
	T, L, C float32
}

// EnvTLC represents an array of (time, level, curve) triplets
type EnvTLC []TLC

func (self EnvTLC) Len() int {
	return len(self)
}

func (self EnvTLC) Less(i, j int) bool {
	return self[i].T < self[j].T
}

func (self EnvTLC) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self EnvTLC) InputsArray() []Input {
	sort.Sort(self)
	lp := len(self)
	levels := make([]Input, lp)
	times := make([]Input, lp-1)
	cts := make([]Input, lp-1)
	for i, tlc := range self {
		levels[i] = C(tlc.L)
		if i > 0 {
			times[i-1] = C(tlc.T - tlc.T)
			cts[i-1] = C(tlc.C)
		}
	}
	e := Env{levels, times, cts, C(0), C(-99), C(-99)}
	return e.InputsArray()
}

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

func (self Env) InputsArray() []Input {
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
