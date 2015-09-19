package sc

import (
	"testing"
)

func TestDust(t *testing.T) {
	defName := "DustTest"
	def := NewSynthdef(defName, func(p Params) Ugen {
		// Out.ar(0, Dust.ar(XLine.kr(20000, 2, 10), 0.5));
		bus, line := C(0), XLine{Start: C(20000), End: C(2), Dur: C(10)}.Rate(KR)
		gain, sig := C(0.5), Dust{Density: line}.Rate(AR)
		return Out{bus, sig.Mul(gain)}.Rate(AR)
	})
	compareAndWrite(t, defName, def)
}
