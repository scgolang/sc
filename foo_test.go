package sc

import . "github.com/scgolang/sc/types"
import . "github.com/scgolang/sc/ugens"
import "testing"

func TestFoo(t *testing.T) {
	name := "foo"
	def := NewSynthdef(name, func(p Params) Ugen {
		bus := C(0)
		blip := Blip{}.Rate(AR)
		sine := SinOsc{}.Rate(AR)
		return Out{bus, sine.Mul(blip)}.Rate(AR)
	})
	same, err := def.Compare(`{
        Out.ar(0, SinOsc.ar() * Blip.ar());
    }`)
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdefs were different")
	}
}
