package sc

import "testing"

func TestTGrains(t *testing.T) {
	name := "TGrainsExample"
	def := NewSynthdef(name, func(params Params) Ugen {
		bus, numChannels, trate := C(0), 2, Impulse{Freq: C(4)}.Rate(AR)
		tg := TGrains{NumChannels: numChannels, Trigger: trate}.Rate(AR)
		return Out{bus, tg}.Rate(AR)
	})
	same, err := def.CompareToFile("testdata/TGrainsExample.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang version")
	}
}
