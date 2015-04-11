package pattern

import "testing"

func TestWhite(t *testing.T) {
	pat := White(2, 5, 3)
	for v := range pat {
		if v < float64(2) || v > float64(5) {
			t.Fatal("value outside expected range")
		}
	}
}
