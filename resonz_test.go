package sc

import (
	"testing"
)

func TestResonz(t *testing.T) {
	const defName = "ResonzTest"

	// Out.ar(0, Resonz.ar(WhiteNoise.ar(0.5), XLine.kr(1000,8000,10), 0.05));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		return Out{
			Bus: C(0),
			Channels: A(Resonz{
				In: WhiteNoise{}.Rate(AR).Mul(C(0.5)),
				Freq: K(XLine{
					Start: C(1000),
					End:   C(8000),
					Dur:   C(10),
				}),
				BWR: C(0.05),
			}),
		}.Rate(AR)
	}))
}
