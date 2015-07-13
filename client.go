package sc

import (
	"fmt"
	"github.com/scgolang/osc"
	"github.com/scgolang/sc/types"
	"io"
	"net"
	"reflect"
	"sync/atomic"
)

const (
	ScsynthDefaultPort = 57120
	listenPort         = 57200
	listenAddr         = "127.0.0.1"
	statusOscAddress   = "/status.reply"
	gqueryTreeAddress  = "/g_queryTree.reply"
	doneOscAddress     = "/done"
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
	RootNodeID     = 0
	DefaultGroupID = 1
)

// Client manages all communication with scsynth
type Client struct {
	// OscErrChan is a channel that emits errors from
	// the goroutine that runs the OSC server that is
	// used to receive messages from scsynth
	oscErrChan chan error
	addr       net.Addr
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

type defLoaded struct {
	Name string
}

// Status gets the status of scsynth
func (self *Client) Status() (*ServerStatus, error) {
	statusReq := osc.NewMessage("/status")
	err := self.oscServer.SendTo(self.addr, statusReq)
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
	msg := osc.NewMessage("/d_recv")
	db, err := def.Bytes()
	if err != nil {
		return err
	}
	msg.Append(db)
	self.oscServer.SendTo(self.addr, msg)
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
	if addr, isString := done.Arguments[0].(string); !isString || addr != "/d_recv" {
		return fmt.Errorf(errmsg)
	}
	return nil
}

// DumpOSC sends a /dumpOSC message to scsynth
// level should be DumpOff, DumpParsed, DumpContents, DumpAll
func (self *Client) DumpOSC(level int32) error {
	dumpReq := osc.NewMessage("/dumpOSC")
	dumpReq.Append(level)
	return self.oscServer.SendTo(self.addr, dumpReq)
}

// NewSynth creates a synth
func (self *Client) NewSynth(name string, id, action, target int32) error {
	synthReq := osc.NewMessage("/s_new")
	synthReq.Append(name)
	synthReq.Append(id)
	synthReq.Append(action)
	synthReq.Append(target)
	synthReq.Append(int32(0))
	return self.oscServer.SendTo(self.addr, synthReq)
}

// NewGroup creates a group
func (self *Client) NewGroup(id, action, target int32) error {
	dumpReq := osc.NewMessage("/g_new")
	dumpReq.Append(id)
	dumpReq.Append(action)
	dumpReq.Append(target)
	return self.oscServer.SendTo(self.addr, dumpReq)
}

// QueryGroup g_queryTree for a particular group
func (self *Client) QueryGroup(id int32) (*group, error) {
	addr := "/g_queryTree"
	gq := osc.NewMessage(addr)
	gq.Append(int32(RootNodeID))
	err := self.oscServer.SendTo(self.addr, gq)
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
	buf := newBuffer(path)
	pat := "/b_allocRead"
	allocRead := osc.NewMessage(pat)
	allocRead.Append(buf.Num())
	allocRead.Append(path)
	err := self.oscServer.SendTo(self.addr, allocRead)

	var done *osc.Message
	select {
	case done = <-self.doneChan:
		goto ParseMessage
	case err = <-self.oscErrChan:
		return nil, err
	}

ParseMessage:
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
	freeReq := osc.NewMessage("/g_freeAll")
	for _, gid := range gids {
		freeReq.Append(gid)
	}
	return self.oscServer.SendTo(self.addr, freeReq)
}

// ClearSched causes scsynth to clear all scheduled bundles
func (self *Client) ClearSched() error {
	clear := osc.NewMessage("/clearSched")
	return self.oscServer.SendTo(self.addr, clear)
}

// WriteGroupJson writes a json representation of a group to an io.Writer
func (self *Client) WriteGroupJSON(gid int32, w io.Writer) error {
	grp, err := self.QueryGroup(gid)
	if err != nil {
		return err
	}
	return grp.WriteJSON(w)
}

// WriteGroupXml writes a xml representation of a group to an io.Writer
func (self *Client) WriteGroupXML(gid int32, w io.Writer) error {
	grp, err := self.QueryGroup(gid)
	if err != nil {
		return err
	}
	return grp.WriteXML(w)
}

// addOscHandlers adds OSC handlers
func (self *Client) addOscHandlers() {
	self.oscServer.AddMsgHandler(statusOscAddress, func(msg *osc.Message) {
		self.statusChan <- msg
	})
	self.oscServer.AddMsgHandler(doneOscAddress, func(msg *osc.Message) {
		self.doneChan <- msg
	})
	self.oscServer.AddMsgHandler(gqueryTreeAddress, func(msg *osc.Message) {
		self.gqueryTreeChan <- msg
	})
}

// NewClient creates a new SuperCollider client. addr and port are the listening
// address and port, respectively, of scsynth
func NewClient(addr string, port int) (*Client, error) {
	self := new(Client)
	self.nextSynthID = 1000
	var err error
	self.addr, err = net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		return nil, err
	}
	self.oscServer = osc.NewServer(listenAddr, listenPort)
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
		return nil, err
	}
	return self, nil
}
