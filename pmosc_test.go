package sc

import (
	"testing"
)

func TestPMOsc(t *testing.T) {
	defName := "PMOscTest"
	def := NewSynthdef(defName, func(p Params) Ugen {
		// Out.ar(0, PMOsc.ar(Line.kr(600, 900, 5), 600, 3, 0, 0.1));
		bus, gain := C(0), C(0.1)
		line := Line{Start: C(600), End: C(900), Dur: C(5)}.Rate(KR)
		sig := PMOsc{CarFreq: line, ModFreq: C(600), PMIndex: C(3)}.Rate(AR).Mul(gain)
		return Out{bus, sig}.Rate(AR)
	})
	compareAndWrite(t, defName, def)
}
