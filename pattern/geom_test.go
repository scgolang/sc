package pattern

import "reflect"
import "testing"

func TestGeom(t *testing.T) {
	pat := Geom(1, 2, 5)
	l := make([]float64, 0)
	for v := range pat {
		l = append(l, v)
	}
	expect := []float64{1, 2, 4, 8, 16}
	if !reflect.DeepEqual(expect, l) {
		t.Fatal("Geom did not generate expected list")
	}
}
