package sc

import . "github.com/briansorahan/sc/types"

type param struct {
	name string
	defaultValue float32
}

func (self *param) Name() string {
	return self.name
}

func (self *param) GetDefault() float32 {
	return self.defaultValue
}

func (self *param) SetDefault(val float32) Param {
	self.defaultValue = val
	return self
}

func NewParam(name string) Param {
	p := param{name, 0}
	return &p
}
