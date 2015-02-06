package sc

import (
	"fmt"
)

type Input interface {
	IsConstant() bool
	ConstantValue() float32
	UgenValue() *Ugen
}

type constantInput struct {
	value float32
}

func (self *constantInput) IsConstant() bool {
	return true
}

func (self *constantInput) ConstantValue() float32 {
	return self.value
}

func (self *constantInput) UgenValue() *Ugen {
	panic(fmt.Errorf("can not convert ugen to constant"))
}

type ugenInput struct {
	value *Ugen
}

func (self *ugenInput) IsConstant() bool {
	return false
}

func (self *ugenInput) ConstantValue() float32 {
	panic(fmt.Errorf("can not convert constant to ugen"))
}

func (self *ugenInput) UgenValue() *Ugen {
	return self.value
}

func ConstantInput(value float32) Input {
	return &constantInput{value}
}

func UgenInput(value *Ugen) Input {
	return &ugenInput{value}
}
