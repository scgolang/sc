package sc

import (
	"log"
	"os"
	"path"
	"testing"
)

// This test requires a SuperCollider server to be running.
//
//     scsynth -u 57120
//
func TestClient(t *testing.T) {
	client, err := NewClient("udp", "127.0.0.1:57200", "127.0.0.1:57120")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = client.Close() }() // Best effort.

	// get status
	// status, err := client.GetStatus()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// if status == nil {
	// 	t.Fatalf("got nil status")
	// }

	// read a buffer
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	audioFile := path.Join(cwd, "kalimba_mono.wav")

	log.Println("reading buffer...")
	buf, err := client.ReadBuffer(audioFile, 0)
	if err != nil {
		t.Fatal(err)
	}
	log.Println("done reading buffer.")
	if buf == nil {
		t.Fatalf("got nil buffer")
	}
}
