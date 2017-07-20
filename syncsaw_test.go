package sc

import "testing"

func TestSyncSaw(t *testing.T) {
	const defName = "SyncSawTest"

	// Out.ar(0, SyncSaw.ar(800, Line.kr(800, 1600, 0.01)));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		return Out{
			Bus: C(0),
			Channels: A(SyncSaw{
				SyncFreq: C(800),
				SawFreq: K(Line{
					Start: C(800),
					End:   C(1600),
					Dur:   C(0.01),
				}),
			}),
		}.Rate(AR)
	}))
}
