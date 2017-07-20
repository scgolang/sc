package sc

import (
	"testing"
)

func TestSlew(t *testing.T) {
	const defName = "SlewTest"

	// Out.ar(0, Slew.ar(Saw.ar(800,mul:0.2), 400, 400));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		saw := Saw{Freq: C(800)}.Rate(AR)
		return Out{
			Bus: C(0),
			Channels: A(Slew{
				In: saw.Mul(C(0.2)),
				Up: C(400),
				Dn: C(400),
			}),
		}.Rate(AR)
	}))
}
