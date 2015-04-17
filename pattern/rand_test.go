package pattern

import "reflect"
import "testing"

func TestRand(t *testing.T) {
	pat := Rand(10, 1, 2, 3, 4)
	for val := range pat {
		if i, isint := val.(int); isint {
			if i != 1 && i != 2 && i != 3 && i != 4 {
				t.Fatalf("%d was not one of the provided values", val)
			}
		} else {
			t.Fatalf("expected int but got %s", reflect.TypeOf(val).String())
		}
	}
}
