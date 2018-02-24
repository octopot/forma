package cmd

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd is the entry point.
var RootCmd = &cobra.Command{Short: "Form API"}

func init() {
	RootCmd.AddCommand(migrateCmd, runCmd)
}

func must(actions ...func() error) {
	for _, action := range actions {
		if err := action(); err != nil {
			panic(err)
		}
	}
}

func db(cmd *cobra.Command) {
	v := viper.New()
	v.SetEnvPrefix("db")
	must(
		func() error { return v.BindEnv("host") },
		func() error { return v.BindEnv("port") },
		func() error { return v.BindEnv("user") },
		func() error { return v.BindEnv("pass") },
		func() error { return v.BindEnv("name") },
		func() error { return v.BindEnv("timeout") },
		func() error { return v.BindEnv("ssl_mode") },
	)
	{
		v.SetDefault("host", "127.0.0.1")
		v.SetDefault("port", 5432)
		v.SetDefault("user", "postgres")
		v.SetDefault("pass", "postgres")
		v.SetDefault("name", "postgres")
		v.SetDefault("timeout", 1)
		v.SetDefault("ssl_mode", "disable")
	}
	{
		cmd.Flags().String("db_host", v.GetString("host"),
			"database host")
		cmd.Flags().Int("db_port", v.GetInt("port"),
			"database port")
		cmd.Flags().String("db_user", v.GetString("user"),
			"database user name")
		cmd.Flags().String("db_pass", v.GetString("pass"),
			"database user password")
		cmd.Flags().String("db_name", v.GetString("name"),
			"database name")
		cmd.Flags().Int("db_timeout", v.GetInt("timeout"),
			"connection timeout in seconds")
		cmd.Flags().String("db_ssl_mode", v.GetString("ssl_mode"),
			"ssl mode")
	}
}

func dsn(cmd *cobra.Command) (driver, dsn string) {
	uri := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cmd.Flag("db_user").Value.String(), cmd.Flag("db_pass").Value.String()),
		Host:   cmd.Flag("db_host").Value.String() + ":" + cmd.Flag("db_port").Value.String(),
		Path:   "/" + cmd.Flag("db_name").Value.String(),
		RawQuery: func() string {
			query := url.Values{}
			query.Add("connect_timeout", cmd.Flag("db_timeout").Value.String())
			query.Add("sslmode", cmd.Flag("db_ssl_mode").Value.String())
			return query.Encode()
		}(),
	}
	return uri.Scheme, uri.String()
}

func asBool(value fmt.Stringer) bool {
	is, _ := strconv.ParseBool(value.String())
	return is
}

func asDuration(value fmt.Stringer) time.Duration {
	duration, _ := time.ParseDuration(value.String())
	return duration
}

func asInt(value fmt.Stringer) int {
	integer, _ := strconv.ParseInt(value.String(), 10, 0)
	return int(integer)
}
