package sc

import (
	"testing"
)

func TestLFClipNoise(t *testing.T) {
	const defName = "LFClipNoiseTest"

	// Out.ar(0, SinOsc.ar(
	//      LFClipNoise.ar(4, 200, 600),
	//      0, 0.2
	// ));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFClipNoise{
			Freq: C(4),
		})
		sine := A(SinOsc{
			Freq: noise.MulAdd(C(200), C(600)),
		})
		return Out{
			Bus:      C(0),
			Channels: sine.Mul(C(0.2)),
		}.Rate(AR)
	}))
}
