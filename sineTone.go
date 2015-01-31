package gosc

func SineTone() SynthDef {
	return SynthDef{
		Name: NewPstring("SineTone"),
		NumConstants: int32(2),
		Constants: make([]float32, 0),
		NumParams: int32(0),
		InitialParamValues: make([]float32, 0),
		NumParamNames: int32(0),
		ParamNames: make([]ParamName,  0),
		NumUgens: int32(2),
		Ugens: make([]Ugen, 0),
	}
}
