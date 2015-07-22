package sc

import (
	. "github.com/scgolang/sc/types"
	"sync"
	"sync/atomic"
)

// buffer is an implementation of the Buffer interface
type buffer struct {
	num int32
}

// Num Buffer implemenation
func (self *buffer) Num() int32 {
	return self.num
}

// global buffer index
// seems smelly for the client to have to keep track of this:
// what if multiple clients want to use the same scsynth instance?
var bufIndex int32 = 0

// global buffer map
// keys are paths to audio files on disk
var buffers = struct {
	sync.RWMutex
	m map[string]*buffer
}{m: make(map[string]*buffer)}

// newBuffer creates a new buffer
func newBuffer(path string) Buffer {
	// return the existing buffer if there is one
	buffers.RLock()
	if eb, exists := buffers.m[path]; exists {
		buffers.RUnlock()
		return eb
	}
	buffers.RUnlock()
	// make a new one
	b := buffer{bufIndex}
	// add it to the global map
	buffers.Lock()
	buffers.m[path] = &b
	buffers.Unlock()
	// increment the global buffer index and
	// return the new buffer
	atomic.AddInt32(&bufIndex, int32(1))
	return &b
}
