package sc

import (
	"fmt"
	"testing"
)

func TestAr(t *testing.T) {
	_, err := Ar("Foo")
	if err != nil {
		t.Error(fmt.Errorf("failed to create dummy ugen with no arguments"))
	}
}
