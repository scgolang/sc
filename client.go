package sc

import (
	"fmt"
	"github.com/scgolang/osc"
	"github.com/scgolang/sc/types"
	"net"
	"reflect"
	"sync/atomic"
)

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
	// see http://doc.sccode.org/Reference/Server-Command-Reference.html#/dumpOSC
	DumpOff      = 0
	DumpParsed   = 1
	DumpContents = 2
	DumpAll      = 3
	// see http://doc.sccode.org/Reference/Server-Command-Reference.html#/s_new
	AddToHead  = 0
	AddToTail  = 1
	AddBefore  = 2
	AddAfter   = 3
	AddReplace = 4
	// see http://doc.sccode.org/Reference/default_group.html
	RootNodeID      = 0
	DefaultGroupID  = 1
	GenerateSynthID = -1
)

// Client manages all communication with scsynth
type Client struct {
	// OscErrChan is a channel that emits errors from
	// the goroutine that runs the OSC server that is
	// used to receive messages from scsynth
	oscErrChan chan error
	addr       string
	port       int
	conn       net.Addr
	// statusChan relays /status.reply messages
	statusChan chan *osc.Message
	// doneChan relays /done messages
	doneChan chan *osc.Message
	// gqueryTreeChan relays /done messages
	gqueryTreeChan chan *osc.Message
	oscServer      *osc.Server
	// next synth node ID
	nextSynthID int32
}

// Connect connects to an scsynth instance via UDP.
func (self *Client) Connect(addr string) error {
	var err error
	self.conn, err = net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}
	self.oscServer = osc.NewServer(self.addr)
	// OSC relays
	self.statusChan = make(chan *osc.Message)
	self.doneChan = make(chan *osc.Message)
	self.gqueryTreeChan = make(chan *osc.Message)
	self.oscErrChan = make(chan error)
	// listen for OSC messages
	self.addOscHandlers()
	go func() {
		self.oscErrChan <- self.oscServer.ListenAndDispatch()
	}()
	// wait for the OSC server to start
	err = <-self.oscServer.Listening
	if err != nil {
		return err
	}
	return nil
}

// Status gets the status of scsynth.
func (self *Client) GetStatus() (*ServerStatus, error) {
	statusReq := osc.NewMessage(statusAddress)
	err := self.oscServer.SendTo(self.conn, statusReq)
	if err != nil {
		return nil, err
	}
	select {
	case msg := <-self.statusChan:
		return newStatus(msg)
	case err = <-self.oscErrChan:
		return nil, err
	}
}

// SendDef sends a synthdef to scsynth.
// This method blocks until a /done message is received
// indicating that the synthdef was loaded
func (self *Client) SendDef(def *Synthdef) error {
	msg := osc.NewMessage(synthdefReceiveAddress)
	db, err := def.Bytes()
	if err != nil {
		return err
	}
	msg.Append(db)
	self.oscServer.SendTo(self.conn, msg)
	var done *osc.Message
	select {
	case done = <-self.doneChan:
		goto ParseMessage
	case err = <-self.oscErrChan:
		return err
	}

ParseMessage:
	// error if this message was not an ack of the synthdef
	errmsg := "expected /done with /d_recv argument"
	if done.CountArguments() != 1 {
		return fmt.Errorf(errmsg)
	}
	if addr, isString := done.Arguments[0].(string); !isString || addr != synthdefReceiveAddress {
		return fmt.Errorf(errmsg)
	}
	return nil
}

// DumpOSC sends a /dumpOSC message to scsynth
// level should be DumpOff, DumpParsed, DumpContents, DumpAll
func (self *Client) DumpOSC(level int32) error {
	dumpReq := osc.NewMessage(dumpOscAddress)
	dumpReq.Append(level)
	return self.oscServer.SendTo(self.conn, dumpReq)
}

// NewSynth creates a synth
func (self *Client) Synth(defName string, id, action, target int32, ctls map[string]float32) (*Synth, error) {
	synthReq := osc.NewMessage(synthNewAddress)
	synthReq.Append(defName)
	synthReq.Append(id)
	synthReq.Append(action)
	synthReq.Append(target)
	if ctls != nil {
		for k, v := range ctls {
			synthReq.Append(k)
			synthReq.Append(v)
		}
	}
	err := self.oscServer.SendTo(self.conn, synthReq)
	if err != nil {
		return nil, err
	}
	return newSynth(self, defName, id), nil
}

// NewGroup creates a group
func (self *Client) Group(id, action, target int32) (*Group, error) {
	dumpReq := osc.NewMessage(groupNewAddress)
	dumpReq.Append(id)
	dumpReq.Append(action)
	dumpReq.Append(target)
	err := self.oscServer.SendTo(self.conn, dumpReq)
	if err != nil {
		return nil, err
	}
	return newGroup(self, id), nil
}

