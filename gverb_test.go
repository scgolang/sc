package sc

import "testing"

func TestGVerb(t *testing.T) {
	name := "GVerbExample"
	def := NewSynthdef(name, func(params Params) Ugen {
		bus, In := C(0), SinOsc{Freq: C(220)}.Rate(AR)
		src := GVerb{In: In}.Rate(AR)
		return Out{bus, src}.Rate(AR)
	})
	same, err := def.CompareToFile("testdata/GVerbExample.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang version")
	}
}
