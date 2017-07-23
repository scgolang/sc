package sc

// Params is an interface that allows you to add parameters to a synthdef.
type Params interface {
	// Add adds a named parameter to a synthdef, with an initial value
	Add(name string, initialValue float32) Input
	// List returns a list of the params that have been added to a synthdef
	List() []Param
	// Control returns a Ugen that should be used as the first ugen
	// of any synthdef that has parameters
	Control() *Ugen
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

// Control param implementation
func (ps *params) Control() *Ugen {
	return NewControl(len(ps.l))
}

// List param implementation
func (ps *params) List() []Param {
	return ps.l
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

func (p *param) Abs() Input {
	return unaryOpAbs(KR, p, 1)
}

func (p *param) Add(in Input) Input {
	return binOpAdd(KR, p, in, 1)
}

func (p *param) Ceil() Input {
	return unaryOpCeil(KR, p, 1)
}

func (p *param) Cpsmidi() Input {
	return unaryOpCpsmidi(KR, p, 1)
}

func (p *param) Cubed() Input {
	return unaryOpCubed(KR, p, 1)
}

func (p *param) Exp() Input {
	return unaryOpExp(KR, p, 1)
}

func (p *param) Frac() Input {
	return unaryOpFrac(KR, p, 1)
}

func (p *param) Floor() Input {
	return unaryOpFloor(KR, p, 1)
}

func (p *param) Index() int32 {
	return p.index
}

func (p *param) InitialValue() float32 {
	return p.val
}

func (p *param) Max(other Input) Input {
	return binOpMax(KR, p, other, 1)
}

func (p *param) Midicps() Input {
	return unaryOpMidicps(KR, p, 1)
}

func (p *param) Midiratio() Input {
	return unaryOpMidiratio(KR, p, 1)
}

func (p *param) Mul(in Input) Input {
	return binOpMul(KR, p, in, 1)
}

func (p *param) MulAdd(mul, add Input) Input {
	return mulAdd(KR, p, mul, add, 1)
}

func (p *param) Name() string {
	return p.name
}

func (p *param) Neg() Input {
	return unaryOpNeg(KR, p, 1)
}

func (p *param) Ratiomidi() Input {
	return unaryOpRatiomidi(KR, p, 1)
}

func (p *param) Reciprocal() Input {
	return unaryOpReciprocal(KR, p, 1) // TODO
}

func (p *param) Sign() Input {
	return unaryOpSign(KR, p, 1)
}

func (p *param) SoftClip() Input {
	return unaryOpSoftClip(KR, p, 1)
}

func (p *param) Sqrt() Input {
	return unaryOpSqrt(KR, p, 1)
}

func (p *param) Squared() Input {
	return unaryOpSquared(KR, p, 1)
}

func newParam(name string, index int32, initialValue float32) *param {
	p := param{name, index, initialValue}
	return &p
}
