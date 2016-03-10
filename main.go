package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	mp "github.com/mackerelio/go-mackerel-plugin-helper"
)

type opts struct {
	tempfile string
	pluginOpts
}

type showMetricsFlag struct {
	prefix string
	opts   *pluginOpts
}

func (f *showMetricsFlag) String() string {
	if v, ok := f.opts.showMetrics[f.prefix]; !ok || v {
		return "true"
	}
	return "false"
}

func (f *showMetricsFlag) Set(x string) error {
	show := x != "false"
	for g := range graphs {
		if f.prefix == "" || g == f.prefix || strings.HasPrefix(g, f.prefix+".") {
			f.opts.showMetrics[g] = show
		}
	}
	return nil
}

func (f *showMetricsFlag) IsBoolFlag() bool { return true }

func doMain(opts *opts) int {
	helper := mp.NewMackerelPlugin(&opts.pluginOpts)
	helper.Tempfile = opts.tempfile
	helper.Run()
	return 0
}

func main() {
	opts := opts{
		pluginOpts: pluginOpts{
			showMetrics: make(map[string]bool),
		},
	}

	flag.StringVar(&opts.tempfile, "tempfile", "/tmp/mackerel-plugin-linux-network", "Path to temporary file")
	flag.StringVar(&opts.prefix, "prefix", "network.", "Prefix for metrics names")

	flag.Var(&showMetricsFlag{prefix: "", opts: &opts.pluginOpts}, "all", "Show all metrics")
	for g := range graphs {
		for i := len(g); i != -1; i = strings.LastIndexByte(g[:i], '.') {
			if gn := g[:i]; gn != "" {
				if fn := strings.Replace(gn, ".", "-", -1); flag.Lookup(fn) == nil {
					flag.Var(&showMetricsFlag{prefix: gn, opts: &opts.pluginOpts}, fn, fmt.Sprintf("Show %s.* metrics", gn))
				}
			}
		}
	}

	flag.Parse()

	os.Exit(doMain(&opts))
}
