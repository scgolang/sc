package sc

import . "github.com/scgolang/sc/ugens"
import . "github.com/scgolang/sc/types"
import "testing"

func TestIntegrator(t *testing.T) {
	def := NewSynthdef("IntegratorExample", func(p Params) Ugen {
		pulse := LFPulse{C(375), C(0.2), C(0.1)}.Rate(AR)
		x := MouseX{C(0.01), C(0.999), C(1), C(0.2)}.Rate(KR)
		sig := Integrator{pulse, x}.Rate(AR)
		return Out{C(0), sig}.Rate(AR)
	})
	same, err := def.Compare(`{
        var pulse = LFPulse.ar(1500 / 4, 0.2, 0.1);
        var mouse = MouseX.kr(0.01, 0.999, 1);
        Out.ar(0, Integrator.ar(pulse, mouse));
    }`)
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang version")
	}
}
