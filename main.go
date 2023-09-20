package main

import (
	"log"
	"os"

	"github.com/leberjs/pebble/cmd"
	"github.com/leberjs/pebble/internal/tui"
)

func main() {
    cfg, err := cmd.ExecuteArgs()
    if err != nil {
        log.Fatal(err)
    }

    tui.Run(cfg)

    os.Exit(0)
}
