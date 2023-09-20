package main

import (
	"log"
	"os"

	"github.com/leberjs/pebble/cmd"
	"github.com/leberjs/pebble/internal/tui"
)

func main() {
	ec, err := cmd.New(os.Args).Execute()
	if err != nil {
		log.Fatal(err)
	}

	tui.New(ec).Run()
}
