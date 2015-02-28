package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	octets, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	for _, octet := range octets {
		fmt.Printf("%X ", octet)
	}
}
