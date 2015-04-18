package main

import (
	"github.com/briansorahan/osc"
	"log"
	"net"
)

const (
	listenPort  = 57240
	listenAddr  = "127.0.0.1"
)

// Request status from scsynth
func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:57130")
	if err != nil {
		log.Fatal(err)
	}
	oscServer := osc.NewServer(listenAddr, listenPort)
	errChan := make(chan error)
	statusChan := make(chan *osc.Message)
	err = oscServer.AddMsgHandler("/status.reply", func(msg *osc.Message) {
		statusChan <- msg
	})
	if err != nil {
		log.Println("could not send status message")
		log.Fatal(err)
	}
	go func() {
		errChan <- oscServer.ListenAndDispatch()
	}()
	err = <-oscServer.Listening
	if err != nil {
		log.Fatal(err)
	}
	log.Println("sending status request")
	statusReq := osc.NewMessage("/status")
	err = oscServer.SendTo(addr, statusReq)
	if err != nil {
		log.Fatal(err)
	}
	select {
	case statusResp := <-statusChan:
		osc.PrintMessage(statusResp)
	case err := <-errChan:
		log.Fatal(err)
	}
}
