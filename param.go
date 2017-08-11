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

func (p *param) Absdif(val Input) Input {
	return binOpAbsdif(KR, p, val, 1)
}

func (p *param) Acos() Input {
	return unaryOpAcos(KR, p, 1)
}

func (p *param) Add(in Input) Input {
	return binOpAdd(KR, p, in, 1)
}

func (p *param) Amclip(val Input) Input {
	return binOpAmclip(KR, p, val, 1)
}

func (p *param) AmpDb() Input {
	return unaryOpAmpDb(KR, p, 1)
}

func (p *param) Asin() Input {
	return unaryOpAsin(KR, p, 1)
}

func (p *param) Atan() Input {
	return unaryOpAtan(KR, p, 1)
}

func (p *param) Atan2(val Input) Input {
	return binOpAtan2(KR, p, val, 1)
}

func (p *param) Bilinrand() Input {
	return unaryOpBilinrand(KR, p, 1)
}

func (p *param) Ceil() Input {
	return unaryOpCeil(KR, p, 1)
}

func (p *param) Clip2(val Input) Input {
	return binOpClip2(KR, p, val, 1)
}

func (p *param) Coin() Input {
	return unaryOpCoin(KR, p, 1)
}

func (p *param) Cos() Input {
	return unaryOpCos(KR, p, 1)
}

func (p *param) Cosh() Input {
	return unaryOpCosh(KR, p, 1)
}

func (p *param) Cpsmidi() Input {
	return unaryOpCpsmidi(KR, p, 1)
}

func (p *param) Cpsoct() Input {
	return unaryOpCpsoct(KR, p, 1)
}

func (p *param) Cubed() Input {
	return unaryOpCubed(KR, p, 1)
}

func (p *param) DbAmp() Input {
	return unaryOpDbAmp(KR, p, 1)
}

func (p *param) Difsqr(val Input) Input {
	return binOpDifsqr(KR, p, val, 1)
}

func (p *param) Distort() Input {
	return unaryOpDistort(KR, p, 1)
}

func (p *param) Div(val Input) Input {
	return binOpDiv(KR, p, val, 1)
}

func (p *param) Excess(val Input) Input {
	return binOpExcess(KR, p, val, 1)
}

func (p *param) Exp() Input {
	return unaryOpExp(KR, p, 1)
}

func (p *param) Expon(val Input) Input {
	return binOpExpon(KR, p, val, 1)
}

func (p *param) Floor() Input {
	return unaryOpFloor(KR, p, 1)
}

func (p *param) Fold2(val Input) Input {
	return binOpFold2(KR, p, val, 1)
}

func (p *param) Frac() Input {
	return unaryOpFrac(KR, p, 1)
}

func (p *param) GCD(val Input) Input {
	return binOpGCD(KR, p, val, 1)
}

func (p *param) GT(val Input) Input {
	return binOpGT(KR, p, val, 1)
}

func (p *param) GTE(val Input) Input {
	return binOpGTE(KR, p, val, 1)
}

func (p *param) Hypot(val Input) Input {
	return binOpHypot(KR, p, val, 1)
}

func (p *param) HypotApx(val Input) Input {
	return binOpHypotApx(KR, p, val, 1)
}

func (p *param) Index() int32 {
	return p.index
}

func (p *param) InitialValue() float32 {
	return p.val
}

func (p *param) LCM(val Input) Input {
	return binOpLCM(KR, p, val, 1)
}

func (p *param) LT(val Input) Input {
	return binOpLT(KR, p, val, 1)
}

func (p *param) LTE(val Input) Input {
	return binOpLTE(KR, p, val, 1)
}

func (p *param) Linrand() Input {
	return unaryOpLinrand(KR, p, 1)
}

func (p *param) Log() Input {
	return unaryOpLog(KR, p, 1)
}

func (p *param) Log10() Input {
	return unaryOpLog10(KR, p, 1)
}

func (p *param) Log2() Input {
	return unaryOpLog2(KR, p, 1)
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

func (p *param) Min(other Input) Input {
	return binOpMin(KR, p, other, 1)
}

func (p *param) Moddif(y, mod Input) Input {
	return moddif(KR, p, y, mod, 1)
}

func (p *param) Modulo(val Input) Input {
	return binOpModulo(KR, p, val, 1)
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

func (p *param) Octcps() Input {
	return unaryOpOctcps(KR, p, 1)
}

func (p *param) Pow(val Input) Input {
	return binOpPow(KR, p, val, 1)
}

func (p *param) Rand() Input {
	return unaryOpRand(KR, p, 1)
}

func (p *param) Rand2() Input {
	return unaryOpRand2(KR, p, 1)
}

func (p *param) Ratiomidi() Input {
	return unaryOpRatiomidi(KR, p, 1)
}

func (p *param) Reciprocal() Input {
	return unaryOpReciprocal(KR, p, 1)
}

func (p *param) Ring1(val Input) Input {
	return binOpRing1(KR, p, val, 1)
}

func (p *param) Ring2(val Input) Input {
	return binOpRing2(KR, p, val, 1)
}

func (p *param) Ring3(val Input) Input {
	return binOpRing3(KR, p, val, 1)
}

func (p *param) Ring4(val Input) Input {
	return binOpRing4(KR, p, val, 1)
}

func (p *param) Round(val Input) Input {
	return binOpRound(KR, p, val, 1)
}

func (p *param) Scaleneg(val Input) Input {
	return binOpScaleneg(KR, p, val, 1)
}

func (p *param) Sign() Input {
	return unaryOpSign(KR, p, 1)
}

func (p *param) Sin() Input {
	return unaryOpSin(KR, p, 1)
}

func (p *param) Sinh() Input {
	return unaryOpSinh(KR, p, 1)
}

func (p *param) SoftClip() Input {
	return unaryOpSoftClip(KR, p, 1)
}

func (p *param) Sqrdif(val Input) Input {
	return binOpSqrdif(KR, p, val, 1)
}

func (p *param) Sqrsum(val Input) Input {
	return binOpSqrsum(KR, p, val, 1)
}

func (p *param) Sqrt() Input {
	return unaryOpSqrt(KR, p, 1)
}

func (p *param) Squared() Input {
	return unaryOpSquared(KR, p, 1)
}

func (p *param) Sum3rand() Input {
	return unaryOpSum3rand(KR, p, 1)
}

func (p *param) Sumsqr(val Input) Input {
	return binOpSumsqr(KR, p, val, 1)
}

func (p *param) Tan() Input {
	return unaryOpTan(KR, p, 1)
}

func (p *param) Tanh() Input {
	return unaryOpTanh(KR, p, 1)
}

func (p *param) Thresh(val Input) Input {
	return binOpThresh(KR, p, val, 1)
}

func (p *param) Trunc(val Input) Input {
	return binOpTrunc(KR, p, val, 1)
}

func (p *param) Wrap2(val Input) Input {
	return binOpWrap2(KR, p, val, 1)
}

func newParam(name string, index int32, initialValue float32) *param {
	p := param{name, index, initialValue}
	return &p
}
