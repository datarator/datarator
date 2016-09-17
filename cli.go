package main

import (
	"os"

	flags "github.com/jessevdk/go-flags"
)

type options struct {
	Embed   bool `short:"e" long:"embed" description:"Use embedded rather than external static resources" default:"false"`
	Port    int  `short:"p" long:"port" description:"Port to listen on" default:"9292"`
	Timeout int  `short:"t" long:"timeout" description:"Timeout in [ms] for maximum request processing" default:"3000"`
}

var (
	opts   options
	parser = flags.NewParser(&opts, flags.Default)
)

func parseFlags() {
	if _, err := parser.Parse(); err != nil {
		os.Exit(1)
	}
}
