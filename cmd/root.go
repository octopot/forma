package cmd

import (
	"net/url"

	"github.com/kamilsk/form-api/pkg/config"
	"github.com/spf13/cobra"
)

// RootCmd is the entry point.
var RootCmd = &cobra.Command{Use: "form-api", Short: "Forma"}

func init() {
	RootCmd.AddCommand(completionCmd, controlCmd, migrateCmd, runCmd)
}

// TODO issue#140 start
func dsn(cmd *cobra.Command) (driver, dsn string, config config.DBConfig) {
	uri := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cmd.Flag("db-user").Value.String(), cmd.Flag("db-pass").Value.String()),
		Host:   cmd.Flag("db-host").Value.String() + ":" + cmd.Flag("db-port").Value.String(),
		Path:   "/" + cmd.Flag("db-name").Value.String(),
		RawQuery: func() string {
			query := url.Values{}
			query.Add("connect_timeout", cmd.Flag("db-timeout").Value.String())
			query.Add("sslmode", cmd.Flag("db-ssl-mode").Value.String())
			return query.Encode()
		}(),
	}
	return uri.Scheme, uri.String(), cnf.Union.DBConfig
}

// TODO issue#140 end
