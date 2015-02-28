package sc

import (
	"fmt"
	"github.com/briansorahan/go-osc/osc"
)

const (
	StatusOscAddress = "/status.reply"
)

type serverStatus struct {
	numUgens          int32
	numSynths         int32
	numGroups         int32
	numSynthdefs      int32
	avgCpu            float32
	peakCpu           float32
	nominalSampleRate float64
	actualSampleRate  float64
}

func (self *serverStatus) NumUgens() int32 {
	return self.numUgens
}

func (self *serverStatus) NumSynths() int32 {
	return self.numUgens
}

func (self *serverStatus) NumGroups() int32 {
	return self.numUgens
}

func (self *serverStatus) NumSynthdefs() int32 {
	return self.numUgens
}

func (self *serverStatus) AvgCpu() float32 {
	return self.avgCpu
}

func (self *serverStatus) PeakCpu() float32 {
	return self.avgCpu
}

func (self *serverStatus) NominalSampleRate() float64 {
	return self.nominalSampleRate
}

func (self *serverStatus) ActualSampleRate() float64 {
	return self.nominalSampleRate
}

func newStatus(msg *osc.OscMessage) (*serverStatus, error) {
	if msg.Address != StatusOscAddress {
		errmsg := "Can not get status from message with address %s"
		return nil, fmt.Errorf(errmsg, msg.Address)
	}
	numArgs := msg.CountArguments()
	status := new(serverStatus)
	if numArgs != 9 || len(msg.Arguments) != 9 {
		errmsg := "Only got %d arguments in /status.reply message"
		return nil, fmt.Errorf(errmsg, numArgs)
	}
	status.numUgens = msg.Arguments[1].(int32)
	status.numSynths = msg.Arguments[2].(int32)
	status.numGroups = msg.Arguments[3].(int32)
	status.numSynthdefs = msg.Arguments[4].(int32)
	status.avgCpu = msg.Arguments[5].(float32)
	status.peakCpu = msg.Arguments[6].(float32)
	status.nominalSampleRate = msg.Arguments[7].(float64)
	status.actualSampleRate = msg.Arguments[8].(float64)
	return status, nil
}
