package sc

import (
	"testing"
)

func TestPan4(t *testing.T) {
	defName := "Pan4Test"

	// Out.ar(0, Pan4.ar(PinkNoise.ar, FSinOsc.kr(2), FSinOsc.kr(1.2), 0.3))
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		return Out{
			Bus: C(0),
			Channels: A(Pan4{
				In: A(PinkNoise{}),
				XPos: K(FSinOsc{
					Freq: C(2),
				}),
				YPos: K(FSinOsc{
					Freq: C(1.2),
				}),
				Level: C(0.3),
			}),
		}.Rate(AR)
	}))
}
