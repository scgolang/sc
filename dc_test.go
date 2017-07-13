package sc

import (
	"testing"
)

func TestDC(t *testing.T) {
	defName := "DCTest"
	def := NewSynthdef(defName, func(p Params) Ugen {
		return Out{
			Bus:      C(0),
			Channels: DC{}.Rate(AR),
		}.Rate(AR)
	})
	compareAndWrite(t, defName, def)
}
