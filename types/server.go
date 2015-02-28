package types

type Server interface {
	Start() error
	Status() (ServerStatus, error)
	Close() error
}

type ServerStatus interface {
	NumUgens()          int32
	NumSynths()         int32
	NumGroups()         int32
	NumSynthdefs()      int32
	AvgCpu()            float32
	PeakCpu()           float32
	NominalSampleRate() float64
	ActualSampleRate()  float64
}
