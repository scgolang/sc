package sc

import (
	"testing"
)

func TestCascade(t *testing.T) {
	def := NewSynthdef("CascadeExample", func(p Params) Ugen {
		bus := C(0)
		freq := Multi(C(440), C(441))
		mod1 := SinOsc{Freq: freq}.Rate(AR)
		mod2 := SinOsc{Freq: mod1}.Rate(AR)
		return Out{bus, SinOsc{Freq: mod2}.Rate(AR)}.Rate(AR)
	})
	same, err := def.CompareToFile("fixtures/CascadeExample.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang version")
	}
}
