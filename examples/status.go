package main

import (
	"encoding/json"
	"github.com/scgolang/sc"
	"log"
	"os"
)

// Request status from scsynth
func main() {
	client := sc.NewClient("127.0.0.1:57121")
	err := client.Connect("127.0.0.1:57120")
	if err != nil {
		log.Fatal(err)
	}
	status, err := client.GetStatus()
	if err != nil {
		log.Fatal(err)
	}
	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(status)
	if err != nil {
		log.Fatal(err)
	}
}
