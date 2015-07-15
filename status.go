package sc

import (
	"fmt"
	"github.com/scgolang/osc"
)

type ServerStatus struct {
	NumUgens          int32   `json:"numUgens"`
	NumSynths         int32   `json:"numSynths"`
	NumGroups         int32   `json:"numGroups"`
	NumSynthdefs      int32   `json:"numSynthdefs"`
	AvgCpu            float32 `json:"avgCpu"`
	PeakCpu           float32 `json:"peakCpu"`
	NominalSampleRate float64 `json:"nominalSampleRate"`
	ActualSampleRate  float64 `json:"actualSampleRate"`
}

func newStatus(msg *osc.Message) (*ServerStatus, error) {
	if msg.Address != statusReplyAddress {
		errmsg := "Can not get status from message with address %s"
		return nil, fmt.Errorf(errmsg, msg.Address)
	}
	numArgs := msg.CountArguments()
	status := new(ServerStatus)
	if numArgs != 9 || len(msg.Arguments) != 9 {
		errmsg := "Only got %d arguments in /status.reply message"
		return nil, fmt.Errorf(errmsg, numArgs)
	}
	status.NumUgens = msg.Arguments[1].(int32)
	status.NumSynths = msg.Arguments[2].(int32)
	status.NumGroups = msg.Arguments[3].(int32)
	status.NumSynthdefs = msg.Arguments[4].(int32)
	status.AvgCpu = msg.Arguments[5].(float32)
	status.PeakCpu = msg.Arguments[6].(float32)
	status.NominalSampleRate = msg.Arguments[7].(float64)
	status.ActualSampleRate = msg.Arguments[8].(float64)
	return status, nil
}
