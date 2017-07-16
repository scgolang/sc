package sc

import (
	"fmt"
	"os"
	"testing"
	// "github.com/scgolang/syndef/defdiff"
)

// func compareAndWriteStructure(t *testing.T, name string, def *Synthdef) {
// 	sclangPath := filepath.Join("testdata", fmt.Sprintf("%s.scsyndef", name))
// 	sclangVersion, err := ReadSynthdef(sclangPath)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	diff, err := defdiff.Do(def, sclangVersion)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	for _, msgs := range diff {
// 		t.Error(msgs[0] + ", " + msgs[1])
// 	}
// }

func compareAndWrite(t *testing.T, name string, def *Synthdef) {
	f, err := os.Create(fmt.Sprintf("%s.gosyndef", name))
	if err != nil {
		t.Fatal(err)
	}
	err = def.Write(f)
	if err != nil {
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
