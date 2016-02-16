package sc

import (
	"testing"
)

func TestSinOsc(t *testing.T) {
	name := "SineTone"
	def := NewSynthdef(name, func(params Params) Ugen {
		bus, freq := C(0), C(440)
		sine := SinOsc{Freq: freq}.Rate(AR)
		return Out{bus, sine}.Rate(AR)
	})
	same, err := def.CompareToFile("fixtures/SineTone.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang version")
	}
}
