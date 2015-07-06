package main

import (
	"flag"
	"github.com/scgolang/sc"
	"log"
	"os"
)

func main() {
	formatp := flag.String("format", "json", "output format")
	flag.Parse()
	client, err := sc.NewClient("127.0.0.1", 57120)
	if err != nil {
		log.Fatal(err)
	}
	if *formatp == "json" {
		err = client.WriteGroupJSON(int32(1), os.Stdout)
	} else if *formatp == "xml" {
		err = client.WriteGroupXML(int32(1), os.Stdout)
	}
	if err != nil {
		log.Fatal(err)
	}
}
