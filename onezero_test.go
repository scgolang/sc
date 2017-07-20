package sc

import (
	"testing"
)

func TestOneZero(t *testing.T) {
	const defName = "OneZeroTest"

	// Out.ar(0, OneZero.ar(WhiteNoise.ar(0.5), Line.kr(-0.5, 0.5, 10)));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		return Out{
			Bus: C(0),
			Channels: A(OneZero{
				In: A(WhiteNoise{}).Mul(C(0.5)),
				Coeff: K(Line{
					Start: C(-0.5),
					End:   C(0.5),
					Dur:   C(10),
				}),
			}),
		}.Rate(AR)
	}))
}
