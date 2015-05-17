package sc

import (
	"fmt"
	"github.com/scgolang/osc"
	. "github.com/scgolang/sc/types"
	"net"
	"sync/atomic"
)

const (
	scsynthPort      = 57100
	listenPort       = 57200
	listenAddr       = "127.0.0.1"
	statusOscAddress = "/status.reply"
	doneOscAddress   = "/done"
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

type Client struct {
	// OscErrChan is a channel that emits errors from
	// the goroutine that runs the OSC server that is
	// used to receive messages from scsynth
	OscErrChan chan error
	addr       net.Addr
	statusChan chan *osc.Message
	oscServer  *osc.Server
	// doneChan relays the /done message that comes
	// from scsynth
	doneChan chan *osc.Message
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
	msg := <-self.statusChan
	return newStatus(msg)
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
	done := <-self.doneChan
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
	err := self.oscServer.SendTo(self.addr, dumpReq)
	if err != nil {
		return err
	}
	return nil
}

func (self *Client) NewSynth(name string, id, action, target int32) error {
	synthReq := osc.NewMessage("/s_new")
	synthReq.Append(name)
	synthReq.Append(id)
	synthReq.Append(action)
	synthReq.Append(target)
	synthReq.Append(int32(0))
	err := self.oscServer.SendTo(self.addr, synthReq)
	if err != nil {
		return err
	}
	return nil
}

// NewGroup
func (self *Client) NewGroup(id, action, target int32) error {
	dumpReq := osc.NewMessage("/g_new")
	dumpReq.Append(id)
	dumpReq.Append(action)
	dumpReq.Append(target)
	err := self.oscServer.SendTo(self.addr, dumpReq)
	if err != nil {
		return err
	}
	return nil
}

func (self *Client) ReadBuffer(path string) Buffer {
	return newBuffer(path)
}

// NextSynthID
func (self *Client) NextSynthID() int32 {
	return atomic.AddInt32(&self.nextSynthID, 1)
}

func (self *Client) ClearSched() error {
	clear := osc.NewMessage("/clearSched")
	err := self.oscServer.SendTo(self.addr, clear)
	if err != nil {
		return err
	}
	return nil
}

func NewClient(addr string, port int) (*Client, error) {
	oscServer := osc.NewServer(listenAddr, listenPort)
	statusChan := make(chan *osc.Message)
	oscServer.AddMsgHandler(statusOscAddress, func(msg *osc.Message) {
		statusChan <- msg
	})
	doneChan := make(chan *osc.Message)
	oscServer.AddMsgHandler(doneOscAddress, func(msg *osc.Message) {
		doneChan <- msg
	})
	errChan := make(chan error)
	go func() {
		errChan <- oscServer.ListenAndDispatch()
	}()
	// wait for the server to start running
	err := <-oscServer.Listening
	if err != nil {
		return nil, err
	}
	netAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		return nil, err
	}
	s := Client{
		errChan,
		netAddr,
		statusChan,
		oscServer,
		doneChan,
		1000,
	}
	return &s, nil
}
