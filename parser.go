package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type parser func(io.Reader) map[string]uint64

// Split "A: B C" into ("A", ["B", "C"])
func splitKeyValues(s, kvSep string) (key string, vals []string, err error) {
	if kv := strings.SplitN(s, kvSep, 2); len(kv) == 2 {
		key = kv[0]
		vals = strings.Fields(kv[1])
	} else {
		err = fmt.Errorf("No key-value separator found in input string.")
	}

	return
}

// Split "A      B" into ("A", "B")
func splitKeyValue(s string) (key string, val string, err error) {
	if kv := strings.Fields(s); len(kv) == 2 {
		key = kv[0]
		val = kv[1]
	} else {
		err = fmt.Errorf("Input string should have exactly two space-separated fields.")
	}

	return
}

// Parse a file in the following format (eg. /proc/net/netstat)
// Keys are listed in a line by categories and the corresponding values in the following line.
//
//   Key: Foo Bar Baz
//   Key: 0 1 42
//   OtherKey: A B C
//   OtherKey: 0 1 2
//   ...
//
// into a map {"KeyFoo": 0, "KeyBar": 1, "KeyBaz": 42, ...}.
func parseCompact(file io.Reader) map[string]uint64 {
	result := make(map[string]uint64)

	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		key, names, err := splitKeyValues(scanner.Text(), ": ")

		if err != nil || !scanner.Scan() {
			break
		}

		key2, values, err := splitKeyValues(scanner.Text(), ": ")

		if err != nil || key != key2 || len(names) != len(values) {
			break
		}

		for i := 0; i < len(names); i++ {
			val, err := strconv.ParseUint(values[i], 10, 64)
			if err == nil {
				result[key+names[i]] = val
			}
		}
	}

	return result
}

// Parse a file in the following format (eg. /proc/net/snmp6)
// Single key-value pair on each line, with the key and value separated by many spaces.
//
//   KeyFoo      0
//   KeyBar      1
//   KeyBaz      42
//   ...
//
// into a map {"KeyFoo": 0, "KeyBar": 1, "KeyBaz": 42, ...}.
func parseTable(file io.Reader) map[string]uint64 {
	result := make(map[string]uint64)

	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		key, value, err := splitKeyValue(scanner.Text())

		if err != nil {
			continue
		}

		valueInt, err := strconv.ParseUint(value, 10, 64)
		if err == nil {
			result[key] = valueInt
		}
	}

	return result
}
