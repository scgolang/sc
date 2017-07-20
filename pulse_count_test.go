package sc

import (
	"testing"
)

func TestPulseCount(t *testing.T) {
	const defName = "PulseCountTest"

	// Out.ar(0, SinOsc.ar(PulseCount.ar(Impulse.ar(10), Impulse.ar(0.4)) * 200, 0, 0.05));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		return Out{
			Bus: C(0),
			Channels: A(SinOsc{
				Freq: A(PulseCount{
					Trig:  Impulse{Freq: C(10)}.Rate(AR),
					Reset: Impulse{Freq: C(0.4)}.Rate(AR),
				}).Mul(C(200)),
			}).Mul(C(0.05)),
		}.Rate(AR)
	}))
}
