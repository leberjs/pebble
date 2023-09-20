package tui

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/leberjs/pebble/cmd"
	"golang.org/x/exp/slices"
)

type RunContext struct {
	profileName string
	syncBucket  string
	queueUrl    string
}

func New(ec *cmd.ExecutionContext) *RunContext {
    var profileName, syncBucket, queueUrl string

	// NOTE: This feels bad but I'm tired
	var idx int

	opts := ec.Command.Opts

	idx = slices.IndexFunc(opts, func(o cmd.Opt) bool { return o.Key == "profile-name" })
	if idx == -1 {
		profileName =  ec.Config.Settings.AwsProfile
	} else {
		profileName =  opts[idx].Value
	}

	idx = slices.IndexFunc(opts, func(o cmd.Opt) bool { return o.Key == "sync-bucket-name" })
	if idx == -1 {
		syncBucket =  ec.Config.Settings.SyncBucket
	} else {
		syncBucket =  opts[idx].Value
	}

	idx = slices.IndexFunc(opts, func(o cmd.Opt) bool { return o.Key == "queue-url" })
	if idx == -1 {
		queueUrl =  ec.Config.Settings.QueueUrl
	} else {
		queueUrl =  opts[idx].Value
	}

    return  &RunContext{
		profileName: profileName,
		syncBucket: syncBucket, 
		queueUrl: queueUrl,   
	}
}

func (rc *RunContext) Run() {
    m:= NewModel(rc.profileName, rc.syncBucket, rc.queueUrl)

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
