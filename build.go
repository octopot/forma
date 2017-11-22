package main

import (
	_ "github.com/lib/pq"
	_ "github.com/rubenv/sql-migrate"
	_ "github.com/spf13/viper"
)

var (
	commit  = "none"
	date    = "unknown"
	version = "dev"
)
