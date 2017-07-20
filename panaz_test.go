package sc

import (
	"testing"
)

func TestPanAz(t *testing.T) {
	defName := "PanAzTest"

	// Out.ar(0, PanAz.ar(2, DC.ar(1), Line.ar(0, 1/2, 0.1)));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		return Out{
			Bus: C(0),
			Channels: A(PanAz{
				NumChans: 2,
				In:       DC{In: C(1)}.Rate(AR),
				Pos: A(Line{
					Start: C(0),
					End:   C(0.5),
					Dur:   C(0.1),
				}),
			}),
		}.Rate(AR)
	}))
}
