package sc

import (
	. "github.com/scgolang/sc/types"
	. "github.com/scgolang/sc/ugens"
	"testing"
)

func TestDelayC(t *testing.T) {
	defName := "DelayCTest"
	def := NewSynthdef(defName, func(p Params) Ugen {
		// var z = Decay.ar(Dust.ar(1, 0.5), 0.3, WhiteNoise.ar());
		// Out.ar(0, DelayC.ar(z, 0.2, 0.2, 1, z));
		bus, dust := C(0), Dust{Density: C(1)}.Rate(AR).Mul(C(0.5))
		noise := WhiteNoise{}.Rate(AR)
		// BinaryOpUGen's are
		// 1) dust  (mul)
		// 2) decay (mul)
		// 3) sig   (add)
		decay := Decay{In: dust, Decay: C(0.3)}.Rate(AR).Mul(noise)
		sig := DelayC{
			In:           decay,
			MaxDelayTime: C(0.2),
			DelayTime:    C(0.2),
		}.Rate(AR).Add(decay)
		return Out{bus, sig}.Rate(AR)
	})
	compareAndWrite(t, defName, def)
}
