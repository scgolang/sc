package synthdefs

import (
	"fmt"
	"github.com/briansorahan/gosc"
)

var Out = NewMetaOut()

type MetaOut struct {
	index int32
	channels []gosc.Synth
}

func (self *MetaOut) Ar(args ...interface{}) (*MetaOut, error) {
	var index int32

	channels := make([]gosc.Synth, 0)

	if len(args) < 2 {
		return nil, fmt.Errorf("not enough arguments to Out.Ar")
	}
	if index, ok := args[0].(int32); !ok {
		return nil, fmt.Errorf("first argument to Out.Ar must be an int32")
	}
	// just a single synth?
	if channels, ok := args[1].(gosc.Synth); !ok {
		// an array of synths?
		if channels, ok := args[1].([]gosc.Synth); !ok {
			return nil, fmt.Errorf("second argument to Out.Ar must be a synth or an array of synths")
		}
	}
	return self, nil
}

func (self *MetaOut) Kr(args ...interface{}) (*MetaOut, error) {
	return self, nil
}

func NewMetaOut() (MetaOut, error) {
}
