package sc

import (
	"testing"
)

func TestClipNoise(t *testing.T) {
	defName := "ClipNoiseTest"
	def := NewSynthdef(defName, func(p Params) Ugen {
		bus, gain := C(0), C(0.2)
		sig := ClipNoise{}.Rate(AR).Mul(gain)
		return Out{bus, sig}.Rate(AR)
	})
	compareAndWrite(t, defName, def)
}
