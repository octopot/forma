package main

import (
	_ "github.com/pkg/errors"
	_ "github.com/rubenv/sql-migrate"
	_ "github.com/spf13/cobra"
	_ "github.com/spf13/viper"
)

var (
	commit  = "none"
	date    = "unknown"
	version = "dev"
)
