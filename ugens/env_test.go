package ugens

import (
	"testing"
)

func TestPerc(t *testing.T) {
	e := Env.Perc()
	expect := []float32{0, 2, -99, -99, 1, 0.01, 5, -4, 0, 1, 5, -4}
	errmsg := "expected %f for input %d but got %f"
	arr := e.InputsArray()
	t.Logf("arr = %v\n", arr)
	for i, in := range arr {
		if floatVal, isFloat := in.(float32); isFloat {
			if floatVal != expect[i] {
				t.Fatalf(errmsg, expect[i], i, floatVal)
			}
		} else {
			t.Fatalf("input %d was not a float (%v)", i, in)
		}
	}
}
