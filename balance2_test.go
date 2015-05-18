package sc

import . "github.com/scgolang/sc/types"
import . "github.com/scgolang/sc/ugens"
import "os"
import "testing"

func NoTestBalance2(t *testing.T) {
    // var l = LFSaw.ar(44);
    // var r = Pulse.ar(33);
    // var pos = FSinOsc.kr(0.5);
    // var level = 0.1;
    // Out.ar(0, Balance2.ar(l, r, pos, level));
	def := NewSynthdef("Balance2Example", func(p *Params) Ugen {
		// l, r := LFSaw{44}.Rate(AR), Pulse{33}.Rate(AR)
		return nil
	})
	f, err := os.Create("Balance2Example.gosyndef")
	if err != nil {
		t.Fatal(err)
	}
	err = def.Write(f)	
	if err != nil {
		t.Fatal(err)
	}
	same, err := def.CompareToFile("Balance2Example.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef is not the same as sclang version")
	}
}
