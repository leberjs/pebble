package cmd

import (
	"flag"

	"github.com/leberjs/pebble/constants"
	"github.com/leberjs/pebble/internal/config"
)

func ExecuteArgs() (*config.Config, error) {
	profileName := flag.String(
		constants.ProfileName,
		"",
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

	cfg, err := config.EnsureConfig(*profileName, *syncBucket, *queueUrl)

	return cfg, err
}
