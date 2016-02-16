package sc

import "testing"

func TestIntegrator(t *testing.T) {
	def := NewSynthdef("IntegratorExample", func(p Params) Ugen {
		pulse := LFPulse{C(375), C(0.2), C(0.1)}.Rate(AR)
		x := MouseX{C(0.01), C(0.999), C(1), C(0.2)}.Rate(KR)
		sig := Integrator{pulse, x}.Rate(AR)
		return Out{C(0), sig}.Rate(AR)
	})
	same, err := def.CompareToFile("fixtures/IntegratorExample.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang version")
	}
}
