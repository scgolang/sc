package sc

import (
	"testing"
)

func TestRingz(t *testing.T) {
	const defName = "RingzTest"

	// Out.ar(0, Ringz.ar(Impulse.ar(6, 0, 0.3), 2000, XLine.kr(4, 0.04, 8)))
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		pulse := A(Impulse{
			Freq: C(6),
		})
		return Out{
			Bus: C(0),
			Channels: A(Ringz{
				In:   pulse.Mul(C(0.3)),
				Freq: C(2000),
				DecayTime: K(XLine{
					Start: C(4),
					End:   C(0.04),
					Dur:   C(8),
				}),
			}),
		}.Rate(AR)
	}))
}
