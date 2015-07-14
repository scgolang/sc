package main

import (
	"flag"
	. "github.com/scgolang/sc"
	"log"
	"os"
)

func main() {
	var err error
	formatp := flag.String("format", "json", "output format")
	flag.Parse()
	if *formatp == "json" {
		err = DefaultClient.WriteGroupJSON(int32(0), os.Stdout)
	} else if *formatp == "xml" {
		err = DefaultClient.WriteGroupXML(int32(0), os.Stdout)
	}
	if err != nil {
		log.Fatal(err)
	}
}
