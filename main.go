package main

import (
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/spf13/cobra"

	cmd "go.octolab.org/ecosystem/forma/internal/cmd/legacy"
	"go.octolab.org/ecosystem/forma/internal/errors"
)

const (
	success = 0
	failed  = 1
)

var (
	commit  = "none"
	date    = "unknown"
	version = "dev"
)

func main() { application{Commander: cmd.RootCmd, Output: os.Stderr, Shutdown: os.Exit}.run() }

type application struct {
	Commander interface {
		AddCommand(...*cobra.Command)
		Execute() error
	}
	Output   io.Writer
	Shutdown func(code int)
}

func (app application) run() {
	var err error
	defer func() {
		errors.Recover(&err)
		if err != nil {
			// so, when `issue` project will be ready
			// I have to integrate it to open GitHub issues
			// with stack trace from terminal
			app.Shutdown(failed)
		}
	}()
	app.Commander.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Show application version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(app.Output,
				"Version %s (commit: %s, build date: %s, go version: %s, compiler: %s, platform: %s)\n",
				version, commit, date, runtime.Version(), runtime.Compiler, runtime.GOOS+"/"+runtime.GOARCH)
		},
		Version: version,
	})
	if err = app.Commander.Execute(); err != nil {
		app.Shutdown(failed)
	}
	app.Shutdown(success)
}
