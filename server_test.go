package sc

import (
	"fmt"
	"testing"
)

func TestServerStatus(t *testing.T) {
	addr := NetAddr{"127.0.0.1", defaultPort}
	options := ServerOptions{
		EchoScsynthStdout: true,
	}
	s := NewServer(addr, options)
	if s == nil {
		t.Fatal(fmt.Errorf("NewServer returned nil"))
	}
	err := s.Start()
	if err != nil {
		t.Fatal(err)
	}
}
