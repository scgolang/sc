package sc

import (
	"errors"
	"fmt"
	"sync"

	"github.com/scgolang/osc"
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

// Buffer is a client-side representation of an scsynth audio buffer
type Buffer struct {
	Num    int32
	client *Client
}

// Gen generates a buffer using a routine.
// A runtime panic will occur if routine is not one of the
// BufferRoutine constants.
func (buffer *Buffer) Gen(routine string, flags int, args ...float32) error {
	if err := checkBufferRoutine(routine); err != nil {
		return err
	}
	if err := checkBufferGenFlags(flags); err != nil {
		return err
	}
	if err := buffer.sendGenMsg(routine, flags, args...); err != nil {
		return err
	}
	if err := buffer.awaitGenReply(); err != nil {
		return err
	}
	return nil
}

// sendGenMsg sends a /b_gen command.
func (buffer *Buffer) sendGenMsg(routine string, flags int, args ...float32) error {
	msg := osc.Message{
		Address: bufferGenAddress,
		Arguments: osc.Arguments{
			osc.Int(buffer.Num),
			osc.String(routine),
			osc.Int(int32(flags)),
		},
	}
	for _, arg := range args {
		msg.Arguments = append(msg.Arguments, osc.Float(arg))
	}
	if err := buffer.client.oscConn.Send(msg); err != nil {
		return err
	}
	return nil
}

// awaitGenReply waits for a reply to the /b_gen command
func (buffer *Buffer) awaitGenReply() error {
	var done osc.Message
	select {
	case done = <-buffer.client.doneChan:
	case err := <-buffer.client.errChan:
		return err
	}
	if len(done.Arguments) != 2 {
		return errors.New("expected two arguments to /done message")
	}
	addr, err := done.Arguments[0].ReadString()
	if err != nil {
		return err
	}
	// If reply address is not /b_gen, requeue the done event.
	if addr != bufferGenAddress {
		buffer.client.doneChan <- done
		return nil
	}
	bufnum, err := done.Arguments[1].ReadInt32()
	if err != nil {
		return err
	}

	// TODO: Don't error if we get a done message for a different buffer.
	// We should probably requeue this particular done message on doneChan.
	if bufnum != buffer.Num {
		buffer.client.doneChan <- done
	}
	return nil
}

// checkBufferRoutine panics if routine is not one of the
// supported BufferRoutine constants
func checkBufferRoutine(routine string) error {
	if routine != BufferRoutineSine1 &&
		routine != BufferRoutineSine2 &&
		routine != BufferRoutineSine3 &&
		routine != BufferRoutineCheby {
		return fmt.Errorf("unsupported buffer routine %s", routine)
	}
	return nil
}

// checkBufferGenFlags panics if not 0 <= flags <= 4
func checkBufferGenFlags(flags int) error {
	if flags < 0 && flags > 4 {
		return fmt.Errorf("unsupported buffer flags %d", flags)
	}
	return nil
}

// global buffer map (keys are paths to audio files on disk)
var buffers = struct {
	sync.RWMutex
	m map[string]*Buffer
}{m: make(map[string]*Buffer)}

// newReadBuffer creates a new buffer for /b_allocRead
func newReadBuffer(path string, num int32, c *Client) *Buffer {
	buffers.RLock()
	// return the existing buffer if there is one
	if existing, exists := buffers.m[path]; exists {
		buffers.RUnlock()
		return existing
	}
	buffers.RUnlock()

	// make a new one
	b := &Buffer{Num: num, client: c}

	// add it to the global map
	buffers.Lock()
	buffers.m[path] = b
	buffers.Unlock()

	return b
}
