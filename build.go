package main

import (
	_ "github.com/lib/pq"
	_ "github.com/rubenv/sql-migrate"
)

var (
	commit  = "none"
	date    = "unknown"
	version = "dev"
)
