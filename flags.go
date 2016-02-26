package main

import (
	"github.com/codegangsta/cli"
)

var cliFlags = []cli.Flag{
	cliFlagTempfile,
}

var cliFlagTempfile = cli.StringFlag{
	Name:   "tempfile",
	Value:  "/tmp/mackerel-plugin-linux-network",
	Usage:  "set temporary file path.",
	EnvVar: "ENVVAR_TEMPFILE",
}
