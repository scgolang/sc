package sc

import (
	"testing"
)

func TestGendy1(t *testing.T) {
	const defName = "Gendy1Test"

	// Out.ar(0, Pan2.ar(Gendy1.ar));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		return Out{
			Bus: C(0),
			Channels: A(Pan2{
				In: A(Gendy1{}),
			}),
		}.Rate(AR)
	}))
}
