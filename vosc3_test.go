package sc

import (
	"testing"
)

func TestVOsc3(t *testing.T) {
	defName := "VOsc3Test"

	// arg bufnum = 0;
	// var line1 = XLine.kr(2000, 200, 0.5);
	// var line2 = XLine.kr(2000, 200, 1.5);
	// var line3 = XLine.kr(2000, 200, 4.5);
	// Out.ar(0, VOsc3.ar(bufnum, line1, line2, line3));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		var (
			bufnum = p.Add("bufnum", 0)
		)
		line1 := K(XLine{
			Start: C(2000),
			End:   C(200),
			Dur:   C(0.5),
		})
		line2 := K(XLine{
			Start: C(2000),
			End:   C(200),
			Dur:   C(1.5),
		})
		line3 := K(XLine{
			Start: C(2000),
			End:   C(200),
			Dur:   C(4.5),
		})
		return Out{
			Bus: C(0),
			Channels: A(VOsc3{
				BufNum: bufnum,
				Freq1:  line1,
				Freq2:  line2,
				Freq3:  line3,
			}),
		}.Rate(AR)
	}))
}
