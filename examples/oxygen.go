package main

import (
	"encoding/json"
	"fmt"
	"github.com/rakyll/portmidi"
	"github.com/scgolang/sc"
	. "github.com/scgolang/sc/types"
	. "github.com/scgolang/sc/ugens"
	"os"
)

func main() {
	var synthId int32
	const synthName = "sineTone"
	// setup supercollider client
	client := sc.NewClient("127.0.0.1", 57121)
	err := client.Connect("127.0.0.1", 57120)
	if err != nil {
		panic(err)
	}
	def := sc.NewSynthdef(synthName, func(p Params) Ugen {
		bus, env := C(0), EnvGen{Env: EnvPerc{}, Done: FreeEnclosing}.Rate(KR)
		sig := SinOsc{}.Rate(AR).Mul(env)
		return Out{bus, sig}.Rate(AR)
	})
	err = client.SendDef(def)
	if err != nil {
		panic(err)
	}

	// initialize midi
	portmidi.Initialize()
	// this code can be uncommented to discover the
	// device ID's portmidi comes up with
	// deviceCount := portmidi.CountDevices()
	// enc := json.NewEncoder(os.Stdout)
	// for i := 0; i < deviceCount; i++ {
	// 	info := portmidi.GetDeviceInfo(portmidi.DeviceId(i))
	// 	fmt.Printf("device %d - ", i)
	// 	err = enc.Encode(info)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// setup midi input stream and listen for midi events
	in, err := portmidi.NewInputStream(3, 1024)
	if err != nil {
		panic(err)
	}
	ch := in.Listen()
	for event := range ch {
		if event.Status == 144 {
			// MIDI note
			// fmt.Printf("Note %-3d Velocity %-3d\n", event.Data1, event.Data2)
			if event.Data2 > 0 {
				// Note On
				synthId = client.NextSynthId()
				err = client.Synth(synthName)
			}
		}
	}
	portmidi.Terminate()
}
