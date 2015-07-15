package main

import (
	"flag"
	. "github.com/scgolang/sc"
	"log"
	"os"
)

func main() {
	var err error
	// create a client and connect to the server
	client := NewClient("127.0.0.1", 57121)
	err = client.Connect("127.0.0.1", 57120)
	if err != nil {
		log.Fatal(err)
	}
	formatp := flag.String("format", "json", "output format")
	flag.Parse()
	// query the root node
	grp, err := client.QueryGroup(0)
	if err != nil {
		log.Fatal(err)
	}
	if *formatp == "json" {
		err = grp.WriteJSON(os.Stdout)
	} else if *formatp == "xml" {
		err = grp.WriteXML(os.Stdout)
	}
	if err != nil {
		log.Fatal(err)
	}
}
