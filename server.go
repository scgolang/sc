package sc

import (
	"fmt"
	// TODO: swap back to hypebeast
	"github.com/briansorahan/go-osc/osc"
	"log"
	"os/exec"
	"strconv"
)

const (
	scsynth = "/usr/bin/scsynth"
	DefaultPort = 57116
	listenPort = 57118
	listenAddr = "127.0.0.1"
)

type Server interface {
	Addr() NetAddr
	Close() error
	Boot() error
}

type server struct {
	addr NetAddr
	statusChan chan *osc.OscMessage
	oscClient *osc.OscClient
	oscServer *osc.OscServer
}

func (self *server) Addr() NetAddr {
	return self.addr
}

func (self *server) Boot() error {
	go self.oscServer.ListenAndDispatch()
	cmd := exec.Command(scsynth, "-u", strconv.Itoa(self.addr.Port))
	// TODO: handle error from Run
	go cmd.Run()
	// if err != nil {
	// 	return err
	// }
	statusReq := osc.NewOscMessage("/status")
	self.oscClient.Send(statusReq)
	// TODO: wait for status reply
	// fmt.Println("waiting for status reply")
	// _ = <-self.statusChan
	// fmt.Println("done waiting for status reply")
	return nil
}

func (self *server) Close() error {
	if self.oscServer == nil {
		log.Fatal(fmt.Errorf("self.oscServer is nil"))
	}
	return self.oscServer.Close()
}

func NewServer(addr NetAddr) Server {
	oscClient := osc.NewOscClient(addr.Addr, addr.Port)
	oscServer := osc.NewOscServer(listenAddr, listenPort)
	statusChan := make(chan *osc.OscMessage)

	oscServer.AddMsgHandler("/status.reply", func(msg *osc.OscMessage) {
		statusChan <- msg
	})

	s := server{
		addr,
		statusChan,
		oscClient,
		oscServer,
	}
	return &s
}