// AddDefaltGroup adds the default group
func (self *Client) AddDefaultGroup() (*Group, error) {
	return self.Group(DefaultGroupID, AddToTail, RootNodeID)
}

// QueryGroup g_queryTree for a particular group
func (self *Client) QueryGroup(id int32) (*Group, error) {
	addr := gqueryTreeAddress
	gq := osc.NewMessage(addr)
	gq.Append(int32(RootNodeID))
	err := self.oscServer.SendTo(self.conn, gq)
	if err != nil {
		return nil, err
	}
	// wait for response
	resp := <-self.gqueryTreeChan
	return parseGroup(resp)
}

// ReadBuffer tells the server to read an audio file and
// load it into a buffer
func (self *Client) ReadBuffer(path string) (types.Buffer, error) {
	buf := newReadBuffer(path, self)
	pat := bufferReadAddress
	allocRead := osc.NewMessage(pat)
	allocRead.Append(buf.Num())
	allocRead.Append(path)
	err := self.oscServer.SendTo(self.conn, allocRead)

	var done *osc.Message
	select {
	case done = <-self.doneChan:
		break
	case err = <-self.oscErrChan:
		return nil, err
	}

	// error if this message was not an ack of the synthdef
	if done.CountArguments() != 2 {
		return nil, fmt.Errorf("expected two arguments to /done message")
	}
	if addr, isString := done.Arguments[0].(string); !isString || addr != pat {
		return nil, fmt.Errorf("expected first argument to be %s but got %s", pat, addr)
	}
	var bufnum int32
	var isInt32 bool
	if bufnum, isInt32 = done.Arguments[1].(int32); !isInt32 {
		m := "expected int32 as second argument, but got %s (%v)"
		return nil, fmt.Errorf(m, reflect.TypeOf(done.Arguments[1]), done.Arguments[1])
	}
	// TODO:
	// Don't error if we get a done message for a different buffer.
	// We should probably requeue this particular done message on doneChan.
	if bufnum != buf.Num() {
		m := "expected done message for buffer %d, but got one for buffer %d"
		return nil, fmt.Errorf(m, buf.Num(), bufnum)
	}
	return buf, nil
}

// AllocBuffer allocates a buffer on the server
func (self *Client) AllocBuffer(frames, channels int) (types.Buffer, error) {
	buf := newBuffer(self)
	pat := bufferAllocAddress
	alloc := osc.NewMessage(pat)
	alloc.Append(buf.Num())
	alloc.Append(int32(frames))
	alloc.Append(int32(channels))
	err := self.oscServer.SendTo(self.conn, alloc)
	if err != nil {
		return nil, err
	}

	var done *osc.Message
	select {
	case done = <-self.doneChan:
		break
	case err = <-self.oscErrChan:
		return nil, err
	}

	// error if this message was not an ack of the synthdef
	if done.CountArguments() != 2 {
		return nil, fmt.Errorf("expected two arguments to /done message")
	}
	if addr, isString := done.Arguments[0].(string); !isString || addr != pat {
		return nil, fmt.Errorf("expected first argument to be %s but got %s", pat, addr)
	}
	var bufnum int32
	var isInt32 bool
	if bufnum, isInt32 = done.Arguments[1].(int32); !isInt32 {
		m := "expected int32 as second argument, but got %s (%v)"
		return nil, fmt.Errorf(m, reflect.TypeOf(done.Arguments[1]), done.Arguments[1])
	}
	// TODO:
	// Don't error if we get a done message for a different buffer.
	// We should probably requeue this particular done message on doneChan.
	if bufnum != buf.Num() {
		m := "expected done message for buffer %d, but got one for buffer %d"
		return nil, fmt.Errorf(m, buf.Num(), bufnum)
	}
	return buf, nil
}

// NextSynthID gets the next available ID for creating a synth
func (self *Client) NextSynthID() int32 {
	return atomic.AddInt32(&self.nextSynthID, 1)
}

// FreeAll frees all nodes in a group
func (self *Client) FreeAll(gids ...int32) error {
	freeReq := osc.NewMessage(groupFreeAllAddress)
	for _, gid := range gids {
		freeReq.Append(gid)
	}
	return self.oscServer.SendTo(self.conn, freeReq)
}

// addOscHandlers adds OSC handlers
func (self *Client) addOscHandlers() {
	self.oscServer.AddMsgHandler(statusReplyAddress, func(msg *osc.Message) {
		self.statusChan <- msg
	})
	self.oscServer.AddMsgHandler(doneOscAddress, func(msg *osc.Message) {
		self.doneChan <- msg
	})
	self.oscServer.AddMsgHandler(gqueryTreeReplyAddress, func(msg *osc.Message) {
		self.gqueryTreeChan <- msg
	})
}

// NewClient creates a new SuperCollider client.
// The client will bind to the provided address and port
// to receive messages from scsynth.
func NewClient(addr string) *Client {
	self := new(Client)
	self.addr = addr
	self.nextSynthID = 1000
	return self
}
