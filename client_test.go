package sc

import (
	"fmt"
	"testing"
)

func TestClientStatus(t *testing.T) {
	s, err := NewClient("127.0.0.1", scsynthPort)
	if err != nil {
		t.Fatal(err)
	}
	if s == nil {
		t.Fatal(fmt.Errorf("NewClient returned nil"))
	}
	// get status
	err = s.Status()
	if err != nil {
		t.Fatal(err)
	}
}
