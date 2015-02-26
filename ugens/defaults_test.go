package ugens

import (
	"testing"
)

func TestApplyDefaults(t *testing.T) {
	defaults := []float32{440, 2}
	vals := applyDefaults(defaults, 1)
	numVals := len(vals)
	if numVals != 2 {
		t.Fatalf("expected numVals to be 2, but got %d", numVals)
	}
	if vals[0] != 1 {
		t.Fatalf("expected vals[0] to be 1, but got %d", vals[0])
	}
	if vals[1] != float32(2) {
		t.Fatalf("expected vals[1] to be 2, but got %d", vals[1])
	}
}
