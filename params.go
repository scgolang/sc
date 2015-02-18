package sc

import . "github.com/briansorahan/sc/types"

type params struct {
	params []Param
}

func (self *params) Add(name string, defaultValue ...float32) {
	var p param
	if len(defaultValue) == 0 {
		p = param{name, float32(0)}
	} else {
		p = param{name, float32(defaultValue[0])}
	}
	self.params = append(self.params, &p)
}

func (self *params) Get() []Param {
	return self.params
}

func newParams() *params {
	return &params{make([]Param, 0)}
}

type param struct {
	name string
	defaultValue float32
}

func (self *param) Name() string {
	return self.name
}

func (self *param) DefaultValue() float32 {
	return self.defaultValue
}
