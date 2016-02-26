package main

import (
	"flag"
	"os"

	mp "github.com/mackerelio/go-mackerel-plugin-helper"
)

type opts struct {
	tempfile string
	pluginOpts
}

func doMain(opts *opts) int {
	helper := mp.NewMackerelPlugin(&opts.pluginOpts)
	helper.Tempfile = opts.tempfile
	helper.Run()
	return 0
}

func main() {
	var opts opts

	flag.StringVar(&opts.tempfile, "tempfile", "/tmp/mackerel-plugin-linux-network", "Path to temporary file")
	flag.StringVar(&opts.prefix, "prefix", "network.", "Prefix for metrics names")
	flag.Parse()

	os.Exit(doMain(&opts))
}
