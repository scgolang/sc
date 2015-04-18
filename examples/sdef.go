package main

import (
	"flag"
	"fmt"
	. "github.com/briansorahan/sc"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "%s [OPTIONS] FILE\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "OPTIONS\n")
	fmt.Fprintf(os.Stderr, "  -format json|xml|dot       Output format\n")
}

// Write json data describing the structure of a synthdef
func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
	format := flag.String("format", "json", "Output format")
	flag.Parse()
	r, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
	d, err := ReadSynthdef(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
	if *format == "json" {
		d.WriteJSON(os.Stdout)
	} else if *format == "dot" {
		d.WriteGraph(os.Stdout)
	} else {
		d.WriteXML(os.Stdout)
	}
}
