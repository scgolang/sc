package sc

import (
	"testing"
)

func TestLFCub(t *testing.T) {
	defName := "LFCubTest"
	def := NewSynthdef(defName, func(p Params) Ugen {
		// var freq = LFCub.kr(LFCub.kr(0.2, 0, 8, 10), 0, 400, 800);
		// var sig = LFCub.ar(freq, 0, 0.1);
		// Out.ar(0, sig);
		bus, gain := C(0), C(0.1)
		lfo1 := LFCub{Freq: C(0.2)}.Rate(KR).MulAdd(C(8), C(10))
		lfo2 := LFCub{Freq: lfo1}.Rate(KR).MulAdd(C(400), C(800))
		sig := LFCub{Freq: lfo2}.Rate(AR).Mul(gain)
		return Out{bus, sig}.Rate(AR)
	})
	compareAndWrite(t, defName, def)
}
