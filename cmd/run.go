package cmd

import (
	"log"
	"net/http"

	"github.com/kamilsk/form-api/dao"
	"github.com/kamilsk/form-api/server"
	"github.com/kamilsk/form-api/server/router/chi"
	"github.com/kamilsk/form-api/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var run = &cobra.Command{
	Use:   "run",
	Short: "Start HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		addr := cmd.Flag("bind").Value.String() + ":" + cmd.Flag("port").Value.String()
		log.Println("starting server at", addr)
		log.Fatal(http.ListenAndServe(addr,
			chi.NewRouter(
				server.New(
					service.New(
						dao.Must(dao.Connection(dsn(cmd))))),
				cmd.Flag("with-profiler").Value.String() == "true")))
	},
}

func init() {
	v := viper.New()
	must(
		func() error { return v.BindEnv("bind") },
		func() error { return v.BindEnv("port") },
	)
	{
		v.SetDefault("bind", "127.0.0.1")
		v.SetDefault("port", 8080)
	}
	{
		run.Flags().String("bind", v.GetString("bind"), "interface to which the server will bind")
		run.Flags().Int("port", v.GetInt("port"), "port on which the server will listen")
		run.Flags().Bool("with-profiler", false, "enable pprof on /debug/pprof")
	}
	db(run)
}
