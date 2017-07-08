package sc

import "testing"

func TestWarp1(t *testing.T) {
	name := "Warp1Example"
	def := NewSynthdef(name, func(params Params) Ugen {
		bus, numChannels := C(0), 2
		src := Warp1{NumChannels: numChannels}.Rate(AR)
		return Out{bus, src}.Rate(AR)
	})
	same, err := def.CompareToFile("testdata/Warp1Example.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang version")
	}
}
