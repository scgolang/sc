package sc

import . "github.com/scgolang/sc/types"
import . "github.com/scgolang/sc/ugens"

// params provides a way to add parameters to a synthdef
type params struct {
	l []Param
}

// Add param implementation
func (ps *params) Add(name string, initialValue float32) Input {
	idx := len(ps.l)
	p := newParam(name, int32(idx), initialValue)
	ps.l = append(ps.l, p)
	return p
}

// List param implementation
func (ps *params) List() []Param {
	return ps.l
}

// Control param implementation
func (ps *params) Control() Ugen {
	return NewControl(len(ps.l))
}

// newParams creates a new params instance
func newParams() *params {
	return &params{make([]Param, 0)}
}

type param struct {
	name  string
	index int32
	val   float32
}

func (p *param) Name() string {
	return p.name
}

func (p *param) Index() int32 {
	return p.index
}

func (p *param) InitialValue() float32 {
	return p.val
}

func (p *param) Mul(in Input) Input {
	return BinOpMul(KR, p, in, 1)
}

func (p *param) Add(in Input) Input {
	return BinOpAdd(KR, p, in, 1)
}

func (p *param) MulAdd(mul, add Input) Input {
	return MulAdd(KR, p, mul, add, 1)
}

func newParam(name string, index int32, initialValue float32) *param {
	p := param{name, index, initialValue}
	return &p
}
