package sc

import (
	"testing"
)

func TestBrownNoise(t *testing.T) {
	defName := "BrownNoiseTest"
	def := NewSynthdef(defName, func(p Params) Ugen {
		bus, gain := C(0), C(0.1)
		noise := BrownNoise{}.Rate(AR).MulAdd(C(100), C(200))
		sig := SinOsc{Freq: noise}.Rate(AR)
		return Out{bus, sig.Mul(gain)}.Rate(AR)
	})
	same, err := def.CompareToFile("fixtures/BrownNoiseTest.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang version")
	}
}
