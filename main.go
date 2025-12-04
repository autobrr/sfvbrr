package main

import (
	"github.com/autobrr/sfvbrr/cmd"
)

var (
	version   string
	buildTime string
)

func main() {
	cmd.SetVersion(version, buildTime)
	cmd.Execute()
}

