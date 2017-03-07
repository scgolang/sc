package sc

// Params is an interface that allows you to add parameters to
// a synthdef.
type Params interface {
	// Add adds a named parameter to a synthdef, with an initial value
	Add(name string, initialValue float32) Input
	// List returns a list of the params that have been added to a synthdef
	List() []Param
	// Control returns a Ugen that should be used as the first ugen
	// of any synthdef that has parameters
	Control() Ugen
}

// Param is the interface of a single synthdef parameter.
type Param interface {
	// Name returns the name of the synthdef param
	Name() string
	// Index returns the index of the synthdef param
	Index() int32
	// InitialValue returns the initial value of the synthdef param
	InitialValue() float32
}

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

func (p *param) Max(other Input) Input {
	return BinOpMax(KR, p, other, 1)
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

func (p *param) SoftClip() Input {
	return UnaryOpSoftClip(KR, p, 1)
}

func newParam(name string, index int32, initialValue float32) *param {
	p := param{name, index, initialValue}
	return &p
}
