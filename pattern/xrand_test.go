package pattern

import "reflect"
import "testing"

func TestXrand(t *testing.T) {
	pat := Xrand(10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	last := -1
	for v := range pat {
		if i, isInt := v.(int); isInt {
			if i == last {
				t.Fatal("got same value twice in a row")
				last = i
			}
		} else {
			t.Fatalf("value was of type %s", reflect.TypeOf(v).String())
		}
	}
}
