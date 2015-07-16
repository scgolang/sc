package sc

import (
	. "github.com/scgolang/sc/types"
	. "github.com/scgolang/sc/ugens"
	"testing"
)

func TestMulti(t *testing.T) {
	def := NewSynthdef("SimpleMulti", func(p Params) Ugen {
		bus, freq := C(0), Multi(C(440), C(441))
		sine := SinOsc{Freq: freq}.Rate(AR)
		return Out{bus, sine}.Rate(AR)
	})
	same, err := def.Compare(`{
        var sine = SinOsc.ar([440, 441]);
        Out.ar(0, sine);
    }`)
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang version")
	}
}
