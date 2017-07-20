package sc

import (
	"testing"
)

func TestKlang(t *testing.T) {
	const defName = "KlangTest"

	// var sig = Pan2.ar(
	//         Klang.ar(`[ Array.rand(12, 200.0, 2000.0), nil, nil ], 1, 0),
	// 	1.0.rand
	// );
	// Out.ar(0, sig * EnvGen.kr(Env.sine(4), 1, 0.02, doneAction: 2);
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		freqs := []Input{
			C(561.384644),
			C(1043.168701),
			C(237.107315),
			C(303.264008),
			C(927.150208),
			C(833.526123),
			C(509.927826),
			C(946.380005),
			C(752.409973),
			C(525.558716),
			C(1111.182129),
			C(715.820068),
		}
		sig := A(Pan2{
			In: A(Klang{
				Spec: ArraySpec{freqs, nil, nil},
			}),
			Pos: C(0.307131),
		})
		env := K(EnvGen{
			Env:        EnvSine{Dur: C(4)},
			Gate:       C(1),
			LevelScale: C(0.02),
			Done:       FreeEnclosing,
		})
		return Out{
			Bus:      C(0),
			Channels: sig.Mul(env),
		}.Rate(AR)
	}))
}
