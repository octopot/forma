package cmd

import (
	"time"

	"github.com/kamilsk/form-api/pkg/config"
	"github.com/kamilsk/go-kit/pkg/fn"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cnf      = config.ApplicationConfig{}
	defaults = map[string]interface{}{
		"max_cpus":            1,
		"bind":                "127.0.0.1",
		"http_port":           80,
		"profiling_port":      8090,
		"monitoring_port":     8091,
		"grpc_port":           8092,
		"read_timeout":        time.Duration(0),
		"read_header_timeout": time.Duration(0),
		"write_timeout":       time.Duration(0),
		"idle_timeout":        time.Duration(0),
		"base_url":            "http://localhost/",
		"template_dir":        "static/templates",
		"dsn":                 "postgres://postgres:postgres@127.0.0.1:5432/postgres?connect_timeout=1&sslmode=disable",
		"open_conn":           1,
		"idle_conn":           1,
		"conn_max_lt":         0,
		"table":               "migration",
		"schema":              "public",
	}
)

func db(cmd *cobra.Command) {
	v := viper.New()
	v.SetEnvPrefix("db")
	fn.Must(
		func() error { return v.BindEnv("dsn") },
		func() error { return v.BindEnv("open_conn") },
		func() error { return v.BindEnv("idle_conn") },
		func() error { return v.BindEnv("conn_max_lt") },
		func() error {
			v.SetDefault("dsn", defaults["dsn"])
			v.SetDefault("open_conn", defaults["open_conn"])
			v.SetDefault("idle_conn", defaults["idle_conn"])
			v.SetDefault("conn_max_lt", defaults["conn_max_lt"])
			return nil
		},
		func() error {
			flags := cmd.Flags()
			flags.StringVarP((*string)(&cnf.Union.DBConfig.DSN),
				"dsn", "", v.GetString("dsn"), "data source name")
			flags.IntVarP(&cnf.Union.DBConfig.MaxOpen,
				"db-open-conn", "", v.GetInt("open_conn"), "maximum number of open connections to the database")
			flags.IntVarP(&cnf.Union.DBConfig.MaxIdle,
				"db-idle-conn", "", v.GetInt("idle_conn"), "maximum number of connections in the idle connection pool")
			flags.DurationVarP(&cnf.Union.DBConfig.MaxLifetime,
				"db-conn-max-lt", "", v.GetDuration("conn_max_lt"), "maximum amount of time a connection may be reused")
			return nil
		},
	)
}
