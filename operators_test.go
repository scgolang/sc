package sc

import (
	"testing"
)

func TestAbs(t *testing.T) {
	const defName = "absExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Abs(),
		}.Rate(AR)
	}))
}

func TestAbsdif(t *testing.T) {
	const defName = "absdifExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Absdif(A(SinOsc{})),
		}.Rate(AR)
	}))
}

func TestAcos(t *testing.T) {
	const defName = "acosExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Acos(),
		}.Rate(AR)
	}))
}

func TestAmclip(t *testing.T) {
	const defName = "amclipExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Amclip(C(0.5)),
		}.Rate(AR)
	}))
}

func TestAsin(t *testing.T) {
	const defName = "asinExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Asin(),
		}.Rate(AR)
	}))
}

func TestAtan(t *testing.T) {
	const defName = "atanExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Atan(),
		}.Rate(AR)
	}))
}

func TestAtan2(t *testing.T) {
	const defName = "atan2Example"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Atan2(C(0.5)),
		}.Rate(AR)
	}))
}

func TestAmpDb(t *testing.T) {
	const defName = "ampdbExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.AmpDb(),
		}.Rate(AR)
	}))
}

func TestBilinrand(t *testing.T) {
	const defName = "bilinrandExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Bilinrand(),
		}.Rate(AR)
	}))
}

func TestCeil(t *testing.T) {
	const defName = "ceilExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Ceil(),
		}.Rate(AR)
	}))
}

func TestCoin(t *testing.T) {
	const defName = "coinExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Coin(),
		}.Rate(AR)
	}))
}

func TestCos(t *testing.T) {
	const defName = "cosExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Cos(),
		}.Rate(AR)
	}))
}

func TestCosh(t *testing.T) {
	const defName = "coshExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Cosh(),
		}.Rate(AR)
	}))
}

func TestCpsmidi(t *testing.T) {
	const defName = "cpsmidiExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Cpsmidi(),
		}.Rate(AR)
	}))
}

func TestCpsoct(t *testing.T) {
	const defName = "cpsoctExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Cpsoct(),
		}.Rate(AR)
	}))
}

func TestCubed(t *testing.T) {
	const defName = "cubedExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Cubed(),
		}.Rate(AR)
	}))
}

func TestDbAmp(t *testing.T) {
	const defName = "dbampExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.DbAmp(),
		}.Rate(AR)
	}))
}

func TestDifsqr(t *testing.T) {
	const defName = "difsqrExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Difsqr(A(SinOsc{})),
		}.Rate(AR)
	}))
}

func TestDistort(t *testing.T) {
	const defName = "distortExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Distort(),
		}.Rate(AR)
	}))
}

func TestDiv(t *testing.T) {
	const defName = "divExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Div(C(2)),
		}.Rate(AR)
	}))
}

func TestExp(t *testing.T) {
	const defName = "expExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Exp(),
		}.Rate(AR)
	}))
}

func TestExpon(t *testing.T) {
	const defName = "exponExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Expon(C(2)),
		}.Rate(AR)
	}))
}

func TestFloor(t *testing.T) {
	const defName = "floorExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Floor(),
		}.Rate(AR)
	}))
}

func TestFrac(t *testing.T) {
	const defName = "fracExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Frac(),
		}.Rate(AR)
	}))
}

func TestGCD(t *testing.T) {
	const defName = "gcdExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.GCD(C(2)),
		}.Rate(AR)
	}))
}

func TestGT(t *testing.T) {
	const defName = "gtExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.GT(C(2)),
		}.Rate(AR)
	}))
}

func TestGTE(t *testing.T) {
	const defName = "gteExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.GTE(C(2)),
		}.Rate(AR)
	}))
}

func TestHypot(t *testing.T) {
	const defName = "hypotExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Hypot(C(2)),
		}.Rate(AR)
	}))
}

func TestHypotApx(t *testing.T) {
	const defName = "hypotapxExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.HypotApx(C(2)),
		}.Rate(AR)
	}))
}

func TestLCM(t *testing.T) {
	const defName = "lcmExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.LCM(C(2)),
		}.Rate(AR)
	}))
}

func TestLT(t *testing.T) {
	const defName = "ltExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.LT(C(2)),
		}.Rate(AR)
	}))
}

func TestLTE(t *testing.T) {
	const defName = "lteExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.LTE(C(2)),
		}.Rate(AR)
	}))
}

