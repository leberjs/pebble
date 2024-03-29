package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/leberjs/pebble/constants"
	"github.com/leberjs/pebble/internal/config"
	"github.com/leberjs/pebble/internal/syncer"
	"github.com/leberjs/pebble/ui"
	"github.com/spf13/cobra"
)

var (
	pn string
	qu string

	rc = &cobra.Command{
		Use:   "pbl",
		Short: "A simple tool to push messages to an AWS SQS queue",
	}
)

func Execute() {
	if err := rc.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rc.Flags().StringVar(&pn, constants.ProfileName, "", "aws profile name")
	rc.Flags().StringVar(&qu, constants.QueueUrl, "", "aws queue url")

	rc.Run = func(_ *cobra.Command, _ []string) {
		cfg, err := config.GetConfig(pn, qu)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err = cfg.EnsureConfigValues(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		files := syncer.GetSyncFiles()

		m := ui.NewModel(cfg, files)

		p := tea.NewProgram(m)
		if _, err := p.Run(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
