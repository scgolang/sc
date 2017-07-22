package sc

import (
	"testing"
)

func TestLFDClipNoise(t *testing.T) {
	const defName = "LFDClipNoiseTest"

	// Out.ar(0, SinOsc.ar(
	//      LFDClipNoise.ar(4, 200, 600),
	//      0, 0.2
	// ));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		noise := A(LFDClipNoise{
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
