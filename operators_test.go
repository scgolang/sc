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
