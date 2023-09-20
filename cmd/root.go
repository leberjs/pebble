package cmd

import (
	"errors"
	"fmt"

	"github.com/leberjs/pebble/config"
)

type ExecutionContext struct {
	Command Command
	Config  config.PebbleConfig
}

type Command struct {
	Opts []Opt
}

type Opt struct {
	Key   string
	Value string
}

func New(args []string) *ExecutionContext {
	c := parseArgs(args)

	return &ExecutionContext{Command: c}
}

func (ec *ExecutionContext) Execute() (*ExecutionContext, error) {
	ce, cp := config.EnsureConfig()
	if !ce {
		e := fmt.Sprintf("Fresh config created at %s. Please update config or pass in needed args", cp)
		return nil, errors.New(e)
	}

	cfg := config.ReadConfig()
	ec.Config = *cfg

	if len(ec.Command.Opts) == 0 {
		ce, cp = cfg.EnsureConfigValues()
		if !ce {
			e := fmt.Sprintf("One or more params missing. Please pass them in or update config located at %s", cp)
			return nil, errors.New(e)
		}
	}

	if len(ec.Command.Opts) == 3 {
		return ec, nil
	}

	diff := optDifference(ec.Command.Opts)
	if len(diff) > 0 {
		b := false
		for _, opt := range diff {
			if opt == "profile-name" && ec.Config.Settings.AwsProfile == "" {
				b = true
			}

			if opt == "sync-bucket-name" && ec.Config.Settings.SyncBucket == "" {
				b = true
			}

			if opt == "queue-url" && ec.Config.Settings.QueueUrl == "" {
				b = true
			}
		}

		if b {
			e := fmt.Sprintf("One or more params missing. Please pass them in or update config located at %s", cp)
			return nil, errors.New(e)
		} else {
			return ec, nil
		}
	}

	return ec, nil
}

func optDifference(opts []Opt) (diff []string) {
	m := make(map[string]bool)

	for _, opt := range opts {
		m[opt.Key] = true
	}

	for _, aOpt := range allowedOpts {
		if _, ok := m[aOpt]; !ok {
			diff = append(diff, aOpt)
		}
	}

	return
}
