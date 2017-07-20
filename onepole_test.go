package sc

import (
	"testing"
)

func TestOnePole(t *testing.T) {
	const defName = "OnePoleTest"

	// Out.ar(0, OnePole.ar(WhiteNoise.ar(0.5), Line.kr(-0.99, 0.99, 10)));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		return Out{
			Bus: C(0),
			Channels: A(OnePole{
				In: A(WhiteNoise{}).Mul(C(0.5)),
				Coeff: K(Line{
					Start: C(-0.99),
					End:   C(0.99),
					Dur:   C(10),
				}),
			}),
		}.Rate(AR)
	}))
}
