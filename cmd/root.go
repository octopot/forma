package cmd

import (
	"net/url"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RootCmd.AddCommand(migrate, run)
}

var RootCmd = &cobra.Command{
	Short: "Form API",
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
		func() error { return v.BindEnv("sslmode") },
	)
	{
		v.SetDefault("host", "127.0.0.1")
		v.SetDefault("port", 5432)
		v.SetDefault("user", "postgres")
		v.SetDefault("pass", "postgres")
		v.SetDefault("name", "postgres")
		v.SetDefault("timeout", 1)
		v.SetDefault("sslmode", "disable")
	}
	{
		cmd.Flags().String("db_host", v.GetString("host"), "database host")
		cmd.Flags().Int("db_port", v.GetInt("port"), "database port")
		cmd.Flags().String("db_user", v.GetString("user"), "user name")
		cmd.Flags().String("db_pass", v.GetString("pass"), "user password")
		cmd.Flags().String("db_name", v.GetString("name"), "database name")
		cmd.Flags().Int("db_timeout", v.GetInt("timeout"), "connection timeout")
		cmd.Flags().String("db_sslmode", v.GetString("sslmode"), "ssl mode")
	}
}

func dsn(cmd *cobra.Command) *url.URL {
	return &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cmd.Flag("db_user").Value.String(), cmd.Flag("db_pass").Value.String()),
		Host:   cmd.Flag("db_host").Value.String() + ":" + cmd.Flag("db_port").Value.String(),
		Path:   "/" + cmd.Flag("db_name").Value.String(),
		RawQuery: func() string {
			query := url.Values{}
			query.Add("connect_timeout", cmd.Flag("db_timeout").Value.String())
			query.Add("sslmode", cmd.Flag("db_sslmode").Value.String())
			return query.Encode()
		}(),
	}
}
