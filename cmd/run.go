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

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		addr := cmd.Flag("bind").Value.String() + ":" + cmd.Flag("port").Value.String()
		log.Println("starting server at", addr)
		log.Fatal(http.ListenAndServe(addr,
			chi.NewRouter(
				server.New(
					cmd.Flag("baseURL").Value.String(),
					cmd.Flag("tplDir").Value.String(),
					service.New(
						dao.Must(dao.Connection(dsn(cmd))))),
				isTrue(cmd.Flag("with-profiler").Value))))
	},
}

func init() {
	v := viper.New()
	must(
		func() error { return v.BindEnv("base_url") },
		func() error { return v.BindEnv("bind") },
		func() error { return v.BindEnv("port") },
		func() error { return v.BindEnv("template_dir") },
	)
	{
		v.SetDefault("base_url", "http://127.0.0.1:8080/")
		v.SetDefault("bind", "127.0.0.1")
		v.SetDefault("port", 8080)
		v.SetDefault("template_dir", "")
	}
	{
		runCmd.Flags().String("baseURL", v.GetString("base_url"),
			"hostname (and path) to the root, e.g. https://kamil.samigullin.info/")
		runCmd.Flags().String("bind", v.GetString("bind"), "interface to which the server will bind")
		runCmd.Flags().Int("port", v.GetInt("port"), "port on which the server will listen")
		runCmd.Flags().String("tplDir", v.GetString("template_dir"), "filesystem path to custom template directory")
		runCmd.Flags().Bool("with-profiler", false, "enable pprof on /debug/pprof")
	}
	db(runCmd)
}
