package sc

import (
	"fmt"
)

// Ugen
type Ugen struct {
	name   string
	rate   int8
	inputs []Input
}

// AppendConstant appends a constant to the list of inputs
func (self *Ugen) AppendConstant(value float32) {
	self.inputs = append(self.inputs, ConstantInput(value))
}

// AppendUgen appends a ugen to the list of inputs
func (self *Ugen) AppendUgen(value *Ugen) {
	self.inputs = append(self.inputs, UgenInput(value))
}

// Ar creates a new audio-rate ugen
func Ar(name string, args ...interface{}) (*Ugen, error) {
	return newUgen(name, 2, args...)
}

// Kr creates a new control-rate ugen
func Kr(name string, args ...interface{}) (*Ugen, error) {
	return newUgen(name, 1, args...)
}

// Ir creates a new initialization-rate ugen
func Ir(name string, args ...interface{}) (*Ugen, error) {
	return newUgen(name, 0, args...)
}

func newUgen(name string, rate int8, args ...interface{}) (*Ugen, error) {
	u := Ugen{
		name,             // name
		rate,             // rate
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
