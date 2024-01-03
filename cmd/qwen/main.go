package main

import (
	"os"

	"qwen-cli/cmd/qwen/app"
)

var version string

func main() {
	cmd := app.NewCommand(version)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
