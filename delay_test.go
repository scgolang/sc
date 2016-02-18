package sc

import (
	"testing"
)

func TestDelayC(t *testing.T) {
	defName := "DelayCTest"
	def := NewSynthdef(defName, func(p Params) Ugen {
		bus, dust := C(0), Dust{Density: C(1)}.Rate(AR).Mul(C(0.5))
		noise := WhiteNoise{}.Rate(AR)
		decay := Decay{In: dust, Decay: C(0.3)}.Rate(AR).Mul(noise)
		sig := Delay{
			Interpolation: InterpolationCubic,
			In:            decay,
			MaxDelayTime:  C(0.2),
			DelayTime:     C(0.2),
		}.Rate(AR).Add(decay)
		return Out{bus, sig}.Rate(AR)
	})
	compareAndWrite(t, defName, def)
}

func TestDelayL(t *testing.T) {
	defName := "DelayLTest"
	def := NewSynthdef(defName, func(p Params) Ugen {
		bus, dust := C(0), Dust{Density: C(1)}.Rate(AR).Mul(C(0.5))
		noise := WhiteNoise{}.Rate(AR)
		decay := Decay{In: dust, Decay: C(0.3)}.Rate(AR).Mul(noise)
		sig := Delay{
			Interpolation: InterpolationLinear,
			In:            decay,
			MaxDelayTime:  C(0.2),
			DelayTime:     C(0.2),
		}.Rate(AR).Add(decay)
		return Out{bus, sig}.Rate(AR)
	})
	compareAndWrite(t, defName, def)
}

func TestDelayN(t *testing.T) {
	defName := "DelayNTest"
	def := NewSynthdef(defName, func(p Params) Ugen {
		bus, dust := C(0), Dust{Density: C(1)}.Rate(AR).Mul(C(0.5))
		noise := WhiteNoise{}.Rate(AR)
		decay := Decay{In: dust, Decay: C(0.3)}.Rate(AR).Mul(noise)
		sig := Delay{
			Interpolation: InterpolationNone,
			In:            decay,
			MaxDelayTime:  C(0.2),
			DelayTime:     C(0.2),
		}.Rate(AR).Add(decay)
		return Out{bus, sig}.Rate(AR)
	})
	compareAndWrite(t, defName, def)
}
