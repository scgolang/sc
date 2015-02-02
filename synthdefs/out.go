package synthdefs

import (
	"fmt"
	"github.com/briansorahan/gosc"
)

var Out = NewMetaOut()

type MetaOut struct {}

func (self *MetaOut) Ar(args ...interface{}) (*gosc.Node, error) {
	root := gosc.NewNode(nil)

	if len(args) < 2 {
		return nil, fmt.Errorf("not enough arguments to Out.Ar")
	}
	// get the index
	if index, ok := args[0].(int32); !ok {
		return nil, fmt.Errorf("first argument to Out.Ar must be an int32")
	} else {
		root.Add(gosc.NewNode(index))
	}

	// just a single synth?
	if channels, ok := args[1].(gosc.Synth); !ok {
		// an array of synths?
		if channels, ok := args[1].([]gosc.Synth); !ok {
			return nil, fmt.Errorf("second argument to Out.Ar must be a synth or an array of synths")
		} else {
			if len(channels) == 0 {
				return nil, fmt.Errorf("Out requires at least 1 synth argument")
			}
			root.Add(gosc.NewNode(channels))
		}
	} else {
		root.Add(gosc.NewNode(channels))
	}

	return root, nil
}

func (self *MetaOut) Kr(args ...interface{}) (*gosc.Node, error) {
	return nil, nil
}

func NewMetaOut() *MetaOut {
	return new(MetaOut)
}
