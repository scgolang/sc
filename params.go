package sc

import . "github.com/briansorahan/sc/types"

type params struct {
	l []Param
}

func (self *params) Add(name string) Param {
	p := newParam(name)
	self.l = append(self.l, p)
	return p
}

func (self *params) List() []Param {
	return self.l
}

func newParams() Params {
	p := params{make([]Param, 0)}
	return &p
}

type param struct {
	name string
	index int32
	defaultValue float32
}

func (self *param) Name() string {
	return self.name
}

func (self *param) Index() int32 {
	return self.index
}

// HACK
func (self *param) IsConstant() bool {
	return false
}

func (self *param) GetDefault() float32 {
	return self.defaultValue
}

func (self *param) SetDefault(val float32) Param {
	self.defaultValue = val
	return self
}

func newParam(name string) Param {
	p := param{name, 0, 0}
	return &p
}

