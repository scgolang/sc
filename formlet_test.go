package sc

import (
	"testing"
)

func TestFormlet(t *testing.T) {
	defName := "FormletTest"
	def := NewSynthdef(defName, func(p Params) Ugen {
		// var in = Blip.ar(SinOsc.kr(5, 0, 20, 300), 1000, 0.1);
		// Out.ar(0, Formlet.ar(in, XLine.kr(1500, 700, 8), 0.005, 0.4));
		bus, sine := C(0), SinOsc{Freq: C(5)}.Rate(KR).MulAdd(C(20), C(300))
		blip := Blip{Freq: sine, Harm: C(1000)}.Rate(AR).Mul(C(0.1))
		line := XLine{Start: C(1500), End: C(700), Dur: C(8)}.Rate(KR)
		sig := Formlet{
			In:         blip,
			Freq:       line,
			AttackTime: C(0.005),
			DecayTime:  C(0.4),
		}.Rate(AR)
		return Out{bus, sig}.Rate(AR)
	})
	compareAndWrite(t, defName, def)
}
