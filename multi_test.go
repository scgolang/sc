package sc

import (
	"testing"
)

func TestMulti(t *testing.T) {
	def := NewSynthdef("SimpleMulti", func(p Params) Ugen {
		bus, freq := C(0), Multi(C(440), C(441))
		sine := SinOsc{Freq: freq}.Rate(AR)
		return Out{bus, sine}.Rate(AR)
	})
	same, err := def.CompareToFile("fixtures/SimpleMulti.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang version")
	}
}
