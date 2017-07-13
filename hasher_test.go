package sc

import (
	"testing"
)

func TestHasher(t *testing.T) {
	defName := "HasherTest"

	def := NewSynthdef(defName, func(p Params) Ugen {
		return Out{
			Bus: C(0),
			Channels: SinOsc{
				Freq: Hasher{
					In: MouseX{Min: C(0), Max: C(10)}.Rate(KR),
				}.Rate(KR).MulAdd(C(300), C(500)),
			}.Rate(AR),
		}.Rate(AR)
	})
	compareAndWrite(t, defName, def)
}
