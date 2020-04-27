package main

import (
	"fmt"
	"runtime/debug"
)

const unknown = "unknown"

var (
	commit  = unknown
	date    = unknown
	version = "dev"
)

//nolint:gochecknoinits
func init() {
	if info, available := debug.ReadBuildInfo(); available && commit == unknown {
		version = info.Main.Version
		commit = fmt.Sprintf("%s, mod sum: %s", commit, info.Main.Sum)
	}
}

func main() {}
