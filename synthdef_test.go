package gosc

import (
	"testing"
)

func TestTestTone(t *testing.T) {
	tt := TestTone()
	if tt == nil {
		t.Fail()
	}
}
