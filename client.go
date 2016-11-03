package sc

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/scgolang/osc"
)

// OSC addresses.
// See http://doc.sccode.org/Reference/Server-Command-Reference.html.
const (
	statusAddress          = "/status"
	statusReplyAddress     = "/status.reply"
	gqueryTreeAddress      = "/g_queryTree"
	gqueryTreeReplyAddress = "/g_queryTree.reply"
	synthdefReceiveAddress = "/d_recv"
	dumpOscAddress         = "/dumpOSC"
	doneOscAddress         = "/done"
	synthNewAddress        = "/s_new"
	groupNewAddress        = "/g_new"
	groupFreeAllAddress    = "/g_freeAll"
	bufferAllocAddress     = "/b_alloc"
	bufferReadAddress      = "/b_allocRead"
	bufferGenAddress       = "/b_gen"
)

// Arguments to dumpOSC command.
// See http://doc.sccode.org/Reference/Server-Command-Reference.html#/dumpOSC
const (
	DumpOff      = 0
	DumpParsed   = 1
	DumpContents = 2
	DumpAll      = 3
)

// Arguments to s_new command.
// See http://doc.sccode.org/Reference/Server-Command-Reference.html#/s_new
const (
	AddToHead  = int32(0)
	AddToTail  = int32(1)
	AddBefore  = int32(2)
	AddAfter   = int32(3)
	AddReplace = int32(4)
)

const (
	// RootNodeID is what sclang uses as the root node ID. See http://doc.sccode.org/Classes/RootNode.html.
	RootNodeID = int32(0)

	// DefaultGroupID is what sclang uses for the default group ID. See http://doc.sccode.org/Reference/default_group.html.
	DefaultGroupID = int32(1)

	// DefaultLocalAddr is the listening address for DefaultClient.
	DefaultLocalAddr = "0.0.0.0:57110"

	// DefaultScsynthAddr is the remote address for DefaultClient.
	DefaultScsynthAddr = "0.0.0.0:57120"

	// DefaultConnectTimeout is the default timeout for connecting to scsynth.
	DefaultConnectTimeout = time.Second
)

// Common errors.
var (
	ErrTimeout = errors.New("timeout error")
)

// Client manages all communication with scsynth
type Client struct {
	// errChan is a channel that emits errors from
	// the goroutine that runs the OSC server that is
	// used to receive messages from scsynth
	errChan    chan error
	closeMutex sync.Mutex
	closed     int32

	addr    *net.UDPAddr
	oscConn osc.Conn

	doneChan       chan *osc.Message // doneChan relays /done messages
	statusChan     chan *osc.Message // statusChan relays /status.reply messages
	gqueryTreeChan chan *osc.Message // gqueryTreeChan relays /done messages

	nextSynthID int32 // next synth node ID
}

// number of concurrent handlers for /done messages.
const numDoneHandlers = 8

// NewClient creates a new SuperCollider client.
// The client will bind to the provided address and port
// to receive messages from scsynth.
func NewClient(network, local, scsynth string, timeout time.Duration) (*Client, error) {
	addr, err := net.ResolveUDPAddr(network, local)
	if err != nil {
		return nil, err
	}
	c := &Client{
		errChan:        make(chan error),
		statusChan:     make(chan *osc.Message),
		gqueryTreeChan: make(chan *osc.Message),
		doneChan:       make(chan *osc.Message, numDoneHandlers),
		addr:           addr,
		nextSynthID:    1000,
	}
	if err := c.Connect(scsynth, timeout); err != nil {
		return nil, err
	}
	return c, nil
}

var (
	defaultClient *Client
	defaultGroup  *Group
)

// DefaultClient returns the default sc client.
func DefaultClient() (*Client, error) {
	var err error

	if defaultClient == nil {
		defaultClient, err = NewClient("udp", DefaultLocalAddr, DefaultScsynthAddr, DefaultConnectTimeout)
		if err != nil {
			return nil, err
		}
		defaultGroup, err = defaultClient.AddDefaultGroup()
		if err != nil {
			return nil, err
		}
	}
	return defaultClient, nil
}

// Connect connects to an scsynth instance via UDP.
func (c *Client) Connect(addr string, timeout time.Duration) error {
	raddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}

	// Attempt connection with a timeout.
	var (
		start    = time.Now()
		timedOut = true
	)
	for time.Now().Sub(start) < timeout {
		oscConn, err := osc.DialUDP("udp", c.addr, raddr)
		if err != nil {
			time.Sleep(100 * time.Millisecond)
			continue
		}
		c.oscConn = oscConn
		timedOut = false
		break
	}
	if timedOut {
		return errors.New("connection timeout")
	}

	// listen for OSC messages
	go func(errChan chan error) {
		var (
			start = time.Now()
			err   error
		)
		for time.Now().Sub(start) < timeout {
			err = c.oscConn.Serve(c.oscHandlers())
			if err != nil {
				time.Sleep(100 * time.Second)
				continue
			}
		}
		if err != nil {
			errChan <- err
		}
	}(c.errChan)

	return nil
}

