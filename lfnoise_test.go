package sc

import "testing"

func TestLFNoise1(t *testing.T) {
	def := NewSynthdef("LFNoise1Example", func(p Params) Ugen {
		start, end, dur, done := C(1000), C(10000), C(10), 0
		bus, gain := C(0), C(0.25)
		freq := XLine{start, end, dur, done}.Rate(KR)
		sig := LFNoise{Interpolation: NoiseLinear, Freq: freq}.Rate(AR).Mul(gain)
		return Out{bus, sig}.Rate(AR)
	})
	same, err := def.CompareToFile("fixtures/LFNoise1Example.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang version")
	}
}
