package main

import (
	"github.com/scgolang/sc"
)

func main() {
	client := sc.NewClient("127.0.0.1:57112")
	err := client.Connect("127.0.0.1:57120")
	if err != nil {
		panic(err)
	}
	err = client.FreeAll(sc.DefaultGroupID)
	if err != nil {
		panic(err)
	}
}
