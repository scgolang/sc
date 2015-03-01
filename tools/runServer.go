package main

import (
	"fmt"
	. "github.com/briansorahan/sc"
	"log"
)

func main() {
	addr := NetAddr{"127.0.0.1", 57113}
	options := ServerOptions{
		EchoScsynthStdout: true,
	}
	server := NewServer(addr, options)
	if server == nil {
		log.Fatal(fmt.Errorf("Could not create server"))
	}
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = server.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
