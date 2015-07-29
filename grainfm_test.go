package sc

import (
	. "github.com/scgolang/sc/types"
	. "github.com/scgolang/sc/ugens"
	"testing"
)

func TestGrainFM(t *testing.T) {
	defName := "GrainFMTest"
	def := NewSynthdef(defName, func(p Params) Ugen {
		// arg gate=1, amp=1;
		// var pan = MouseX.kr(-1, 1);
		// var freqdev = WhiteNoise.kr(MouseY.kr(0, 400));
		// var env = Env([0, 1, 0], [1, 1], \sin, 1);
		// var ampenv = EnvGen.kr(env, gate, levelScale: amp, doneAction: 2);
		// var trig = Impulse.kr(10);
		// var modIndex = LFNoise1.kr.range(1, 10);
		// var sig = GrainFM.ar(
		//     numChannels: 2,
		//     trigger: trig,
		//     dur: 0.1,
		//     carFreq: 440+freqdev,
		//     modFreq: 200,
		//     index: modIndex,
		//     pan: pan
		// );
		// Out.ar(0, sig * ampenv);
		gate := p.Add("gate", 1)
		amp := p.Add("amp", 1)
		bus := C(0)
		mousey := MouseY{Min: C(0), Max: C(400)}.Rate(KR)
		freqdev := WhiteNoise{}.Rate(KR).Mul(mousey)
		env := Env{
			Levels:     []Input{C(0), C(1), C(0)},
			Times:      []Input{C(1), C(1)},
			CurveTypes: []Input{CurveSine, CurveSine},
			// Curvature:   CurveSine,
			ReleaseNode: C(1),
		}
		ampenv := EnvGen{
			Env:        env,
			Gate:       gate,
			LevelScale: amp,
			Done:       FreeEnclosing,
		}.Rate(KR)
		trig := Impulse{Freq: C(10)}.Rate(KR)
		modIndex := LFNoise1{}.Rate(KR).MulAdd(C(5), C(5))
		pan := MouseX{Min: C(-1), Max: C(1)}.Rate(KR)
		sig := GrainFM{
			NumChannels: 2,
			Trigger:     trig,
			Dur:         C(0.1),
			CarFreq:     C(440).Add(freqdev),
			ModFreq:     C(200),
			ModIndex:    modIndex,
			Pan:         pan,
		}.Rate(AR)
		return Out{bus, sig.Mul(ampenv)}.Rate(AR)
	})
	compareAndWrite(t, defName, def)
}
