package sc

import . "github.com/briansorahan/sc/types"
import . "github.com/briansorahan/sc/ugens"
import "testing"

func TestLFSaw(t *testing.T) {
	def := NewSynthdef("LFSawExample", func(p *Params) UgenNode {
		lfoFreq, lfoPhase := C(4), C(0)
		bus, gain := C(0), C(0.1)
		freq := LFSaw{lfoFreq, lfoPhase}.Rate(KR).MulAdd(C(200), C(400))
		sig := LFSaw{freq, C(0)}.Rate(AR).Mul(gain)
		return Out{bus, sig}.Rate(AR)
	})
	compareAndWrite(t, "LFSawExample", def);
}
