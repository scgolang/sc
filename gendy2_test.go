package sc

import (
	"testing"
)

func TestGendy2(t *testing.T) {
	const defName = "Gendy2Test"

	// Out.ar(0, Pan2.ar(Gendy2.ar));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		return Out{
			Bus: C(0),
			Channels: A(Pan2{
				In: A(Gendy2{}),
			}),
		}.Rate(AR)
	}))
}
