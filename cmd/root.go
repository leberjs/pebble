package cmd

import (
	"errors"
	"fmt"

	"github.com/leberjs/pebble/config"
)

type ExecutionContext struct {
	command Command
	config  config.PebbleConfig
}

type Command struct {
	opts []Opt
}

type Opt struct {
	key   string
	value string
}

func New(args []string) *ExecutionContext {
	c := parseArgs(args)

	return &ExecutionContext{command: c}
}

func (ec *ExecutionContext) Execute() error {
	ce, cp := config.EnsureConfig()
	if !ce {
		e := fmt.Sprintf("Fresh config created at %s. Please update config or pass in needed args", cp)
		return errors.New(e)
	}

	cfg := config.ReadConfig()
    	ec.config = *cfg

	if len(ec.command.opts) == 0 {
		ce, cp = cfg.EnsureConfigValues()
		if !ce {
			e := fmt.Sprintf("One or more params missing. Please pass them in or update config located at %s", cp)
			return errors.New(e)
		}
	}

    if len(ec.command.opts) == 3 {
        executeWithContext(ec)
    }

    diff := optDifference(ec.command.opts)
    if len(diff) > 0 {
        b := false
        for _, opt := range diff {
            if opt == "profile-name" && ec.config.Settings.AwsProfile == "" {
                b = true
            }

            if opt == "sync-bucket-name" && ec.config.Settings.SyncBucket == "" {
                b = true
            }

            if opt == "queue-url" && ec.config.Settings.QueueUrl == "" {
                b = true
            }
        }

        if b {
			e := fmt.Sprintf("One or more params missing. Please pass them in or update config located at %s", cp)
			return errors.New(e)
        } else {
            executeWithContext(ec)    
        }
    }

	return nil
}

func executeWithContext(ec *ExecutionContext) {
	fmt.Println(ec)
}

func optDifference(opts []Opt) (diff []string) {
    m := make(map[string]bool)

    for _, opt := range opts {
        m[opt.key] = true
    }

    for _, aOpt := range allowedOpts {
        if _, ok := m[aOpt]; !ok {
            diff = append(diff, aOpt)
        }
    }

    return
}
