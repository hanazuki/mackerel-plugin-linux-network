package main

import (
	"strings"
	"testing"
)

func TestMapping(t *testing.T) {
	for k, m := range mapping {
		if m == "" {
			continue
		}

		found := false
		for g := range graphs {
			if strings.HasPrefix(m, g + ".") {
				found = true
			}
		}
		if !found {
			t.Errorf("Mapping %s -> %s has no corresponding graph definition", k, m)
		}
	}
}

func TestGraphs(t *testing.T) {
	for g := range graphs {
		found := false
		for _, m := range mapping {
			if strings.HasPrefix(m, g + ".") {
				found = true
			}
		}
		if !found {
			t.Errorf("Graph %s has no metrics", g)
		}
	}
}
