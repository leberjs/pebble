package main

import (
	"log"
	"os"

	"github.com/leberjs/pebble/cmd"
	"github.com/leberjs/pebble/internal/syncer"
	"github.com/leberjs/pebble/internal/tui"
)

func main() {
	cfg, err := cmd.ExecuteArgs()
	if err != nil {
		log.Fatal(err)
	}

	files := syncer.GetSyncFiles()

	tui.Run(cfg, files)

	os.Exit(0)
}
