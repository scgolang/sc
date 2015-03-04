package main

import (
	"fmt"
	. "github.com/briansorahan/sc"
	"log"
)

func main() {
	options := ServerOptions{
		EchoScsynthStdout: true,
	}
	server, err := NewServer("127.0.0.1", 57130, options)
	if server == nil {
		log.Fatal(fmt.Errorf("Could not create server"))
	}
	if err != nil {
		log.Fatal(err)
	}
	// error-handling goroutine
	go func() {
		select {
		case err := <-server.OscErrChan:
			if err != nil {
				log.Println("Error with OSC server")
				log.Fatal(err)
			}
		}
	}()
	done := make(chan error)
	go func() {
		done <-server.Run()
	}()
	// time.Sleep(3 * time.Second)
	// status, err := server.Status()
	// if err != nil {
	// 	log.Println("Could not get server status")
	// 	log.Fatal(err)
	// }
	// if status == nil {
	// 	log.Fatalf("status was nil")
	// }
	err = <-done
	if err != nil {
		log.Fatal(err)
	}
}
