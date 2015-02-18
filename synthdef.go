package sc

import (
	. "github.com/briansorahan/sc/types"
	"io"
)

type synthdef struct {
	name string
	constants []float32
	initialParamValues []float32
	paramNames []string
	ugens []Ugen
	variants []Variant
}

func (self *synthdef) Name() string {
	return self.name
}

func (self *synthdef) Constants() []float32 {
	return self.constants
}

func (self *synthdef) InitialParamValues() []float32 {
	return self.initialParamValues
}

func (self *synthdef) ParamNames() []string {
	return self.paramNames
}

func (self *synthdef) Ugens() []Ugen {
	return self.ugens
}

func (self *synthdef) Variants() []Variant {
	return self.variants
}

func (self *synthdef) Print(w io.Writer) {
}

// NewSynthdef creates a new synthdef
func NewSynthdef(name string, f UgenGraphFunc) Synthdef {
	sd := synthdef{
		name,
		make([]float32, 0),
		make([]float32, 0),
		make([]string, 0),
		make([]Ugen, 0),
		make([]Variant, 0),
	}
	return &sd
}
