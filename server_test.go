package sc

import (
	"fmt"
	"testing"
)

func TestNewServer(t *testing.T) {
	s := NewServer(NetAddr{"localhost", DefaultPort})
	if s == nil {
		t.Fatal(fmt.Errorf("NewServer returned nil"))
	}
}

func TestServerBoot(t *testing.T) {
	s := NewServer(NetAddr{"localhost", DefaultPort})
	if s == nil {
		t.Fatal(fmt.Errorf("NewServer returned nil"))
	}
	err := s.Boot()
	if err != nil {
		t.Fatal(err)
	}
	err = s.Close()
	if err != nil {
		t.Fatal(err)
	}
}
