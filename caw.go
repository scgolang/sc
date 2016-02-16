package sc

import "fmt"
import "os"
import "testing"

func compareAndWrite(t *testing.T, name string, def *Synthdef) {
	f, err := os.Create(fmt.Sprintf("%s.gosyndef", name))
	if err != nil {
		t.Fatal(err)
	}
	err = def.Write(f)
	if err != nil {
		t.Fatal(err)
	}
	same, err := def.CompareToFile(fmt.Sprintf("fixtures/%s.scsyndef", name))
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("%s is not the same as sclang version", name)
	}
}
