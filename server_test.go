package sc

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestServerStatus(t *testing.T) {
	options := ServerOptions{
		EchoScsynthStdout: false,
	}
	s, err := NewServer("127.0.0.1", scsynthPort, options)
	if err != nil {
		t.Fatal(err)
	}
	if s == nil {
		t.Fatal(fmt.Errorf("NewServer returned nil"))
	}
	done := s.Run()
	if err != nil {
		s.Close()
		t.Fatal(err)
	}

	time.Sleep(1000 * time.Millisecond)

	// get status
	log.Println("getting status")
	err = s.Status()
	if err != nil {
		s.Close()
		t.Fatal(err)
	}
	log.Println("got status")
	// quit
	time.Sleep(10 * time.Millisecond)
	err = s.Quit()
	if err != nil {
		t.Fatal(err)
	}
	log.Println("sent quit message")
	err = <-done
	if err != nil {
		t.Fatal(err)
	}
}
