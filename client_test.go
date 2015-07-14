package sc

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func TestClient(t *testing.T) {
	client := NewClient("127.0.0.1", ScsynthDefaultPort)
	err := client.Connect("127.0.0.1", 57200)
	if err != nil {
		t.Fatal(err)
	}
	if client == nil {
		t.Fatal(fmt.Errorf("NewClient returned nil"))
	}
	
	// get status
	status, err := client.Status()
	if err != nil {
		t.Fatal(err)
	}
	if status == nil {
		t.Fatalf("got nil status")
	}
	
	// read a buffer
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	audioFile := path.Join(cwd, "kalimba_mono.wav")
	buf, err := client.ReadBuffer(audioFile)
	if err != nil {
		t.Fatal(err)
	}
	if buf == nil {
		t.Fatalf("got nil buffer")
	}

}
