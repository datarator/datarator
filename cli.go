package main

import (
	"io"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
)

type options struct {
	Embed   bool `short:"e" long:"embed" description:"Use embedded rather than external static resources"`
	ManPage bool `short:"m" long:"man" description:"Dump man page to stdout" hidden:"true"`
	Port    int  `short:"p" long:"port" description:"Port to listen on" default:"9292"`
	Timeout int  `short:"t" long:"timeout" description:"Timeout in [ms] for maximum request processing" default:"3000"`
}

var (
	opts   options
	parser = flags.NewParser(&opts, flags.Default)
)

func parseFlags() bool {

	parser.ShortDescription = "stateless data generator with HTTP based JSON API"
	parser.LongDescription = "Datarator is the stateless data generator with HTTP based JSON API.\n\nFor full documentation, refer to: http://datarator.readthedocs.io"

	// parser.AddGroup
	if _, err := parser.Parse(); err != nil {
		// TODO detect help usage properly
		if strings.HasPrefix(err.Error(), "Usage:\n") {
			return true
		}
		panic(err.Error())
	}

	if opts.ManPage {
		parser.WriteManPage(os.Stdout)
		return true
	}

	return false
}

func writeManPage(writer io.Writer) {
	parser.WriteManPage(writer)
}
