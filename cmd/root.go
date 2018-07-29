package cmd

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/spf13/cobra"
)

// RootCmd is the entry point.
var RootCmd = &cobra.Command{Use: "form-api", Short: "Forma"}

func init() {
	RootCmd.AddCommand(completionCmd, controlCmd, migrateCmd, runCmd)
}

// TODO issue#147 start
func must(actions ...func() error) {
	for _, action := range actions {
		if err := action(); err != nil {
			panic(err)
		}
	}
}

// TODO issue#147 end

// TODO issue#140 start
func dsn(cmd *cobra.Command) (driver, dsn string, open, idle int) {
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
	return uri.Scheme, uri.String(), asInt(cmd.Flag("db_open_conn").Value), asInt(cmd.Flag("db_idle_conn").Value)
}

func asInt(value fmt.Stringer) int {
	integer, _ := strconv.ParseInt(value.String(), 10, 0)
	return int(integer)
}

// TODO issue#140 end
