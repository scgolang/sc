package sc

import "testing"

// Out.ar(0, AllpassC.ar(Decay.ar(Dust.ar(1,0.5), 0.2, WhiteNoise.ar), 0.2, 0.2, 3));
func TestAllpassnExample(t *testing.T) {
	const defName = "AllpassExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		var (
			noise = WhiteNoise{}.Rate(AR)
			dust  = Dust{C(1)}.Rate(AR).Mul(C(0.5))
			decay = Decay{dust, C(0.2)}.Rate(AR).Mul(noise)
			sig   = Allpass{
				Interpolation: InterpolationCubic,
				In:            decay,
				MaxDelayTime:  C(0.2),
				DelayTime:     C(0.2),
				DecayTime:     C(3),
			}.Rate(AR)
		)
		return Out{
			Bus:      C(0),
			Channels: sig,
		}.Rate(AR)
	}))
}
