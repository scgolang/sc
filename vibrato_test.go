package sc

import (
	"testing"
)

func TestVibrato(t *testing.T) {
	const defName = "VibratoTest"

	// var freq    = DC.ar(400.0);
	// var rate    = MouseX.kr(2.0, 100.0);
	// var ratevar = MouseY.kr(0.0, 1.0);
	// Out.ar(0, SinOsc.ar(Vibrato.ar(freq, rate, 0.1, 1.0, 1.0, ratevar, 0.1)));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		var (
			freq    = DC{In: C(400)}.Rate(AR)
			rate    = MouseX{Min: C(2), Max: C(100)}.Rate(KR)
			ratevar = MouseY{Min: C(0), Max: C(1)}.Rate(KR)
		)
		return Out{
			Bus: C(0),
			Channels: A(SinOsc{
				Freq: A(Vibrato{
					Freq:           freq,
					Speed:          rate,
					Depth:          C(0.1),
					Delay:          C(1),
					Onset:          C(1),
					RateVariation:  ratevar,
					DepthVariation: C(0.1),
				}),
			}),
		}.Rate(AR)
	}))
}
