package sc

import "testing"

func TestLFTri(t *testing.T) {
	def := NewSynthdef("LFTriExample", func(p Params) Ugen {
		bus := C(0)
		freq := LFTri{C(4), C(0)}.Rate(KR).MulAdd(C(200), C(400))
		sig := LFTri{freq, C(0)}.Rate(AR).Mul(C(0.1))
		return Out{bus, sig}.Rate(AR)
	})
	same, err := def.CompareToFile("fixtures/LFTriExample.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang version")
	}
}
