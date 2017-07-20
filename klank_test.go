package sc

import (
	"testing"
)

func TestKlank1(t *testing.T) {
	var defName = "KlankTest1"

	// Out.ar(0, Klank.ar(`[[800, 1071, 1353, 1723], nil, [1, 1, 1, 1]], PinkNoise.ar(0.007)));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		return Out{
			Bus: C(0),
			Channels: A(Klank{
				Spec: ArraySpec{
					{
						C(800),
						C(1071),
						C(1353),
						C(1723),
					},
					nil,
					Fill(4, C(1)),
				},
				In: PinkNoise{}.Rate(AR).Mul(C(0.007)),
			}),
		}.Rate(AR)
	}))
}
