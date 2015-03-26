package ugens

import (
	. "github.com/briansorahan/sc/types"
	"reflect"
	"testing"
)

func TestexpandInputs(t *testing.T) {
	a := []Input{Multi(C(1), C(2), C(3)), Multi(C(4), C(5))}
	b := expand(a...)
	// should be [ [1,4], [2,5], [3,4] ]
	expect := [][]Input{
		{ C(1), C(4) },
		{ C(2), C(5) },
		{ C(3), C(4) },
	}
	if !reflect.DeepEqual(expect, b) {
		t.Fatalf("expandInputs returned %v", b)
	}
}
