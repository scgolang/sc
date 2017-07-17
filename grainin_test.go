package sc

import "testing"

func TestGrainIn(t *testing.T) {
	const defName = "GrainInTest"

	// arg gate = 1, amp = 1, envbuf;
	// var pan, env;
	// // use mouse x to control panning
	// pan = MouseX.kr(-1, 1);
	// env = EnvGen.kr(
	//     Env([0, 1, 0], [1, 1], \sin, 1),
	//     gate,
	//     levelScale: amp,
	//     doneAction: 2);
	// Out.ar(0,
	//     GrainIn.ar(2, Impulse.kr(32), 1, PinkNoise.ar * 0.05, pan, envbuf) * env)
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		var (
			gate   = p.Add("gate", 1)
			amp    = p.Add("amp", 1)
			envbuf = p.Add("envbuf", 0)
			pan    = MouseX{Min: C(-1), Max: C(1)}.Rate(KR)
		)
		env := EnvGen{
			Env: Env{
				Levels:      []Input{C(0), C(1), C(0)},
				Times:       []Input{C(1), C(1)},
				Curve:       "sine",
				ReleaseNode: C(1),
			},
			Gate:       gate,
			LevelScale: amp,
			Done:       FreeEnclosing,
		}.Rate(KR)

		return Out{
			Bus: C(0),
			Channels: GrainIn{
				NumChannels: 2,
				Trigger:     Impulse{Freq: C(32)}.Rate(KR),
				Dur:         C(1),
				In:          PinkNoise{}.Rate(AR).Mul(C(0.05)),
				Pan:         pan,
				EnvBuf:      envbuf,
			}.Rate(AR).Mul(env),
		}.Rate(AR)
	}))
}
