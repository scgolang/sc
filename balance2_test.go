package sc

// import (
// 	. "github.com/scgolang/sc/types"
// 	. "github.com/scgolang/sc/ugens"
// 	"testing"
// )

// func TestBalance2(t *testing.T) {
// 	defName := "Balance2Test"
// 	def := NewSynthdef(defName, func(p Params) Ugen {
// 		// var l = LFSaw.ar(44);
// 		// var r = Pulse.ar(33);
// 		// var pos = FSinOsc.kr(0.5);
// 		// var level = 0.1;
// 		// Out.ar(0, Balance2.ar(l, r, pos, level));
// 		l, r := LFSaw{C(44)}.Rate(AR), Pulse{33}.Rate(AR)
// 		return nil
// 	})
// 	f, err := os.Create("Balance2Example.gosyndef")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	err = def.Write(f)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	same, err := def.CompareToFile("Balance2Example.scsyndef")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if !same {
// 		t.Fatalf("synthdef is not the same as sclang version")
// 	}
// }
