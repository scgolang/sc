package pattern

import "reflect"
import "testing"

func TestShuf(t *testing.T) {
	n := 10
	pat := Shuf(n, 1, 2, 3)
	count := 0
	for v := range pat {
		if i, isInt := v.(int); isInt {
			if i != 1 && i != 2 && i != 3 {
				t.Fatalf("%d was not in the provided list", i)
			}
		} else {
			t.Fatalf("value was of type %s", reflect.TypeOf(v).String())
		}
		count++
	}
	if count != n*3 {
		t.Fatalf("count was %d", count)
	}
}
