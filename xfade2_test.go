package sc

import (
	"testing"
)

func TestXFade2(t *testing.T) {
	defName := "XFade2Test"

	// Out.ar(0, XFade2.ar( Saw.ar, SinOsc.ar , LFTri.kr(0.1) ));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		return Out{
			Bus: C(0),
			Channels: A(XFade2{
				A: Saw{}.Rate(AR),
				B: SinOsc{}.Rate(AR),
				Pan: K(LFTri{
					Freq: C(0.1),
				}),
			}),
		}.Rate(AR)
	}))
}
