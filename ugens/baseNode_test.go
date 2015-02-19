package ugens

import (
	"testing"
)

func TestEnsureOutput(t *testing.T) {
	n := newNode("foo", 2)
	n.EnsureOutput()
}
