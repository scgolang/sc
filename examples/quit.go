package main

import (
	"github.com/scgolang/osc"
	"log"
	"net"
)

const (
	listenAddr = "127.0.0.1:57110"
)

// Send a /quit message to scsynth
func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:57130")
	if err != nil {
		log.Fatal(err)
	}
	oscServer := osc.NewServer(listenAddr)
	errChan := make(chan error)
	doneChan := make(chan *osc.Message)
	err = oscServer.AddMsgHandler("/done", func(msg *osc.Message) {
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
	quitReq := osc.NewMessage("/quit")
	err = oscServer.SendTo(addr, quitReq)
	if err != nil {
		log.Fatal(err)
	}
	select {
	case quitResp := <-doneChan:
		osc.PrintMessage(quitResp)
	case err := <-errChan:
		log.Fatal(err)
	}
}
