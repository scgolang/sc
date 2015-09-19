package sc

import "testing"

func TestPlayBuf(t *testing.T) {
	def := NewSynthdef("PlayBufExample", func(p Params) Ugen {
		// sclang synthdef:
		// arg bufnum = 0;
		// Out.ar(0, PlayBuf.ar(1, bufnum, 1.0, 1.0, 0, 0, 2));
		bufnum := p.Add("bufnum", 0)
		bus, channels := C(0), 1
		sig := PlayBuf{
			NumChannels: channels,
			BufNum:      bufnum,
			Done:        FreeEnclosing,
		}.Rate(AR)
		return Out{bus, sig}.Rate(AR)
	})
	compareAndWrite(t, "PlayBufExample", def)
}