// Status gets the status of scsynth with a timeout.
// If the status request times out it returns ErrTimeout.
func (c *Client) Status(timeout time.Duration) (*ServerStatus, error) {
	statusReq, err := osc.NewMessage(statusAddress)
	if err != nil {
		return nil, err
	}
	if err := c.oscConn.Send(statusReq); err != nil {
		return nil, err
	}

	after := time.After(timeout)

	select {
	case _ = <-after:
		return nil, ErrTimeout
	case msg := <-c.statusChan:
		return newStatus(msg)
	case err = <-c.errChan:
		return nil, err
	}
}

// SendDef sends a synthdef to scsynth.
// This method blocks until a /done message is received
// indicating that the synthdef was loaded
func (c *Client) SendDef(def *Synthdef) error {
	msg, err := osc.NewMessage(synthdefReceiveAddress)
	if err != nil {
		return err
	}
	db, err := def.Bytes()
	if err != nil {
		return err
	}
	if err := msg.WriteBlob(db); err != nil {
		return err
	}
	if err := c.oscConn.Send(msg); err != nil {
		return err
	}
	var done *osc.Message
	select {
	case done = <-c.doneChan:
		goto ParseMessage
	case err = <-c.errChan:
		return err
	}

ParseMessage:
	// error if this message was not an ack of the synthdef
	errmsg := "expected /done with /d_recv argument"
	if done.CountArguments() != 1 {
		return fmt.Errorf(errmsg)
	}
	addr, err := done.ReadString()
	if err != nil {
		return err
	}
	if addr != synthdefReceiveAddress {
		return errors.New(errmsg)
	}
	return nil
}

// DumpOSC sends a /dumpOSC message to scsynth
// level should be DumpOff, DumpParsed, DumpContents, DumpAll
func (c *Client) DumpOSC(level int32) error {
	dumpReq, err := osc.NewMessage(dumpOscAddress)
	if err != nil {
		return err
	}
	if err := dumpReq.WriteInt32(level); err != nil {
		return err
	}
	return c.oscConn.Send(dumpReq)
}

// Synth creates a synth node.
func (c *Client) Synth(defName string, id, action, target int32, ctls map[string]float32) (*Synth, error) {
	synthReq, err := osc.NewMessage(synthNewAddress)
	if err != nil {
		return nil, err
	}
	if err := synthReq.WriteString(defName); err != nil {
		return nil, err
	}
	if err := synthReq.WriteInt32(id); err != nil {
		return nil, err
	}
	if err := synthReq.WriteInt32(action); err != nil {
		return nil, err
	}
	if err := synthReq.WriteInt32(target); err != nil {
		return nil, err
	}
	if ctls != nil {
		for k, v := range ctls {
			if err := synthReq.WriteString(k); err != nil {
				return nil, err
			}
			if err := synthReq.WriteFloat32(v); err != nil {
				return nil, err
			}
		}
	}
	if err := c.oscConn.Send(synthReq); err != nil {
		return nil, err
	}
	return newSynth(c, defName, id), nil
}

// Group creates a group.
func (c *Client) Group(id, action, target int32) (*Group, error) {
	dumpReq, err := osc.NewMessage(groupNewAddress)
	if err != nil {
		return nil, err
	}
	if err := dumpReq.WriteInt32(id); err != nil {
		return nil, err
	}
	if err := dumpReq.WriteInt32(action); err != nil {
		return nil, err
	}
	if err := dumpReq.WriteInt32(target); err != nil {
		return nil, err
	}
	if err := c.oscConn.Send(dumpReq); err != nil {
		return nil, err
	}
	return newGroup(c, id), nil
}

// AddDefaultGroup adds the default group.
func (c *Client) AddDefaultGroup() (*Group, error) {
	return c.Group(DefaultGroupID, AddToTail, RootNodeID)
}

// QueryGroup g_queryTree for a particular group.
func (c *Client) QueryGroup(id int32) (*Group, error) {
	addr := gqueryTreeAddress
	gq, err := osc.NewMessage(addr)
	if err != nil {
		return nil, err
	}
	if err := gq.WriteInt32(int32(RootNodeID)); err != nil {
		return nil, err
	}
	if err := c.oscConn.Send(gq); err != nil {
		return nil, err
	}
	// wait for response
	resp := <-c.gqueryTreeChan
	return parseGroup(resp)
}

// ReadBuffer tells the server to read an audio file and
// load it into a buffer
func (c *Client) ReadBuffer(path string, num int32) (*Buffer, error) {
	buf, err := c.sendBufReadMsg(path, num)
	if err != nil {
		return nil, err
	}
	if err := c.awaitBufReadReply(buf); err != nil {

	}
	return buf, nil
}

