package main

import (
	"os"

	"github.com/autobrr/sfvbrr/cmd"
)

var (
	version   = "dev"
	buildTime = "unknown"
)

func main() {
	cmd.SetVersion(version, buildTime)
	if cmd.Execute() != nil {
		os.Exit(1)
	}
}
