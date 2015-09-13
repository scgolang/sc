package sc

import (
	"github.com/scgolang/osc"
)

const (
	freeSynthNodeAddress = "/n_free"
	setSynthNodeAddress  = "/n_set"
)

// Synth encapsulates a synth node.
type Synth struct {
	defName string `json:"defName"`
	id      int32  `json:"id"`
	client  *Client
}

// Get the value of a synth control.
func (self *Synth) Get(controlName string) (float32, error) {
	return 0, nil
}

// Set the value of a synth control.
func (self *Synth) Set(ctls map[string]float32) error {
	set := osc.NewMessage(setSynthNodeAddress)
	for name, value := range ctls {
		set.Append(name, value)
	}
	return self.client.oscServer.SendTo(self.client.conn, set)
}

// Free a synth node.
func (self *Synth) Free() error {
	free := osc.NewMessage(freeSynthNodeAddress)
	free.Append(self.id)
	return self.client.oscServer.SendTo(self.client.conn, free)
}

// newSynth creates a new synth structure.
func newSynth(client *Client, defName string, id int32) *Synth {
	return &Synth{
		defName: defName,
		id:      id,
		client:  client,
	}
}
