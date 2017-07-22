package sc

import (
	"testing"
)

func TestReciprocal(t *testing.T) {
	const defName = "reciprocalExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFNoise{
			Interpolation: NoiseLinear,
			Freq:          C(1500),
		})
		return Out{
			Bus:      C(0),
			Channels: noise.Reciprocal(),
		}.Rate(AR)
	}))
}
