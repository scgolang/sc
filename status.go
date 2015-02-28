package sc

import (
	"fmt"
	"github.com/briansorahan/go-osc/osc"
)

const (
	StatusOscAddress = "/status.reply"
)

type ServerStatus struct {
	NumUgens          int32
	NumSynths         int32
	NumGroups         int32
	NumSynthdefs      int32
	AvgCpu            float32
	PeakCpu           float32
	NominalSampleRate float64
	ActualSampleRate  float64
}

func newStatus(msg *osc.OscMessage) (*ServerStatus, error) {
	if msg.Address != StatusOscAddress {
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
