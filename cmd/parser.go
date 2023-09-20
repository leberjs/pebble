package cmd

import (
	"strings"

	"golang.org/x/exp/slices"
)

var allowedOpts []string = []string{
	"profile-name",
	"sync-bucket-name",
	"queue-url",
}

func parseArgs(args []string) Command {
	var c Command

	// one arg means only program passed in
	if len(args) == 1 {
		c = Command{Opts: make([]Opt, 0)}
	}

	/*  NOTE:
	    Assuming only opts for now
	    Add parsing commands vs opts if ever needed
	*/
	opts := parseOpts(args[1:])

	c = Command{Opts: opts}

	return c
}

func parseOpts(optArgs []string) []Opt {
	var opts []Opt
	idx := 0

	for i, optArg := range optArgs {
		trimOpt := strings.TrimPrefix(optArg, "--")
		if strings.Contains(trimOpt, "=") {
			s := strings.Split(trimOpt, "=")
			if slices.Contains(allowedOpts, s[0]) {
				o := Opt{Key: s[0], Value: s[1]}
				opts = append(opts, o)
				idx++
			}
		} else {
			if slices.Contains(allowedOpts, trimOpt) {
				o := Opt{Key: trimOpt, Value: optArgs[i+1]}
				opts = append(opts, o)
				idx = idx + 2
			}
		}
	}

	return opts
}
