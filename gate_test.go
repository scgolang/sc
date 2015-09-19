package sc

import (
	"testing"
)

func TestGate(t *testing.T) {
	defName := "GateTest"
	def := NewSynthdef(defName, func(p Params) Ugen {
		// var noise = WhiteNoise.kr(1, 0);
		// var pulse = LFPulse.kr(1.333, 0.5);
		// Out.ar(0, Gate.ar(noise, pulse));
		bus, noise := C(0), WhiteNoise{}.Rate(KR)
		pulse := LFPulse{Freq: C(1.333), Iphase: C(0.5)}.Rate(KR)
		sig := Gate{In: noise, Trig: pulse}.Rate(AR)
		return Out{bus, sig}.Rate(AR)
	})
	compareAndWrite(t, defName, def)
}
