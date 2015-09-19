package sc

import (
	"testing"
)

func TestGrainBuf(t *testing.T) {
	defName := "GrainBufTest"
	def := NewSynthdef(defName, func(p Params) Ugen {
		bus, bufnum := C(0), C(0)
		channels := 1
		sig := GrainBuf{
			NumChannels: channels,
			BufNum:      bufnum,
		}.Rate(AR)
		return Out{bus, sig}.Rate(AR)
	})
	compareAndWrite(t, defName, def)
}
