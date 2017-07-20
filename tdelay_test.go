package sc

import (
	"testing"
)

func TestTDelay(t *testing.T) {
	const defName = "TDelayTest"

	// z = Impulse.ar(2);
	// Out.ar(0, [z * 0.1, ToggleFF.ar(TDelay.ar(z, 0.5)) * SinOsc.ar(mul: 0.1)]);
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		var (
			z    = Impulse{Freq: C(2)}.Rate(AR)
			sine = SinOsc{}.Rate(AR).Mul(C(0.1))
		)
		toggle := A(ToggleFF{
			Trig: A(TDelay{
				In:  z,
				Dur: C(0.5),
			}),
		})
		return Out{
			Bus: C(0),
			Channels: Multi(
				z.Mul(C(0.1)),
				toggle.Mul(sine),
			),
		}.Rate(AR)
	}))
}
