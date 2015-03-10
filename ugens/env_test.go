package ugens

import (
	"testing"
)

func TestPerc(t *testing.T) {
	e := Env.Perc()
	expect := []float32{0, 2, -99, -99, 1, 0.01, 5, -4, 0, 1, 5, -4}
	verifyInputs(t, expect, e.InputsArray())
}

func TestLinen(t *testing.T) {
	e := Env.Linen()
	expect := []float32{0, 3, -99, -99, 1, 0.01, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0}
	verifyInputs(t, expect, e.InputsArray())
	expectWelch := []float32{0, 3, -99, -99, 1, 0.01, 4, 0, 1, 1, 4, 0, 0, 1, 4, 0}
	f := Env.Linen(0.01, 1, 1, 1, CurveWelch)
	verifyInputs(t, expectWelch, f.InputsArray())
}

func verifyInputs(t *testing.T, expect []float32, inputs []interface{}) {
	t.Logf("%v\n", inputs)
	errmsg := "expected %f for input %d but got %f"
	for i, in := range inputs {
		var val float32
		switch v := in.(type) {
		case int:
			val = float32(v)
		case float64:
			val = float32(v)
		case float32:
			val = v
		default:
			t.Fatalf("input %d was not a float or int (%v)", i, in)
		}
		if val != expect[i] {
			t.Fatalf(errmsg, expect[i], i, val)
		}
		// if floatVal, isFloat := in.(float32); isFloat {
		// 	if floatVal != expect[i] {
		// 		t.Fatalf(errmsg, expect[i], i, floatVal)
		// 	}
		// } else if intVal, isInt := in.(int); isInt {
		// 	if float32(intVal) != expect[i] {
		// 		t.Fatalf(errmsg, expect[i], i, intVal)
		// 	}
		// } else {
		// 	t.Fatalf("input %d was not a float or int (%v)", i, in)
		// }
	}
}
