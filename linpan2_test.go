package sc

import (
	"testing"
)

func TestLinPan2(t *testing.T) {
	const defName = "LinPan2Test"

	// Out.ar(0, LinPan2.ar(FSinOsc.ar(800, 0, 0.1), FSinOsc.kr(3)))
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		sig := A(FSinOsc{
			Freq:  C(800),
			Phase: C(0),
		})
		pos := K(FSinOsc{
			Freq: C(3),
		})
		return Out{
			Bus: C(0),
			Channels: A(LinPan2{
				In:  sig.Mul(C(0.1)),
				Pos: pos,
			}),
		}.Rate(AR)
	}))
}
