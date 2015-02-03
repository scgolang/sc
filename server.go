package gosc

import (
	"github.com/hypebeast/go-osc/osc"
)

const (
	scsynth = "/usr/bin/scsynth"
	listenPort = 57118
	listenAddr = "127.0.0.1"
)

type ServerBootFunc func(Server)

type Server interface {
	Addr() NetAddr
	Boot(ServerBootFunc)
}

type server struct {
	addr NetAddr
	oscClient *osc.OscClient
	oscServer *osc.OscServer
}

func (self *server) Addr() NetAddr {
	return self.addr
}

func (self *server) Boot(f ServerBootFunc) {
}

func NewServer(addr NetAddr) Server {
	oscClient := osc.NewOscClient(addr.Addr, addr.Port)
	oscServer := osc.NewOscServer(listenAddr, listenPort)

	go oscServer.ListenAndDispatch()

	oscServer.AddMsgHandler("/status.reply", func(msg *osc.OscMessage) {
		osc.PrintOscMessage(msg)
	})

	s := server{
		addr,
		oscClient,
		oscServer,
	}
	return &s
}
