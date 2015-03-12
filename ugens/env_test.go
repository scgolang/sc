package ugens

import (
	. "github.com/briansorahan/sc/types"
	"testing"
)

func TestPerc(t *testing.T) {
	e := Env.Perc(C(0.01), C(1), C(1), C(-4))
	expect := []float32{0, 2, -99, -99, 1, 0.01, 5, -4, 0, 1, 5, -4}
	verifyInputs(t, expect, e.InputsArray())
}

func TestLinen(t *testing.T) {
	e := Env.Linen(C(0.01), C(1), C(1), C(1), CurveLinear)
	expect := []float32{0, 3, -99, -99, 1, 0.01, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0}
	verifyInputs(t, expect, e.InputsArray())
	expectWelch := []float32{0, 3, -99, -99, 1, 0.01, 4, 0, 1, 1, 4, 0, 0, 1, 4, 0}
	f := Env.Linen(C(0.01), C(1), C(1), C(1), CurveWelch)
	verifyInputs(t, expectWelch, f.InputsArray())
}

func verifyInputs(t *testing.T, expect []float32, inputs []Input) {
	t.Logf("%v\n", inputs)
	errmsg := "expected %f for input %d but got %f"
	for i, in := range inputs {
		var val float32
		switch v := in.(type) {
		case C:
			val = float32(v)
		default:
			t.Fatalf("input %d was not a float or int (%v)", i, in)
		}
		if val != expect[i] {
			t.Fatalf(errmsg, expect[i], i, val)
		}
	}
}
