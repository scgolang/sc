package sc

import (
	"bytes"
	"fmt"
	"github.com/briansorahan/go-osc/osc"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
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
	DumpOff          = 0x00
	DumpParsed       = 0x01
	DumpContents     = 0x02
	DumpAll          = 0x03
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
	doneChan chan error
}

// Status gets the status of scsynth
func (self *Server) Status() error {
	statusReq := osc.NewOscMessage("/status")
	err := self.oscServer.SendTo(self.addr, statusReq)
	if err != nil {
		return err
	}
	log.Println("status message sent")
	return nil
}

// SendDef sends a synthdef to scsynth
func (self *Server) SendDef(def *Synthdef) error {
	buf := bytes.NewBuffer(make([]byte, 0))
	err := def.Write(buf)
	if err != nil {
		return err
	}
	msg := osc.NewOscMessage("/d_recv")
	msg.Append(buf.Bytes())
	// sclang seems to do this, not quite sure why
	// the second argument is supposed to be an osc
	// message it will send when it loads the synthdef
	msg.Append(int32(0))
	self.oscServer.SendTo(self.addr, msg)
	return nil
}

// DumpOSC sends a /dumpOSC message to scsynth
func (self *Server) DumpOSC(level int32) error {
	dumpReq := osc.NewOscMessage("/dumpOSC")
	dumpReq.Append(level)
	err := self.oscServer.SendTo(self.addr, dumpReq)
	if err != nil {
		return err
	}
	log.Println("dumpOSC message sent")
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
	log.Println("s_new message sent")
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
	log.Println("g_new message sent")
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
		log.Println("scsynth done")
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
		err := self.NewGroup(1, 0, 0)
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
	doneChan := make(chan error)
	oscServer.AddMsgHandler(doneOscAddress, func(msg *osc.OscMessage) {
		// TODO: figure out if there was an error?
		// Maybe also relay /fail messages on this channel?
		doneChan <- nil
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
	log.Println("server listening")	
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
