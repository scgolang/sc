package sc

import "testing"

// Out.ar(0, SoundIn.ar([0, 1]));
func TestSoundIn(t *testing.T) {
	defName := "SoundInTest0"
	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		return Out{
			Bus: C(0),
			Channels: SoundIn{
				Bus: C(0),
			}.Rate(AR),
		}.Rate(AR)
	}))

	defName = "SoundInTest00"
	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		return Out{
			Bus: C(0),
			Channels: SoundIn{
				Bus: Multi(C(0), C(0)),
			}.Rate(AR),
		}.Rate(AR)
	}))

	defName = "SoundInTest01"
	compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		return Out{
			Bus: C(0),
			Channels: SoundIn{
				Bus: Multi(C(0), C(1)),
			}.Rate(AR),
		}.Rate(AR)
	}))

	// FIXME
	// defName = "SoundInTest02"
	// compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
	// 	return Out{
	// 		Bus: C(0),
	// 		Channels: SoundIn{
	// 			Bus: Multi(C(0), C(2)),
	// 		}.Rate(AR),
	// 	}.Rate(AR)
	// }))

	// FIXME
	// defName = "SoundInTest12"
	// compareAndWrite(t, defName, NewSynthdef(defName, func(p Params) Ugen {
	// 	return Out{
	// 		Bus: C(0),
	// 		Channels: SoundIn{
	// 			Bus: Multi(C(1), C(2)),
	// 		}.Rate(AR),
	// 	}.Rate(AR)
	// }))
}
