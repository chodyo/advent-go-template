package main

import (
	"log"
	"os"

	"github.com/jessevdk/go-flags"
)

type options struct {
	Start bool `command:"start" description:"Start a new Advent of Code solution"`

	Advent bool `command:"advent" description:"Run a specific day of Advent of Code"`
}

func main() {
	var opts options
	_, err := flags.ParseArgs(&opts, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if opts.Start {
		if err := start(); err != nil {
			log.Fatal(err)
		}
	}

	if opts.Advent {
		if err := advent(); err != nil {
			log.Fatal(err)
		}
	}

	if !opts.Start && !opts.Advent {
		log.Fatal("No command specified")
	}
}
