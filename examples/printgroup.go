package main

import (
	"github.com/scgolang/sc"
	"log"
	"os"
)

func main() {
	client, err := sc.NewClient("127.0.0.1", 57120)
	if err != nil {
		log.Fatal(err)
	}
	err = client.WriteGroupJSON(int32(1), os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
