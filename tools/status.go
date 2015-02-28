package main

import (
	"github.com/briansorahan/go-osc/osc"
	"log"
)

const (
	scsynth = "/usr/bin/scsynth"
	scsynthPort = 57116
	listenPort = 57119
	listenAddr = "127.0.0.1"
)

// Request status from scsynth
func main() {
	oscClient := osc.NewOscClient("127.0.0.1", scsynthPort)
	oscServer := osc.NewOscServer(listenAddr, listenPort)
	errChan := make(chan error)
	statusChan := make(chan *osc.OscMessage)
	oscClient.SetLocalAddr(listenAddr, listenPort)
	err := oscServer.AddMsgHandler("/status.reply", func(msg *osc.OscMessage) {
		statusChan <- msg
	})
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		errChan <-oscServer.ListenAndDispatch()
	}()
	log.Println("sending status request")
	statusReq := osc.NewOscMessage("/status")
	err = oscClient.Send(statusReq)
	if err != nil {
		log.Fatal(err)
	}
	select {
	case statusResp := <-statusChan:
		osc.PrintOscMessage(statusResp)
	case err := <-errChan:
		log.Fatal(err)
	}
}
