package sc

import "testing"

func TestFoo(t *testing.T) {
	name := "foo"
	def := NewSynthdef(name, func(p Params) Ugen {
		bus := C(0)
		blip := Blip{}.Rate(AR)
		sine := SinOsc{}.Rate(AR)
		return Out{bus, sine.Mul(blip)}.Rate(AR)
	})
	same, err := def.CompareToFile("fixtures/foo.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdefs were different")
	}
}
