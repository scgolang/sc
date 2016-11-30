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
	msg := osc.Message{
		Address: setSynthNodeAddress,
		Arguments: osc.Arguments{
			osc.Int(s.ID),
		},
	}
	for name, value := range ctls {
		msg.Arguments = append(msg.Arguments, osc.String(name))
		msg.Arguments = append(msg.Arguments, osc.Float(value))
	}
	return s.client.oscConn.Send(msg)
}

// newSynth creates a new synth structure.
func newSynth(client *Client, defName string, id int32) *Synth {
	return &Synth{
		DefName: defName,
		ID:      id,
		client:  client,
	}
}
