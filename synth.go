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
	set, err := osc.NewMessage(setSynthNodeAddress)
	if err != nil {
		return err
	}
	if err := set.WriteInt32(self.id); err != nil {
		return err
	}
	for name, value := range ctls {
		if err := set.WriteString(name); err != nil {
			return err
		}
		if err := set.WriteFloat32(value); err != nil {
			return err
		}
	}
	return self.client.oscConn.Send(set)
}

// Free a synth node.
func (self *Synth) Free() error {
	free, err := osc.NewMessage(freeSynthNodeAddress)
	if err != nil {
		return err
	}
	if err := free.WriteInt32(self.id); err != nil {
		return err
	}
	return self.client.oscConn.Send(free)
}

// newSynth creates a new synth structure.
func newSynth(client *Client, defName string, id int32) *Synth {
	return &Synth{
		defName: defName,
		id:      id,
		client:  client,
	}
}
