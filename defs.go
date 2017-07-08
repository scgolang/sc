package sc

import (
	"errors"
	"sync"
)

// Synthdefs is a map of synthdefs.
var Synthdefs = map[string]*Synthdef{}

var synthdefsMu sync.RWMutex

// RegisterSynthdef registers a synthdef with this package.
// It returns an error if a synthdef is already registered with the provided name.
func RegisterSynthdef(name string, f UgenFunc) error {
	synthdefsMu.Lock()
	defer synthdefsMu.Unlock()
	if _, ok := Synthdefs[name]; ok {
		return errors.New("synthdef already registered: " + name)
	}
	Synthdefs[name] = NewSynthdef(name, f)
	return nil
}
