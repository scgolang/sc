package sc

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func compareAndWriteStructure(t *testing.T, name string, def *Synthdef) {
	f, err := os.Create(fmt.Sprintf("%s.gosyndef", name))
	if err != nil {
		t.Fatal(err)
	}
	if err := def.Write(f); err != nil {
		t.Fatal(err)
	}
	scf, err := os.Open(filepath.Join("testdata", fmt.Sprintf("%s.scsyndef", name)))
	if err != nil {
		t.Fatal(err)
	}
	sclangVersion, err := ReadSynthdef(scf)
	if err != nil {
		t.Fatal(err)
	}
	for _, msgs := range def.Diff(sclangVersion) {
		t.Error(msgs[0] + ", " + msgs[1])
	}
}

func compareAndWrite(t *testing.T, name string, def *Synthdef) {
	f, err := os.Create(fmt.Sprintf("%s.gosyndef", name))
	if err != nil {
		t.Fatal(err)
	}
	if err := def.Write(f); err != nil {
		t.Fatal(err)
	}
	same, err := def.CompareToFile(fmt.Sprintf("testdata/%s.scsyndef", name))
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("%s is not the same as sclang version", name)
	}
}
