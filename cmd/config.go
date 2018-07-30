package cmd

import (
	"github.com/kamilsk/form-api/pkg/config"
	"github.com/kamilsk/go-kit/pkg/fn"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cnf = config.ApplicationConfig{}

func db(cmd *cobra.Command) {
	v := viper.New()
	v.SetEnvPrefix("db")
	fn.Must(

		// TODO issue#140 start
		func() error { return v.BindEnv("host") },
		func() error { return v.BindEnv("port") },
		func() error { return v.BindEnv("user") },
		func() error { return v.BindEnv("pass") },
		func() error { return v.BindEnv("name") },
		func() error { return v.BindEnv("timeout") },
		func() error { return v.BindEnv("ssl_mode") },
		// TODO issue#140 end
		func() error { return v.BindEnv("dsn") },

		func() error { return v.BindEnv("open_conn") },
		func() error { return v.BindEnv("idle_conn") },
		func() error { return v.BindEnv("conn_max_lt") },
		func() error {

			// TODO issue#140 start
			v.SetDefault("host", "127.0.0.1")
			v.SetDefault("port", 5432)
			v.SetDefault("user", "postgres")
			v.SetDefault("pass", "postgres")
			v.SetDefault("name", "postgres")
			v.SetDefault("timeout", 1)
			v.SetDefault("ssl_mode", "disable")
			// TODO issue#140 end
			v.SetDefault("dsn", "postgres://postgres:postgres@127.0.0.1:5432/postgres?connect_timeout=1&sslmode=disable")

			v.SetDefault("open_conn", 1)
			v.SetDefault("idle_conn", 1)
			v.SetDefault("conn_max_lt", 0)
			return nil
		},
		func() error {
			flags := cmd.Flags()

			// TODO issue#140 start
			flags.String("db-host", v.GetString("host"), "database host")
			flags.Int("db-port", v.GetInt("port"), "database port")
			flags.String("db-user", v.GetString("user"), "database user name")
			flags.String("db-pass", v.GetString("pass"), "database user password")
			flags.String("db-name", v.GetString("name"), "database name")
			flags.Int("db-timeout", v.GetInt("timeout"), "connection timeout in seconds")
			flags.String("db_ssl_mode", v.GetString("ssl_mode"), "ssl mode")
			// TODO issue#140 end
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
