package ugens

import (
	. "github.com/briansorahan/sc/types"
	"testing"
)

func TestPerc(t *testing.T) {
	e := EnvPerc{}
	expect := []float32{0, 2, -99, -99, 1, 0.01, 5, -4, 0, 1, 5, -4}
	verifyInputs(t, expect, e.InputsArray())
}

func TestLinen(t *testing.T) {
	e := EnvLinen{}
	expect := []float32{0, 3, -99, -99, 1, 0.01, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0}
	verifyInputs(t, expect, e.InputsArray())
	f := EnvLinen{CurveType:CurveWelch}
	expectWelch := []float32{0, 3, -99, -99, 1, 0.01, 4, 0, 1, 1, 4, 0, 0, 1, 4, 0}
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
