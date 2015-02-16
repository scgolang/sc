package ugens

import (
	"testing"
)

func TestOutNoArguments(t *testing.T) {
	defer func() {
		if val := recover(); val == nil {
			t.Fatalf("expected panic when Out.Ar is called with no arguments")
		}
	}()
	Out.Ar()
}

