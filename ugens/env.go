package ugens

// 0, 3, -99, -99, -- starting level, num segments, releaseNode, loopNode
// 1, 0.1, 5, 4, -- first segment: level, time, curve type, curvature
// 0.5, 1, 5, -4, -- second segment: level, time, curve type, curvature
// 0, 0.2, 5, 4 -- and so on

import (
	. "github.com/briansorahan/sc/types"
)

const (
	CurveStep        = 0
	CurveLinear      = 1
	CurveExponential = 2
	CurveSine        = 3
	CurveWelch       = 4
	CurveCustom      = 5
	CurveSquared     = 6
	CurveCubed       = 7
)

// Env is not a ugen, but rather a way to generate
// Control arrays that get handed to EnvGen
var Env = newEnv()

type Envelope interface {
	// InputsArray provides EnvGen with the data it needs
	// to get a list of inputs
	InputsArray() []Input
}

type envelopeImpl struct {
	levels      []Input
	times       []Input
	curvetype   int
	curveature  Input
	releaseNode Input
	loopNode    Input
}

func (self *envelopeImpl) InputsArray() []Input {
	numSegments := len(self.levels)
	arr := make([]Input, 4*numSegments)
	arr[0] = self.levels[0]
	arr[1] = C(float32(numSegments - 1))
	arr[2] = self.releaseNode
	arr[3] = self.loopNode
	for i, t := range self.times {
		arr[(4*i)+4] = self.levels[i+1]
		arr[(4*i)+5] = t
		arr[(4*i)+6] = C(float32(self.curvetype))
		arr[(4*i)+7] = self.curveature
	}
	return arr
}

type env struct {
}

// Linen http://doc.sccode.org/Classes/Env.html#*linen
func (self *env) Linen(at, st, rt, level Input, curvetype int) Envelope {
	levels := []Input{C(0), level, level, C(0)}
	times := []Input{at, st, rt}
	e := envelopeImpl{levels, times, curvetype, C(0), C(-99), C(-99)}
	return &e
}

// Triangle http://doc.sccode.org/Classes/Env.html#*triangle
func (self *env) Triangle(dur, level Input) Envelope {
	levels := []Input{C(0), level, C(0)}
	d := dur.Mul(C(0.5))
	times := []Input{d, d}
	e := envelopeImpl{levels, times, CurveLinear, C(0), C(-99), C(-99)}
	return &e
}

// Triangle http://doc.sccode.org/Classes/Env.html#*triangle
func (self *env) Sine(dur, level Input) Envelope {
	levels := []Input{C(0), level, C(0)}
	d := dur.Mul(C(0.5))
	times := []Input{d, d}
	e := envelopeImpl{levels, times, CurveSine, C(0), C(-99), C(-99)}
	return &e
}

// Perc http://doc.sccode.org/Classes/Env.html#*perc
func (self *env) Perc(at, rt, level, curveature Input) Envelope {
	levels := []Input{C(0), level, C(0)}
	times := []Input{at, rt}
	e := envelopeImpl{levels, times, CurveCustom, curveature, C(-99), C(-99)}
	return &e
}

func newEnv() *env {
	eg := env{}
	return &eg
}
