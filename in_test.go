package sc

import (
	"os"
	"testing"
)

func TestIn(t *testing.T) {
	const name = "InTest"

	def := NewSynthdef(name, func(params Params) Ugen {
		var (
			out = params.Add("out", 0)
			in  = params.Add("in", 0)
		)
		return Out{
			Bus: out,
			Channels: SinOsc{
				Freq: In{
					Bus:         in,
					NumChannels: 2,
				}.Rate(KR),
			}.Rate(AR).Mul(C(0.1)),
		}.Rate(AR)
	})
	f, err := os.Create("testdata/" + name + ".gosyndef")
	if err != nil {
		t.Fatal(err)
	}
	if err := def.Write(f); err != nil {
		t.Fatal(err)
	}
	same, err := def.CompareToFile("testdata/" + name + ".scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang version")
	}
}
