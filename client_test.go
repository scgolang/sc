package sc

import (
	"os"
	"path"
	"testing"
	"time"
)

// skipIfNoScsynth skips a test if scsynth is not running.
// A timeout of 2 seconds is used to attempt a connection to scsynth
// since when it is running it should be local on the same machine
// and we should be able to connect very quickly.
func skipIfNoScsynth(t *testing.T, client *Client) {
	if _, err := client.Status(2 * time.Second); err != nil {
		t.SkipNow()
	}
}

// This test requires a SuperCollider server to be running.
//
//     scsynth -u 57120
//
func TestClient(t *testing.T) {
	client, err := NewClient("udp", "127.0.0.1:57200", "127.0.0.1:57120", 5*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// TODO: only skip if there is not a supercollider server running
	skipIfNoScsynth(t, client)

	defer func() {
		_ = <-client.doneChan
		_ = client.Close()
	}() // Best effort.

	// read a buffer
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	audioFile := path.Join(cwd, "kalimba_mono.wav")

	buf, err := client.ReadBuffer(audioFile, 0)
	if err != nil {
		t.Fatal(err)
	}

	if buf == nil {
		t.Fatalf("got nil buffer")
	}
}
