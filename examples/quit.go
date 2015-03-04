package main

import (
	"github.com/briansorahan/go-osc/osc"
	"log"
	"net"
)

const (
	listenPort  = 57150
	listenAddr  = "127.0.0.1"
)

// Request status from scsynth
func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:57130")
	if err != nil {
		log.Fatal(err)
	}
	oscServer := osc.NewOscServer(listenAddr, listenPort)
	errChan := make(chan error)
	doneChan := make(chan *osc.OscMessage)
	err = oscServer.AddMsgHandler("/done", func(msg *osc.OscMessage) {
		doneChan <- msg
	})
	if err != nil {
		log.Println("could not send quit message")
		log.Fatal(err)
	}
	go func() {
		errChan <- oscServer.ListenAndDispatch()
	}()
	err = <-oscServer.Listening
	if err != nil {
		log.Fatal(err)
	}
	log.Println("sending quit request")
	quitReq := osc.NewOscMessage("/quit")
	err = oscServer.SendTo(addr, quitReq)
	if err != nil {
		log.Fatal(err)
	}
	select {
	case quitResp := <-doneChan:
		osc.PrintOscMessage(quitResp)
	case err := <-errChan:
		log.Fatal(err)
	}
}
