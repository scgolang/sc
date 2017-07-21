package sc

import (
	"testing"
)

func TestLinXFade2(t *testing.T) {
	defName := "LinXFade2Test"

	// Out.ar(0, LinXFade2.ar( Saw.ar, SinOsc.ar , LFTri.kr(0.1) ));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		return Out{
			Bus: C(0),
			Channels: A(LinXFade2{
				A: Saw{}.Rate(AR),
				B: SinOsc{}.Rate(AR),
				Pan: K(LFTri{
					Freq: C(0.1),
				}),
			}),
		}.Rate(AR)
	}))
}
