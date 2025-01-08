package main

import (
	"log"
	"os"

	"github.com/jessevdk/go-flags"
)

type options struct {
	Start Start `command:"start" description:"Start a new Advent of Code solution"`

	Advent Advent `command:"advent" description:"Run a specific day of Advent of Code"`
}

func main() {
	var opts options
	_, err := flags.ParseArgs(&opts, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}
