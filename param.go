package sc

import . "github.com/scgolang/sc/types"
import . "github.com/scgolang/sc/ugens"

// params provides a way to add parameters to a synthdef
type params struct {
	l []Param
}

// Add param implementation
func (self *params) Add(name string, initialValue float32) Input {
	idx := len(self.l)
	p := newParam(name, int32(idx), initialValue)
	self.l = append(self.l, p)
	return p
}

// List param implementation
func (self *params) List() []Param {
	return self.l
}

// Control param implementation
func (self *params) Control() Ugen {
	return newControl(len(self.l))
}

// newParams creates a new params instance
func newParams() *params {
	p := params{make([]Param, 0)}
	return &p
}

type param struct {
	name  string
	index int32
	val   float32
}

func (self *param) Name() string {
	return self.name
}

func (self *param) Index() int32 {
	return self.index
}

func (self *param) InitialValue() float32 {
	return self.val
}

func (self *param) Mul(in Input) Input {
	return BinOpMul(KR, self, in)
}

func (self *param) Add(in Input) Input {
	return BinOpAdd(KR, self, in)
}

func (self *param) MulAdd(mul, add Input) Input {
	return MulAdd(KR, self, mul, add)
}

func newParam(name string, index int32, initialValue float32) *param {
	p := param{name, index, initialValue}
	return &p
}

type Control struct {
	inputs  []Input
	outputs []Output
}

func (self *Control) Name() string {
	return "Control"
}

func (self *Control) Rate() int8 {
	return int8(1)
}

func (self *Control) SpecialIndex() int16 {
	return 0
}

func (self *Control) Inputs() []Input {
	return self.inputs
}

func (self *Control) Outputs() []Output {
	return self.outputs
}

func (self *Control) Add(val Input) Input {
	return self
}

func (self *Control) Mul(val Input) Input {
	return self
}

func (self *Control) MulAdd(mul, add Input) Input {
	return self
}

type ControlOutput struct{}

func (self *ControlOutput) Rate() int8 {
	return 1
}

func newControl(numOutputs int) Ugen {
	outputs := make([]Output, numOutputs)
	o := ControlOutput{}
	for i := 0; i < numOutputs; i++ {
		outputs[i] = &o
	}
	c := Control{make([]Input, 0), outputs}
	return &c
}
