package sc

import (
	"testing"

	. "github.com/scgolang/sc/types"
	. "github.com/scgolang/sc/ugens"
)

func TestCrackle(t *testing.T) {
	name := "CrackleTest"
	def := NewSynthdef(name, func(params Params) Ugen {
		//arg bufnum = 0;
		//var crack = Crackle.ar(Line.kr(1.0, 2.0, 3), 0.5, 0.5);
		//Out.ar(0, crack);

		bus, chaos := C(0), Line{C(1.0), C(2.0), C(3), DoNothing}.Rate(KR)
		sig := Crackle{chaos}.Rate(AR).MulAdd(C(0.5), C(0.5))
		return Out{bus, sig}.Rate(AR)

	})
	if def == nil {
		t.Fatalf("nil synthdef")
	}
	compareAndWrite(t, name, def)
}
