package sc

import . "github.com/scgolang/sc/types"
import . "github.com/scgolang/sc/ugens"
import "testing"

func TestSinOsc(t *testing.T) {
	name := "SineTone"
	def := NewSynthdef(name, func(params Params) Ugen {
		bus, freq := C(0), C(440)
		sine := SinOsc{Freq: freq}.Rate(AR)
		return Out{bus, sine}.Rate(AR)
	})
	if def == nil {
		t.Fatalf("nil synthdef")
	}
	compareAndWrite(t, name, def)
}
