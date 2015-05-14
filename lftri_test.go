package sc

import . "github.com/scgolang/sc/types"
import . "github.com/scgolang/sc/ugens"
import "testing"

func TestLFTri(t *testing.T) {
	def := NewSynthdef("LFTriExample", func(p *Params) UgenNode {
		// sclang:
		// var freq = LFTri.kr(4, 0, 200, 400);
		// Out.ar(0, LFTri.ar(freq, 0, 0.1));
		bus := C(0)
		freq := LFTri{C(4), C(0)}.Rate(KR).MulAdd(C(200), C(400))
		sig := LFTri{freq, C(0)}.Rate(AR).Mul(C(0.1))
		return Out{bus, sig}.Rate(AR)
	})
	compareAndWrite(t, "LFTriExample", def);
}
