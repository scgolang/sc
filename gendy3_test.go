package sc

import (
	"testing"
)

func TestGendy3(t *testing.T) {
	const defName = "Gendy3Test"

	// Out.ar(0, Pan2.ar(Gendy3.ar));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		return Out{
			Bus: C(0),
			Channels: A(Pan2{
				In: A(Gendy3{}),
			}),
		}.Rate(AR)
	}))
}
