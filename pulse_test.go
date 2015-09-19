package sc

import (
	"testing"
)

func TestPulse(t *testing.T) {
	defName := "PulseTest"
	def := NewSynthdef(defName, func(p Params) Ugen {
		// Out.ar(0, Pulse.ar(XLine.kr(40,4000,6),0.1, 0.2));
		bus, width, gain := C(0), C(0.1), C(0.2)
		line := XLine{Start: C(40), End: C(4000), Dur: C(6)}.Rate(KR)
		sig := Pulse{Freq: line, Width: width}.Rate(AR).Mul(gain)
		return Out{bus, sig}.Rate(AR)
	})
	compareAndWrite(t, defName, def)
}
