package main

import "flag"

const (
	embedFlagDefault   = true
	embedFlagDesc      = "Whether to use embedded static data or not"
	portFlagDefault    = 9292
	portFlagDesc       = "Port to listen on"
	timeoutFlagDefault = 5000
	timeoutFlagDesc    = "Timeout [ms] for response generation"
)

var (
	embedFlag   = flag.Bool("embed", embedFlagDefault, embedFlagDesc)
	portFlag    = flag.Int("port", portFlagDefault, portFlagDesc)
	timeoutFlag = flag.Int("timeout", 5000, timeoutFlagDesc)
)

func initFlags() {
	flag.BoolVar(embedFlag, "e", embedFlagDefault, embedFlagDesc)
	flag.IntVar(portFlag, "p", portFlagDefault, portFlagDesc)
	flag.IntVar(timeoutFlag, "t", timeoutFlagDefault, timeoutFlagDesc)
}
