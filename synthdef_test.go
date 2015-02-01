package gosc

import (
	"os"
	"testing"
)

func TestSineTone(t *testing.T) {
	_ = SineTone()
	// if st == nil {
	// 	t.Fail()
	// }
	// f, ce := os.Create("SineTone_go.scsyndef")
	// if ce != nil {
	// 	t.Fatal(ce)
	// }
	// we := st.Write(f)
	// if we != nil {
	// 	t.Fatal(we)
	// }
}

func TestReadSynthDef(t *testing.T) {
	f, err := os.Open("SineTone.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	_, err = ReadSynthDef(f)
	if err != nil {
		t.Fatal(err)
	}
}