// sendBufReadMsg sends a /b_allocRead command.
func (c *Client) sendBufReadMsg(path string, num int32) (*Buffer, error) {
	allocRead, err := osc.NewMessage(bufferReadAddress)
	if err != nil {
		return nil, err
	}

	buf := newReadBuffer(path, num, c)
	if err := allocRead.WriteInt32(buf.Num); err != nil {
		return nil, err
	}
	if err := allocRead.WriteString(path); err != nil {
		return nil, err
	}
	if err := c.oscConn.Send(allocRead); err != nil {
		return nil, err
	}
	return buf, nil
}

// awaitBufReadReply waits for a reply to /b_allocRead
func (c *Client) awaitBufReadReply(buf *Buffer) error {
	var done *osc.Message
	select {
	case done = <-c.doneChan:
	case err := <-c.errChan:
		return err
	}

	// error if this message was not an ack of the buffer read
	if done.CountArguments() != 2 {
		return fmt.Errorf("expected two arguments to /done message")
	}
	addr, err := done.ReadString()
	if err != nil {
		return err
	}
	if addr != bufferReadAddress {
		c.doneChan <- done
	}
	bufnum, err := done.ReadInt32()
	if err != nil {
		return err
	}
	if bufnum != buf.Num {
		c.doneChan <- done
	}
	return nil
}

// AllocBuffer allocates a buffer on the server
func (c *Client) AllocBuffer(frames, channels int) (*Buffer, error) {
	buf, err := c.sendBufAllocMsg(frames, channels)
	if err != nil {
		return nil, err
	}
	if err := c.awaitBufAllocReply(buf); err != nil {
		return nil, err
	}
	return buf, nil
}

// sendBufAllocMsg sends a /b_alloc message
func (c *Client) sendBufAllocMsg(frames, channels int) (*Buffer, error) {
	buf := &Buffer{client: c}
	alloc, err := osc.NewMessage(bufferAllocAddress)
	if err != nil {
		return nil, err
	}
	if err := alloc.WriteInt32(buf.Num); err != nil {
		return nil, err
	}
	if err := alloc.WriteInt32(int32(frames)); err != nil {
		return nil, err
	}
	if err := alloc.WriteInt32(int32(channels)); err != nil {
		return nil, err
	}
	if err := c.oscConn.Send(alloc); err != nil {
		return nil, err
	}
	return buf, nil
}

// awaitBufAllocReply waits for a reply to /b_alloc
func (c *Client) awaitBufAllocReply(buf *Buffer) error {
	var done *osc.Message
	select {
	case done = <-c.doneChan:
	case err := <-c.errChan:
		return err
	}
	// error if this message was not an ack of /b_alloc
	if done.CountArguments() != 2 {
		return fmt.Errorf("expected two arguments to /done message")
	}
	addr, err := done.ReadString()
	if err != nil {
		return err
	}
	if addr != bufferAllocAddress {
		c.doneChan <- done

	}
	bufnum, err := done.ReadInt32()
	if err != nil {
		return err
	}
	if bufnum != buf.Num {
		c.doneChan <- done
	}
	return nil
}

// NextSynthID gets the next available ID for creating a synth
func (c *Client) NextSynthID() int32 {
	return atomic.AddInt32(&c.nextSynthID, 1)
}

// FreeAll frees all nodes in a group
func (c *Client) FreeAll(gids ...int32) error {
	freeReq, err := osc.NewMessage(groupFreeAllAddress)
	if err != nil {
		return err
	}
	for _, gid := range gids {
		if err := freeReq.WriteInt32(gid); err != nil {
			return err
		}
	}
	return c.oscConn.Send(freeReq)
}

// addOscHandlers adds OSC handlers
func (c *Client) oscHandlers() osc.Dispatcher {
	return map[string]osc.Method{
		statusReplyAddress: func(msg *osc.Message) error {
			c.statusChan <- msg
			return nil
		},
		doneOscAddress: func(msg *osc.Message) error {
			c.doneChan <- msg
			return nil
		},
		gqueryTreeReplyAddress: func(msg *osc.Message) error {
			c.gqueryTreeChan <- msg
			return nil
		},
	}
}

// PlayDef plays a synthdef by sending the synthdef using
// DefaultClient, then immediately creating a synth node from the def.
func PlayDef(def *Synthdef) (*Synth, error) {
	c, err := DefaultClient()
	if err != nil {
		return nil, err
	}

	if err := c.SendDef(def); err != nil {
		return nil, err
	}

	synthID := c.NextSynthID()
	return defaultGroup.Synth(def.Name, synthID, AddToTail, nil)
}

// Close closes the client.
func (c *Client) Close() error {
	if c.isClosed() {
		return nil
	}
	atomic.StoreInt32(&c.closed, 1)
	if err := c.oscConn.Close(); err != nil {
		return err
	}
	close(c.errChan)
	close(c.doneChan)
	close(c.statusChan)
	close(c.gqueryTreeChan)
	return nil
}

// isClosed says whether or not the client is closed.
func (c *Client) isClosed() bool {
	return atomic.LoadInt32(&c.closed) == int32(1)
}
