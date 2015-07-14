package sc

import (
	"github.com/scgolang/osc"
)

// Synth represents a synth node on the server
type Synth interface {
	// Get fetches the value of a control
	Get(controlName string) (float32, error)
	// Set sets the value of a control
	Set(controlName string, val float32) error
	// Free frees the synth node
	Free() error
}

type synth struct {
	defName string `json:"defName"`
	id      int32  `json:"id"`
	client  *Client
}

func (self *synth) Get(controlName string) (float32, error) {
	return 0, nil
}

func (self *synth) Set(controlName string, val float32) error {
	return nil
}

func (self *synth) Free() error {
	free := osc.NewMessage("/n_free")
	free.Append(self.id)
	return self.client.oscServer.SendTo(self.client.conn, free)
}

func newSynth(client *Client, defName string, id int32) Synth {
	return &synth{
		defName: defName,
		id:      id,
		client:  client,
	}
}
