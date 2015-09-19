package sc

import (
	"testing"
)

func TestDelayL(t *testing.T) {
	defName := "DelayLTest"
	def := NewSynthdef(defName, func(p Params) Ugen {
		bus, dust := C(0), Dust{Density: C(1)}.Rate(AR).Mul(C(0.5))
		noise := WhiteNoise{}.Rate(AR)
		decay := Decay{In: dust, Decay: C(0.3)}.Rate(AR).Mul(noise)
		sig := DelayL{
			In:           decay,
			MaxDelayTime: C(0.2),
			DelayTime:    C(0.2),
		}.Rate(AR).Add(decay)
		return Out{bus, sig}.Rate(AR)
	})
	compareAndWrite(t, defName, def)
}
