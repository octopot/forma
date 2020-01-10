package main

import (
	"os"

	"go.octolab.org/toolkit/cli/cobra"

	"github.com/kamilsk/egg/internal/cmd"
)

var (
	commit  = "none"
	date    = "unknown"
	version = "dev"
)

func main() {
	root := cmd.New()
	root.SetOut(os.Stdout)
	root.SetErr(os.Stderr)
	root.AddCommand(cobra.NewCompletionCommand(), cobra.NewVersionCommand(version, date, commit))
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
