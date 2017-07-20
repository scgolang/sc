package sc

import (
	"testing"
)

func TestTrig1(t *testing.T) {
	const defname = "Trig1Test"

	// Trig1.ar(Dust.ar(1), 0.2) * FSinOsc.ar(800, 0.5)
	compareAndWriteStructure(t, defname, NewSynthdef(defname, func(p Params) Ugen {
		dust := A(Dust{
			Density: C(1),
		})
		trig := A(Trig1{
			In:  dust,
			Dur: C(0.2),
		})
		sine := A(FSinOsc{
			Freq:  C(800),
			Phase: C(0.5),
		})
		return Out{
			Bus:      C(0),
			Channels: trig.Mul(sine),
		}.Rate(AR)
	}))
}
