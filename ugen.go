package sc

import (
	"fmt"
)

// Ugen
type Ugen struct {
	Name   string
	Rate   int8
	Inputs []Input
}

func (self *Ugen) AppendConstant(value float32) {
	self.Inputs = append(self.Inputs, ConstantInput(value))
}

func (self *Ugen) AppendUgen(value *Ugen) {
	self.Inputs = append(self.Inputs, UgenInput(value))
}

func Ar(name string, args ...interface{}) (*Ugen, error) {
	u := Ugen{
		name,             // name
		2,                // rate
		make([]Input, 0), // inputs
	}

	for _, arg := range args {
		if fv, isFloat := arg.(float32); isFloat {
			u.AppendConstant(fv)
		} else if ug, isUgen := arg.(*Ugen); isUgen {
			u.AppendUgen(ug)
		} else {
			return nil, fmt.Errorf("ugen arguments must be float32's or ugen's")
		}
	}

	return &u, nil
}
