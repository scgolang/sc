package sc

import . "github.com/briansorahan/sc/types"

type params struct {
	l []Param
}

func (self *params) Add(name string, initialValue ...float32) Param {
	idx := len(self.l)
	p := newParam(name, int32(idx))
	if len(initialValue) > 0 {
		p.SetDefault(initialValue[0])
	}
	self.l = append(self.l, p)
	return p
}

func (self *params) List() []Param {
	return self.l
}

func (self *params) Control() UgenNode {
	return newControl(len(self.l))
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

func (self *param) GetDefault() float32 {
	return self.defaultValue
}

func (self *param) SetDefault(val float32) Param {
	self.defaultValue = val
	return self
}

func newParam(name string, index int32) Param {
	p := param{name, index, 0}
	return &p
}

type control struct {
	outputs []Output
}

func (self *control) Name() string {
	return "Control"
}

func (self *control) Rate() int8 {
	return int8(1)
}

func (self *control) SpecialIndex() int16 {
	return 0
}

func (self *control) Inputs() []interface{} {
	return make([]interface{}, 0)
}

func (self *control) Outputs() []Output {
	return self.outputs
}

func (self *control) Mul(val float32) UgenNode {
	return self
}

func (self *control) Add(val float32) UgenNode {
	return self
}

type controlOutput struct {}

func (self *controlOutput) Rate() int8 {
	return 1
}

func newControl(numOutputs int) UgenNode {
	outputs := make([]Output, numOutputs)
	o := controlOutput{}
	for i := 0; i < numOutputs; i++ {
		outputs[i] = &o
	}
	c := control{outputs}
	return &c
}
