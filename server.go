package sc

import (
	"fmt"
	"github.com/briansorahan/go-osc/osc"
	"io"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"sync/atomic"
	"syscall"
	"time"
)

const (
	scsynth          = "/usr/bin/scsynth"
	scsynthPort      = 57130
	listenPort       = 57140
	listenAddr       = "127.0.0.1"
	statusOscAddress = "/status.reply"
	doneOscAddress   = "/done"
	// see http://doc.sccode.org/Reference/Server-Command-Reference.html#/dumpOSC
	DumpOff          = 0
	DumpParsed       = 1
	DumpContents     = 2
	DumpAll          = 3
	// see http://doc.sccode.org/Reference/Server-Command-Reference.html#/s_new
	AddToHead        = 0
	AddToTail        = 1
	AddBefore        = 2
	AddAfter         = 3
	AddReplace       = 4
	// see http://doc.sccode.org/Reference/default_group.html
	RootNodeID       = 0
	DefaultGroupID   = 1
)

type Server struct {
	// OscErrChan is a channel that emits errors from
	// the goroutine that runs the OSC server that is
	// used to receive messages from scsynth
	OscErrChan chan error
	addr       net.Addr
	options    ServerOptions
	StatusChan chan *osc.OscMessage
	oscServer *osc.OscServer
	scsynth *exec.Cmd
	// doneChan relays the /done message that comes
	// from scsynth
	doneChan chan *osc.OscMessage
	// next synth node ID
	nextSynthID int32
}

type defLoaded struct {
	Name string
}

// Status gets the status of scsynth
func (self *Server) Status() error {
	statusReq := osc.NewOscMessage("/status")
	err := self.oscServer.SendTo(self.addr, statusReq)
	if err != nil {
		return err
	}
	return nil
}

// SendDef sends a synthdef to scsynth.
// This method blocks until a /done message is received
// indicating that the synthdef was loaded
func (self *Server) SendDef(def *Synthdef) error {
	msg := osc.NewOscMessage("/d_recv")
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
func (self *Server) DumpOSC(level int32) error {
	dumpReq := osc.NewOscMessage("/dumpOSC")
	dumpReq.Append(level)
	err := self.oscServer.SendTo(self.addr, dumpReq)
	if err != nil {
		return err
	}
	return nil
}

func (self *Server) NewSynth(name string, id, action, target int32) error {
	synthReq := osc.NewOscMessage("/s_new")
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
func (self *Server) NewGroup(id, action, target int32) error {
	dumpReq := osc.NewOscMessage("/g_new")
	dumpReq.Append(id)
	dumpReq.Append(action)
	dumpReq.Append(target)
	err := self.oscServer.SendTo(self.addr, dumpReq)
	if err != nil {
		return err
	}
	return nil
}

// NextSynthID
func (self *Server) NextSynthID() int32 {
	return atomic.AddInt32(&self.nextSynthID, 1)
}

func (self *Server) ClearSched() error {
	clear := osc.NewOscMessage("/clearSched")
	err := self.oscServer.SendTo(self.addr, clear)
	if err != nil {
		return err
	}
	return nil
}

// Run runs scsynth in a new goroutine and sends
// any errors on the returned channel.
// This method will not return until
// a status message has been successfully received.
// If scsynth returns an error before a status message
// is received, then a runtime panic occurs.
func (self *Server) Run() chan error {
	running := make(chan error)
	go func() {
		running <-self.scsynth.Run()
	}()
	// give scsynth a little time to get ready
	time.Sleep(200 * time.Millisecond)
	// start trying to get status
	self.Status()
	for {
		select {
		case err := <-running:
			panic(err)
		case <-self.StatusChan:
			goto add_default_group
		default:
			time.Sleep(200 * time.Millisecond)
			self.Status()
		}
	}
add_default_group:
	go func() {
		err := self.NewGroup(DefaultGroupID, AddToHead, RootNodeID)
		if err != nil {
			running <-err
		}
	}()
	return running
}

// Quit sends a /quit message to scsynth
func (self *Server) Quit() error {
	quitReq := osc.NewOscMessage("/quit")
	return self.oscServer.SendTo(self.addr, quitReq)
}

func (self *Server) Close() error {
	var oscErr, stopErr error
	if self.oscServer != nil {
		oscErr = self.oscServer.Close()
	}
	stopErr = self.KillScsynth()
	if oscErr == nil {
		return stopErr
	}
	return oscErr
}

func (self *Server) KillScsynth() error {
	if self.scsynth != nil && self.scsynth.Process != nil {
		return syscall.Kill(self.scsynth.Process.Pid, syscall.SIGKILL)
	}
	return nil
}

type ServerOptions struct {
	EchoScsynthStdout bool
}

func NewServer(addr string, port int, options ServerOptions) (*Server, error) {
	oscServer := osc.NewOscServer(listenAddr, listenPort)
	statusChan := make(chan *osc.OscMessage)
	oscServer.AddMsgHandler(statusOscAddress, func(msg *osc.OscMessage) {
		statusChan <- msg
	})
	doneChan := make(chan *osc.OscMessage)
	oscServer.AddMsgHandler(doneOscAddress, func(msg *osc.OscMessage) {
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
	portStr := strconv.Itoa(port)
	scsynth := exec.Command(scsynth, "-u", portStr)
	if options.EchoScsynthStdout {
		go func() {
			scsynthStdout, err := scsynth.StdoutPipe()
			if err != nil {
				errChan <- err
				return
			}
			for {
				_, err = io.Copy(os.Stdout, scsynthStdout)
				if err != nil {
					errChan <- err
					return
				}
			}
		}()
	}
	netAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		return nil, err
	}
	s := Server{
		errChan,
		netAddr,
		options,
		statusChan,
		oscServer,
		scsynth,
		doneChan,
		1000,
	}
	// stop scsynth on interrupts and kills
	c := make(chan os.Signal)
	go func() {
		<-c
		s.KillScsynth()
		os.Exit(1)
	}()
	signal.Notify(c, os.Interrupt, os.Kill)
	return &s, nil
}
