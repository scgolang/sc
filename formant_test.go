package sc

import (
	"testing"
)

func TestFormant(t *testing.T) {
	defName := "FormantTest"

	// Out.ar(0, Formant.ar(XLine.kr(400,1000, 8), 2000, 800, 0.125));
	def := NewSynthdef(defName, func(p Params) Ugen {
		line := XLine{
			Start: C(400),
			End:   C(1000),
			Dur:   C(8),
		}.Rate(KR)

		return Out{
			Bus: C(0),
			Channels: Formant{
				FundFreq:    line,
				FormantFreq: C(2000),
				BWFreq:      C(800),
			}.Rate(AR).Mul(C(0.125)),
		}.Rate(AR)
	})
	compareAndWriteStructure(t, defName, def)
}
