package sc

import (
	"testing"
)

func TestCascade(t *testing.T) {
	const defName = "CascadeExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		var (
			freq = Multi(C(440), C(441))
			mod1 = SinOsc{Freq: freq}.Rate(AR)
			mod2 = SinOsc{Freq: mod1}.Rate(AR)
		)
		return Out{
			Bus: C(0),
			Channels: SinOsc{
				Freq: mod2,
			}.Rate(AR),
		}.Rate(AR)
	}))
}
