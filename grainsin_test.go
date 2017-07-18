package sc

import (
	"testing"
)

func TestGrainSin(t *testing.T) {
	const defName = "GrainSinTest"

	// arg gate = 1, amp = 1, envbuf;
	// var pan, env;
	// // use mouse x to control panning
	// pan = MouseX.kr(-1, 1);
	// env = EnvGen.kr(
	//         Env([0, 1, 0], [1, 1], \sin, 1),
	//         gate,
	//         levelScale: amp,
	//         doneAction: 2
	// );
	// Out.ar(0, GrainSin.ar(2, Impulse.kr(10), 0.1, 440 + freqdev, pan, envbuf) * env);
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		var (
			amp    = p.Add("amp", 1)
			envbuf = p.Add("envbuf", 0)
			gate   = p.Add("gate", 1)
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

		freqdev := WhiteNoise{}.Rate(KR).MulAdd(MouseY{Max: C(400)}.Rate(KR), C(440))
		return Out{
			Bus: C(0),
			Channels: GrainSin{
				NumChannels: 2,
				Trigger:     Impulse{Freq: C(10)}.Rate(KR),
				Dur:         C(0.1),
				Freq:        freqdev,
				Pan:         pan,
				EnvBuf:      envbuf,
			}.Rate(AR).Mul(env),
		}.Rate(AR)
	}))
}
