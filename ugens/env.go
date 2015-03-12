package ugens

// 0, 3, -99, -99, -- starting level, num segments, releaseNode, loopNode
// 1, 0.1, 5, 4, -- first segment: level, time, curve type, curvature
// 0.5, 1, 5, -4, -- second segment: level, time, curve type, curvature
// 0, 0.2, 5, 4 -- and so on

import . "github.com/briansorahan/sc/types"

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

// EnvLinen http://doc.sccode.org/Classes/Env.html#*linen
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
	e := Env{levels, times, self.CurveType, C(0), C(-99), C(-99)}
	return e.InputsArray()
}

// EnvTriangle http://doc.sccode.org/Classes/Env.html#*triangle
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
	e := Env{levels, times, CurveLinear, C(0), C(-99), C(-99)}
	return e.InputsArray()
}

// EnvSine http://doc.sccode.org/Classes/Env.html#*sine
type EnvSine struct {
	Dur, Level Input
}

func (self EnvSine) InputsArray() []Input {
	levels := []Input{C(0), self.Level, C(0)}
	d := self.Dur.Mul(C(0.5))
	times := []Input{d, d}
	e := Env{levels, times, CurveSine, C(0), C(-99), C(-99)}
	return e.InputsArray()
}

// EnvPerc http://doc.sccode.org/Classes/Env.html#*perc
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
	crv := self.Curvature
	e := Env{levels, times, CurveCustom, crv, C(-99), C(-99)}
	return e.InputsArray()
}

type Env struct {
	Levels                                      []Input
	Times                                       []Input
	CurveType, Curvature, ReleaseNode, LoopNode Input
}

func (self *Env) InputsArray() []Input {
	numSegments := len(self.Levels)
	arr := make([]Input, 4*numSegments)
	arr[0] = self.Levels[0]
	arr[1] = C(float32(numSegments - 1))
	arr[2] = self.ReleaseNode
	arr[3] = self.LoopNode
	for i, t := range self.Times {
		arr[(4*i)+4] = self.Levels[i+1]
		arr[(4*i)+5] = t
		arr[(4*i)+6] = self.CurveType
		arr[(4*i)+7] = self.Curvature
	}
	return arr
}
