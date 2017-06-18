package sc

import "testing"

func TestBAllPass(t *testing.T) {
	const defName = "BAllPassExample"

	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		mousex := MouseX{
			Min:  C(10),
			Max:  C(18000),
			Warp: C(1),
		}.Rate(KR)

		saw := Saw{}.Rate(AR)

		sig := BAllPass{
			In:   saw,
			Freq: mousex,
			RQ:   C(0.8),
		}.Rate(AR)

		return Out{
			Bus:      C(0),
			Channels: sig,
		}.Rate(AR)
	}))
}
