package sc

import (
	"testing"
)

func TestLatch(t *testing.T) {
	defName := "LatchTest"

	// Out.ar(0, Blip.ar(Latch.ar(WhiteNoise.ar, Impulse.ar(9)) * 400 + 500, 4, 0.2));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		latch := Latch{
			In:   WhiteNoise{}.Rate(AR),
			Trig: Impulse{Freq: C(9)}.Rate(AR),
		}.Rate(AR)

		return Out{
			Bus: C(0),
			Channels: Blip{
				Freq: latch.MulAdd(C(400), C(500)),
				Harm: C(4),
			}.Rate(AR).Mul(C(0.2)),
		}.Rate(AR)
	}))
}
