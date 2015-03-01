package sc

import (
	"github.com/briansorahan/go-osc/osc"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
)

const (
	scsynth     = "/usr/bin/scsynth"
	defaultPort = 57117
	listenPort  = 57118
	listenAddr  = "127.0.0.1"
)

type Server struct {
	// ErrChan is a channel that emits errors from
	// the goroutine that runs scsynth
	ErrChan    chan error
	addr       NetAddr
	options    ServerOptions
	statusChan chan *osc.OscMessage
	oscClient  *osc.OscClient
	oscServer  *osc.OscServer
	scsynth    *exec.Cmd
}

func (self *Server) Addr() NetAddr {
	return self.addr
}

// Status gets the status of scsynth
func (self *Server) Status() (*ServerStatus, error) {
	statusReq := osc.NewOscMessage("/status")
	err := self.oscClient.Send(statusReq)
	if err != nil {
		return nil, err
	}
	msg := <-self.statusChan
	return newStatus(msg)
}

// Send a synthdef to scsynth
func (self *Server) Send(def *Synthdef) error {
	return nil
}

// Start starts scsynth
func (self *Server) Start() error {
	port := strconv.Itoa(self.addr.Port)
	self.scsynth = exec.Command(scsynth, "-u", port)
	go func() {
		self.ErrChan <- self.scsynth.Run()
	}()
	if self.options.EchoScsynthStdout {
		go func() {
			scsynthStdout, err := self.scsynth.StdoutPipe()
			if err != nil {
				self.ErrChan <-err
				return
			}
			for {
				_, err = io.Copy(os.Stdout, scsynthStdout)
				if err != nil {
					self.ErrChan <-err
					return
				}
			}
		}()
	}
	// stop scsynth on interrupts and kills
	c := make(chan os.Signal)
	go func() {
		<-c
		self.stopScsynth()
		os.Exit(1)
	}()
	signal.Notify(c, os.Interrupt, os.Kill)
	return nil
}

func (self *Server) Wait() error {
	if self.scsynth == nil {
		return nil
	}
	return self.scsynth.Wait()
}

func (self *Server) Close() error {
	var oscErr error
	if self.oscServer != nil {
		oscErr = self.oscServer.Close()
	}
	select {
	case err := <-self.ErrChan:
		return err
	default:
		if oscErr == nil {
			return self.stopScsynth()
		}
		return oscErr
	}
}

func (self *Server) stopScsynth() error {
	if self.scsynth != nil {
		return syscall.Kill(self.scsynth.Process.Pid, syscall.SIGKILL)
	}
	return nil
}

type ServerOptions struct {
	EchoScsynthStdout bool
}

func NewServer(addr NetAddr, options ServerOptions) *Server {
	oscClient := osc.NewOscClient(addr.Addr, addr.Port)
	oscServer := osc.NewOscServer(listenAddr, listenPort)
	statusChan := make(chan *osc.OscMessage)
	oscClient.SetLocalAddr(listenAddr, listenPort)
	oscServer.AddMsgHandler(StatusOscAddress, func(msg *osc.OscMessage) {
		statusChan <- msg
	})
	errChan := make(chan error)
	go func() {
		errChan <- oscServer.ListenAndDispatch()
	}()
	s := Server{
		errChan,
		addr,
		options,
		statusChan,
		oscClient,
		oscServer,
		nil,
	}
	return &s
}
