package main

import (
	"log"
	"os"

	"github.com/leberjs/pebble/cmd"
)

func main() {
	if err := cmd.New(os.Args).Execute(); err != nil {
		log.Fatal(err)
	}
}
