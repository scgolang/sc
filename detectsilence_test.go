package sc

import "testing"

func TestDetectSilence(t *testing.T) {
	t.SkipNow()

	const name = "DetectTheSilence"

	def := NewSynthdef(name, func(params Params) Ugen {
		out := params.Add("out", 0)

		sine := SinOsc{
			Freq: Rand{
				Lo: C(400),
				Hi: C(700),
			}.Rate(AR),
		}.Rate(AR)
		return Out{
			Bus: out,
			Channels: sine.Mul(LFDNoise{
				Interpolation: InterpolationCubic,
				Freq:          C(8),
			}.Rate(KR)),
		}.Rate(AR)
	})
	same, err := def.CompareToFile("fixtures/DetectSilence.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang version")
	}
}
