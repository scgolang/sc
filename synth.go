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
	DefName string `json:"defName"`
	ID      int32  `json:"id"`
	client  *Client
}

// Get the value of a synth control.
func (s *Synth) Get(controlName string) (float32, error) {
	return 0, nil
}

// Set the value of a synth control.
func (s *Synth) Set(ctls map[string]float32) error {
	set, err := osc.NewMessage(setSynthNodeAddress)
	if err != nil {
		return err
	}
	if err := set.WriteInt32(0, s.ID); err != nil {
		return err
	}
	argidx := 1
	for name, value := range ctls {
		if err := set.WriteString(argidx, name); err != nil {
			return err
		}
		if err := set.WriteFloat32(argidx, value); err != nil {
			return err
		}
	}
	_, err = s.client.oscConn.Send(set)
	return err
}

// Free a synth node.
func (s *Synth) Free() error {
	free, err := osc.NewMessage(freeSynthNodeAddress)
	if err != nil {
		return err
	}
	if err := free.WriteInt32(0, s.ID); err != nil {
		return err
	}
	_, err = s.client.oscConn.Send(free)
	return err
}

// newSynth creates a new synth structure.
func newSynth(client *Client, defName string, id int32) *Synth {
	return &Synth{
		DefName: defName,
		ID:      id,
		client:  client,
	}
}
