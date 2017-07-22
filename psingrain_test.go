package sc

import (
	"testing"
)

func TestPSinGrain(t *testing.T) {
	const defName = "PSinGrainTest"

	// Out.ar(0, PSinGrain.ar(880, 0.1, 0.7();
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		return Out{
			Bus: C(0),
			Channels: A(PSinGrain{
				Freq: C(880),
				Dur:  C(0.1),
				Amp:  C(0.7),
			}),
		}.Rate(AR)
	}))
}
