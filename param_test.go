package sc

// import (
// 	"testing"
// )

// func Test2Params(t *testing.T) {
// 	synthName := "defWith2Params"
// 	def := NewSynthdef(synthName, func(p Params) Ugen {
// 		freq := p.Add("freq", 440)
// 		gain := p.Add("gain", 0.5)
// 		bus := C(0)
// 		env := EnvGen{
// 			Env:        EnvPerc{},
// 			Done:       FreeEnclosing,
// 			LevelScale: gain,
// 		}.Rate(KR)
// 		sig := SinOsc{Freq: freq}.Rate(AR).Mul(env)
// 		return Out{bus, sig}.Rate(AR)
// 	})
// 	compareAndWrite(t, synthName, def)
// 	sclangVersion := `{
// 	    arg freq=440, gain=0.5;
// 	    var env = EnvGen.kr(Env.perc, doneAction: 2, levelScale: gain);
// 	    var sine = SinOsc.ar(freq);
// 	    Out.ar(0, sine * env);
// 	}`
// 	same, err := def.Compare(sclangVersion)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if !same {
// 		f, err := os.Open(synthName + ".gosyndef")
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		def.Write
// 		t.Fatalf("synthdef different from sclang version")
// 	}
// }
