package sc

import (
	"testing"
)

func TestMidicps(t *testing.T) {
	val := Midicps(60)
	if int(val) != 261 {
		t.Fatalf("got %d, expected %d", int(val), 261)
	}
}
