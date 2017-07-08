package sc

import "testing"

func TestHPF(t *testing.T) {
	name := "HPFExample"
	def := NewSynthdef(name, func(params Params) Ugen {
		bus, In := C(0), SinOsc{Freq: C(220)}.Rate(AR)
		src := HPF{In: In}.Rate(AR)
		return Out{bus, src}.Rate(AR)
	})
	same, err := def.CompareToFile("testdata/HPFExample.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang version")
	}
}
