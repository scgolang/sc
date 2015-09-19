package sc

import (
	"testing"
)

func TestCOsc(t *testing.T) {
	const defName = "COscTest"
	def := NewSynthdef(defName, func(p Params) Ugen {
		bus, bufnum := C(0), C(0)
		freq, beats := C(200), C(0.7)
		gain := C(0.25)
		sig := COsc{BufNum: bufnum, Freq: freq, Beats: beats}.Rate(AR)
		return Out{bus, sig.Mul(gain)}.Rate(AR)
	})
	compareAndWrite(t, defName, def)
}
