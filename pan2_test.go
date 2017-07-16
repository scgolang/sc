package sc

import (
	"testing"
)

func TestPan2(t *testing.T) {
	defName := "Pan2Test"

	// Out.ar(0, Pan2.ar(PinkNoise.ar(0.4), FSinOsc.kr(2), 0.3));
	def := NewSynthdef(defName, func(p Params) Ugen {
		return Out{
			Bus: C(0),
			Channels: Pan2{
				In:    PinkNoise{}.Rate(AR).Mul(C(0.4)),
				Pos:   FSinOsc{Freq: C(2)}.Rate(KR),
				Level: C(0.3),
			}.Rate(AR),
		}.Rate(AR)
	})
	compareAndWriteStructure(t, defName, def)
}
