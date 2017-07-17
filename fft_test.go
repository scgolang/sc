package sc

import (
	"testing"
)

func TestFFT(t *testing.T) {
	const defName = "FFTTest"

	// var in, chain;
	// in = WhiteNoise.ar(0.2);
	// chain = FFT(LocalBuf(2048), in);
	// chain = PV_BrickWall(chain, SinOsc.kr(0.1));
	// Out.ar(0, IFFT(chain));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		chain := PVBrickWall{
			Buffer: FFT{
				Buffer: LocalBuf{NumFrames: C(2048)}.Rate(AR),
				In:     WhiteNoise{}.Rate(AR).Mul(C(0.2)),
			}.Rate(AR),
			Wipe: SinOsc{Freq: C(0.1)}.Rate(KR),
		}.Rate(AR)

		return Out{
			Bus:      C(0),
			Channels: IFFT{Buffer: chain}.Rate(AR),
		}.Rate(AR)
	}))
}
