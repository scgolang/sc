package pattern

import "reflect"
import "testing"

func TestArith(t *testing.T) {
	pat := Arith(3, 4, 5)
	l := make([]float64, 0)
	for v := range pat {
		l = append(l, v)
	}
	expect := []float64{3, 7, 11, 15, 19}
	if !reflect.DeepEqual(expect, l) {
		t.Fatal("Arith did not generate the expected list")
	}
}
