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
	if def == nil {
		t.Fatalf("nil synthdef")
	}
	compareAndWrite(t, name, def)
}
