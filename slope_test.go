package sc

import (
	"testing"
)

func TestSlope(t *testing.T) {
	const defName = "SlopeTest"

	// Out.ar(0, SinOsc.ar(Slope.ar(LFNoise2.ar(10))));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		slope := A(Slope{
			In: A(LFNoise{
				Interpolation: NoiseQuadratic,
				Freq:          C(10),
			}),
		})
		return Out{
			Bus:      C(0),
			Channels: A(SinOsc{Freq: slope}),
		}.Rate(AR)
	}))
}
