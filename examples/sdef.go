package main

import (
	"fmt"
	. "github.com/briansorahan/sc"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "%s FILE\n", os.Args[0])
}

// Write json data describing the structure of a synthdef
func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
	r, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
	d, err := ReadSynthdef(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
	d.WriteJSON(os.Stdout)
}
