package sc

import (
	"fmt"

	"github.com/scgolang/osc"
)

// ServerStatus represents the reply to the /status command.
type ServerStatus struct {
	NumUgens          int32   `json:"numUgens"`
	NumSynths         int32   `json:"numSynths"`
	NumGroups         int32   `json:"numGroups"`
	NumSynthdefs      int32   `json:"numSynthdefs"`
	AvgCPU            float32 `json:"avgCPU"`
	PeakCPU           float32 `json:"peakCPU"`
	NominalSampleRate float32 `json:"nominalSampleRate"`
	ActualSampleRate  float32 `json:"actualSampleRate"`
}

func newStatus(msg osc.Message) (*ServerStatus, error) {
	if msg.Address != statusReplyAddress {
		errmsg := "Can not get status from message with address %s"
		return nil, fmt.Errorf(errmsg, msg.Address)
	}
	numArgs := len(msg.Arguments)
	status := &ServerStatus{}
	if numArgs != 9 {
		return nil, fmt.Errorf("Only got %d arguments in /status.reply message", numArgs)
	}
	var err error
	status.NumUgens, err = msg.Arguments[0].ReadInt32()
	if err != nil {
		return nil, err
	}
	status.NumSynths, err = msg.Arguments[1].ReadInt32()
	if err != nil {
		return nil, err
	}
	status.NumGroups, err = msg.Arguments[2].ReadInt32()
	if err != nil {
		return nil, err
	}
	status.NumSynthdefs, err = msg.Arguments[3].ReadInt32()
	if err != nil {
		return nil, err
	}
	status.AvgCPU, err = msg.Arguments[4].ReadFloat32()
	if err != nil {
		return nil, err
	}
	status.PeakCPU, err = msg.Arguments[5].ReadFloat32()
	if err != nil {
		return nil, err
	}
	status.NominalSampleRate, err = msg.Arguments[6].ReadFloat32()
	if err != nil {
		return nil, err
	}
	status.ActualSampleRate, err = msg.Arguments[7].ReadFloat32()
	if err != nil {
		return nil, err
	}
	return status, nil
}
