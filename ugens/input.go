package ugens

import (
	"github.com/briansorahan/sc"
)

type constantInput float32

func (self constantInput) IsConstant() bool {
	return true
}

func (self constantInput) Value() interface{} {
	return self
}

func newConstantInput(val float32) sc.Input {
	return constantInput(val)
}
