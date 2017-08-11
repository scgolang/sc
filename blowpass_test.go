package sc

import (
	"testing"
)

func TestBLowPass(t *testing.T) {
	defName := "BLowPassTest"

	// Out.ar(0, BLowPass.ar(Blip.ar(400, 4), 300, 0.5));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		return Out{
			Bus: C(0),
			Channels: BLowPass{
				In: A(Blip{
					Freq: C(400),
					Harm: C(4),
				}),
				Freq: C(300),
				RQ:   C(0.5),
			}.Rate(AR),
		}.Rate(AR)
	}))
}
