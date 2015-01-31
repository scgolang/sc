package gosc

import (
	"os"
	"testing"
)

func TestSineTone(t *testing.T) {
	st := SineTone()
	if st == nil {
		t.Fail()
	}
	f, ce := os.Create("SineTone_go.scsyndef")
	if ce != nil {
		t.Fatal(ce)
	}
	we := st.Write(f)
	if we != nil {
		t.Fatal(we)
	}
}
