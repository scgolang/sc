package ugens

import . "github.com/briansorahan/sc/types"

type EnvGen struct {
	Env        Envelope
	Gate       Input
	LevelScale Input
	LevelBias  Input
	TimeScale  Input
	Done       int
}

func (self *EnvGen) defaults() {
	if self.Gate == nil {
		self.Gate = C(1)
	}
	if self.LevelScale == nil {
		self.LevelScale = C(1)
	}
	if self.LevelBias == nil {
		self.LevelBias = C(0)
	}
	if self.TimeScale == nil {
		self.TimeScale = C(1)
	}
}

// Rate ugen implementation
func (self EnvGen) Rate(rate int8) Input {
	checkRate(rate)
	(&self).defaults()
	ins := []Input{self.Gate, self.LevelScale, self.LevelBias}
	ins = append(ins, self.TimeScale, C(float32(self.Done)))
	ins = append(ins, self.Env.Inputs()...)
	return UgenInput("EnvGen", rate, 0, ins...)
}
