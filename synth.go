package sc

import (
	"github.com/scgolang/osc"
)

type Synth struct {
	defName string `json:"defName"`
	id      int32  `json:"id"`
	client  *Client
}

func (self *Synth) Get(controlName string) (float32, error) {
	return 0, nil
}

func (self *Synth) Set(controlName string, val float32) error {
	return nil
}

func (self *Synth) Free() error {
	free := osc.NewMessage("/n_free")
	free.Append(self.id)
	return self.client.oscServer.SendTo(self.client.conn, free)
}

func newSynth(client *Client, defName string, id int32) *Synth {
	return &Synth{
		defName: defName,
		id:      id,
		client:  client,
	}
}
