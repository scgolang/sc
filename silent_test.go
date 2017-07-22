package sc

import (
	"testing"
)

func TestSilent(t *testing.T) {
	const defName = "SilentTest"

	// Out.ar(0, Silent.ar(2));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		return Out{
			Bus:      C(0),
			Channels: Silent{NumChannels: 2}.Rate(AR),
		}.Rate(AR)
	}))
}