func TestLinrand(t *testing.T) {
	const defName = "linrandExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Linrand(),
		}.Rate(AR)
	}))
}

func TestLog(t *testing.T) {
	const defName = "logExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Log(),
		}.Rate(AR)
	}))
}

func TestLog10(t *testing.T) {
	const defName = "log10Example"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Log10(),
		}.Rate(AR)
	}))
}

func TestLog2(t *testing.T) {
	const defName = "log2Example"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Log2(),
		}.Rate(AR)
	}))
}

func TestMidiratio(t *testing.T) {
	const defName = "midiratioExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Midiratio(),
		}.Rate(AR)
	}))
}

func TestModdif(t *testing.T) {
	const defName = "moddifExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Moddif(A(SinOsc{}), nil),
		}.Rate(AR)
	}))
}

func TestModulo(t *testing.T) {
	const defName = "moduloExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Modulo(C(0.5)),
		}.Rate(AR)
	}))
}

func TestOctcps(t *testing.T) {
	const defName = "octcpsExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Octcps(),
		}.Rate(AR)
	}))
}

func TestPow(t *testing.T) {
	const defName = "powExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Pow(C(2)),
		}.Rate(AR)
	}))
}

func TestRand(t *testing.T) {
	const defName = "randExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Rand(),
		}.Rate(AR)
	}))
}

func TestRand2(t *testing.T) {
	const defName = "rand2Example"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Rand2(),
		}.Rate(AR)
	}))
}

func TestRatiomidi(t *testing.T) {
	const defName = "ratiomidiExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Ratiomidi(),
		}.Rate(AR)
	}))
}

func TestReciprocal(t *testing.T) {
	const defName = "reciprocalExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Reciprocal(),
		}.Rate(AR)
	}))
}

func TestRing1234(t *testing.T) {
	defName := "ring1Example"
	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Ring1(A(SinOsc{})),
		}.Rate(AR)
	}))

	defName = "ring2Example"
	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Ring2(A(SinOsc{})),
		}.Rate(AR)
	}))

	defName = "ring3Example"
	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Ring3(A(SinOsc{})),
		}.Rate(AR)
	}))

	defName = "ring4Example"
	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Ring4(A(SinOsc{})),
		}.Rate(AR)
	}))
}

func TestRound(t *testing.T) {
	const defName = "roundExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Round(C(0.5)),
		}.Rate(AR)
	}))
}

func TestScaleneg(t *testing.T) {
	const defName = "scalenegExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Scaleneg(C(0.5)),
		}.Rate(AR)
	}))
}

func TestSign(t *testing.T) {
	const defName = "signExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Sign(),
		}.Rate(AR)
	}))
}

func TestSin(t *testing.T) {
	const defName = "sinExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Sin(),
		}.Rate(AR)
	}))
}

func TestSinh(t *testing.T) {
	const defName = "sinhExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Sinh(),
		}.Rate(AR)
	}))
}

func TestSquared(t *testing.T) {
	const defName = "squaredExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Squared(),
		}.Rate(AR)
	}))
}

func TestSqrdif(t *testing.T) {
	const defName = "sqrdifExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Sqrdif(A(SinOsc{})),
		}.Rate(AR)
	}))
}

func TestSqrsum(t *testing.T) {
	const defName = "sqrsumExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Sqrsum(A(SinOsc{})),
		}.Rate(AR)
	}))
}

func TestSqrt(t *testing.T) {
	const defName = "sqrtExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Sqrt(),
		}.Rate(AR)
	}))
}

func TestSum3rand(t *testing.T) {
	const defName = "sum3randExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Sum3rand(),
		}.Rate(AR)
	}))
}

func TestSumsqr(t *testing.T) {
	const defName = "sumsqrExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Sumsqr(A(SinOsc{})),
		}.Rate(AR)
	}))
}

func TestTan(t *testing.T) {
	const defName = "tanExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Tan(),
		}.Rate(AR)
	}))
}

func TestTanh(t *testing.T) {
	const defName = "tanhExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Tanh(),
		}.Rate(AR)
	}))
}

func TestThresh(t *testing.T) {
	const defName = "threshExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Thresh(C(0.5)),
		}.Rate(AR)
	}))
}

func TestTrunc(t *testing.T) {
	const defName = "truncExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Trunc(C(0.5)),
		}.Rate(AR)
	}))
}
