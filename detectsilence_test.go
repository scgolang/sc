package sc

import "testing"

func TestDetectSilence(t *testing.T) {
	// TODO: need a way to compare synthdef blobs where the order of the constants doesn't matter
	t.SkipNow()

	const name = "DetectSilence"

	def := NewSynthdef(name, func(params Params) Ugen {
		out := params.Add("out", 0)

		noise := LFDNoise{
			Interpolation: InterpolationCubic,
			Freq:          C(8),
		}.Rate(KR).Max(C(0))

		sine := SinOsc{
			Freq:  Rand{Lo: C(400), Hi: C(700)}.Rate(AR),
			Phase: C(0),
		}.Rate(AR).Mul(noise).SoftClip().Mul(C(0.3))

		return Out{
			Bus: out,
			Channels: DetectSilence{
				In:   sine,
				Done: FreeEnclosing,
			}.Rate(AR),
		}.Rate(AR)
	})
	compareAndWrite(t, name, def)
}
