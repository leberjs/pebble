package cmd

import (
	"flag"

	"github.com/leberjs/pebble/constants"
	"github.com/leberjs/pebble/internal/config"
	"github.com/leberjs/pebble/internal/syncer"
)

func ExecuteArgs() (*config.Config, error) {
	profileName := flag.String(
		constants.ProfileName,
		"default",
		"aws profile name set in `credentials`",
	)

	syncBucket := flag.String(
		constants.SyncBucket,
		"",
		"aws s3 bucket to sync files from",
	)

	queueUrl := flag.String(
		constants.QueueUrl,
		"",
		"aws queue url",
	)

	flag.Parse()

	cfg, err := config.GetConfig(*profileName, *syncBucket, *queueUrl)

	syncer.EnsureSyncDir()

	return cfg, err
}
