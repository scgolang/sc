package sc

import (
	. "github.com/scgolang/sc/types"
	. "github.com/scgolang/sc/ugens"
	"testing"
)

func TestEnvGen(t *testing.T) {
	def := NewSynthdef("Envgen1", func(p Params) Ugen {
		bus := C(0)
		attack, release := C(0.01), C(1)
		level, curvature := C(1), C(-4)
		perc := EnvPerc{attack, release, level, curvature}
		gate, levelScale, levelBias, timeScale := C(1), C(1), C(0), C(1)
		ampEnv := EnvGen{perc, gate, levelScale, levelBias, timeScale, FreeEnclosing}.Rate(KR)
		noise := PinkNoise{}.Rate(AR).Mul(ampEnv)
		return Out{bus, noise}.Rate(AR)
	})
	same, err := def.Compare(`{
        Out.ar(0, PinkNoise.ar() * EnvGen.kr(Env.perc, doneAction: 2));
    }`)
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang version")
	}
}
