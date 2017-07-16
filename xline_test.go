package sc

import (
	"testing"
)

func TestXLine(t *testing.T) {
	defName := "XLineTest"

	// Out.ar(0, SinOsc.ar(Line.kr(200, 17000, 10), 0, 0.1));
	def := NewSynthdef(defName, func(p Params) Ugen {
		line := XLine{
			Start: C(200),
			End:   C(17000),
			Dur:   C(10),
		}.Rate(KR)

		return Out{
			Bus:      C(0),
			Channels: SinOsc{Freq: line}.Rate(AR).Mul(C(0.1)),
		}.Rate(AR)
	})
	compareAndWriteStructure(t, defName, def)
}
