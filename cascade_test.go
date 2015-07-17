package sc

import (
	. "github.com/scgolang/sc/types"
	. "github.com/scgolang/sc/ugens"
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
	same, err := def.Compare(`{
		var mod1 = SinOsc.ar([440, 441]);
		var mod2 = SinOsc.ar(mod1);
		Out.ar(0, SinOsc.ar(mod2));
    }`)
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang version")
	}
}
