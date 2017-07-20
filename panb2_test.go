package sc

import (
	"testing"
)

func TestPanB2(t *testing.T) {
	defName := "PanB2Test"

	// var w, x, y, p, a, b, c, d;
	// p = PinkNoise.ar; // source
	// // B-format encode
	// #w, x, y = PanB2.ar(p, MouseX.kr(-1,1), 0.1);
	// // B-format decode to quad
	// #a, b, c, d = DecodeB2.ar(4, w, x, y);
	// Out.ar(0, [a, b, d, c]); // reorder to my speaker arrangement: Lf Rf Lr Rr
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		d := A(DecodeB2{
			NumChans: 4,
			In: A(PanB2{
				In: A(PinkNoise{}),
				Azimuth: A(MouseX{
					Min: C(-1),
					Max: C(1),
				}),
				Gain: C(0.1),
			}),
		})
		return Out{
			Bus:      C(0),
			Channels: d,
		}.Rate(AR)
	}))
}
