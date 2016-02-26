package main

import (
	"log"
	"os"
)

type pluginOpts struct {
	prefix string
}

type source struct {
	path   string
	parser parser
}

func (s *source) read() (values map[string]uint64, err error) {
	file, err := os.Open(s.path)
	if err != nil {
		return
	}

	values = s.parser(file)
	return
}

var sources = [...]source{
	{path: "/proc/net/netstat", parser: parseCompact},
	{path: "/proc/net/snmp", parser: parseCompact},
	{path: "/proc/net/snmp6", parser: parseTable},
}

func (p *pluginOpts) FetchMetrics() (metrics map[string]interface{}, err error) {
	metrics = make(map[string]interface{})

	for _, source := range sources {
		stats, err := source.read()
		if err != nil {
			log.Printf("Error reading stats file: %v", err)
			continue
		}

		for k, v := range stats {
			name := mapping[k]
			if name != "" {
				metrics[p.prefix+name] = v
			}
		}
	}

	return
}
