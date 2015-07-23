package sc

import (
	"fmt"
	"github.com/scgolang/osc"
	. "github.com/scgolang/sc/types"
	"reflect"
	"sync"
	"sync/atomic"
)

const (
	bufIndexIncr = int32(1)
	// BufferFlagNormalize causes the peak amplitude of the buffer
	// to be normalized to 1.0 when using the Gen method.
	BufferFlagNormalize = 0x01
	// BufferFlagWavetable causes the buffer to be written in wavetable
	// format when using the Gen method, so that it can be used by
	// interpolating oscillators
	BufferFlagWavetable = 0x02
	// BufferFlagClear causes the buffer to be cleared before new
	// partials are written using the Gen method. Otherwise new partials
	// are summed with the existing contents of the buffer.
	BufferFlagClear = 0x04
	// BufferRoutineSine1 generates a buffer with a series of sine partials.
	// The args of this routine are the amplitudes of the partials.
	BufferRoutineSine1 = "sine1"
	// BufferRoutineSine2 is similar to BufferRoutineSine1 except that the
	// args are pairs of frequency and amplitude (i.e. you can specify the
	// frequency of each partial).
	BufferRoutineSine2 = "sine2"
	// BufferRoutineSine3 is similar to BufferRoutineSine1 except that the
	// args are triplets of frequency, amplitude, and phase (i.e. you can
	// specify the frequency and phase of each partial).
	BufferRoutineSine3 = "sine3"
	// BufferRoutineCheby generates a buffer that contains a series of
	// chebyshev polynomials which can be defined as
	//     cheby(n) = amplitude * cos(n * acos(x))
	// The first arg specifies the amplitude for n = 1, the second arg
	// specifies the amplitude for n = 2, and so on.
	// To eliminate DC offset when used as a waveshaper, the wavetable is
	// offset so that the center value is zero.
	BufferRoutineCheby = "cheby"
)

// buffer is an implementation of the Buffer interface
type buffer struct {
	num int32
	c   *Client
}

// Num returns the buffer number
func (self *buffer) Num() int32 {
	return self.num
}

// Gen generates a buffer using a routine.
// A runtime panic will occur if routine is not one of the
// BufferRoutine constants.
func (self *buffer) Gen(routine string, flags int, args ...float32) error {
	checkBufferRoutine(routine)
	checkBufferGenFlags(flags)
	pat := bufferGenAddress
	gen := osc.NewMessage(pat)
	gen.Append(self.Num())
	gen.Append(routine)
	gen.Append(int32(flags))
	for _, arg := range args {
		gen.Append(arg)
	}
	err := self.c.oscServer.SendTo(self.c.conn, gen)
	if err != nil {
		return err
	}

	var done *osc.Message
	select {
	case done = <-self.c.doneChan:
		break
	case err = <-self.c.oscErrChan:
		return err
	}

	if done.CountArguments() != 2 {
		return fmt.Errorf("expected two arguments to /done message")
	}
	if addr, isString := done.Arguments[0].(string); !isString || addr != pat {
		return fmt.Errorf("expected first argument to be %s but got %s", pat, addr)
	}
	var bufnum int32
	var isInt32 bool
	if bufnum, isInt32 = done.Arguments[1].(int32); !isInt32 {
		m := "expected int32 as second argument, but got %s (%v)"
		return fmt.Errorf(m, reflect.TypeOf(done.Arguments[1]), done.Arguments[1])
	}
	// TODO:
	// Don't error if we get a done message for a different buffer.
	// We should probably requeue this particular done message on doneChan.
	if bufnum != self.Num() {
		m := "expected done message for buffer %d, but got one for buffer %d"
		return fmt.Errorf(m, self.Num(), bufnum)
	}

	return nil
}

// checkBufferRoutine panics if routine is not one of the
// supported BufferRoutine constants
func checkBufferRoutine(routine string) {
	if routine != BufferRoutineSine1 &&
		routine != BufferRoutineSine2 &&
		routine != BufferRoutineSine3 &&
		routine != BufferRoutineCheby {
		panic(fmt.Errorf("unsupported buffer routine %s", routine))
	}
}

// checkBufferGenFlags panics if not 0 <= flags <= 4
func checkBufferGenFlags(flags int) {
	if flags < 0 && flags > 4 {
		panic(fmt.Errorf("unsupported buffer flags %s", flags))
	}
}

// global buffer index
var bufIndex int32 = -1

// global buffer map (keys are paths to audio files on disk)
var buffers = struct {
	sync.RWMutex
	m map[string]*buffer
}{m: make(map[string]*buffer)}

// newReadBuffer creates a new buffer for /b_allocRead
func newReadBuffer(path string, c *Client) Buffer {
	// return the existing buffer if there is one
	buffers.RLock()
	if eb, exists := buffers.m[path]; exists {
		buffers.RUnlock()
		return eb
	}
	buffers.RUnlock()
	// make a new one
	b := buffer{atomic.AddInt32(&bufIndex, bufIndexIncr), c}
	// add it to the global map
	buffers.Lock()
	buffers.m[path] = &b
	buffers.Unlock()
	return &b
}

// newBuffer creates a new buffer for /b_alloc
func newBuffer(c *Client) Buffer {
	return &buffer{atomic.AddInt32(&bufIndex, bufIndexIncr), c}
}
