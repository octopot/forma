package cmd

import (
	"log"
	"net/http"
	"runtime"
	"time"

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
		runtime.GOMAXPROCS(asInt(cmd.Flag("cpus").Value))
		addr := cmd.Flag("bind").Value.String() + ":" + cmd.Flag("port").Value.String()
		handler := chi.NewRouter(
			server.New(
				cmd.Flag("base-url").Value.String(),
				cmd.Flag("tpl-dir").Value.String(),
				service.New(dao.Must(dao.Connection(dsn(cmd))))),
			asBool(cmd.Flag("with-profiler").Value))
		srv := &http.Server{Addr: addr, Handler: handler,
			ReadTimeout:       asDuration(cmd.Flag("read-timeout").Value),
			ReadHeaderTimeout: asDuration(cmd.Flag("read-header-timeout").Value),
			WriteTimeout:      asDuration(cmd.Flag("write-timeout").Value),
			IdleTimeout:       asDuration(cmd.Flag("idle-timeout").Value)}
		log.Println("starting server at", addr)
		log.Fatal(srv.ListenAndServe())
	},
}

func init() {
	v := viper.New()
	must(
		func() error { return v.BindEnv("max_cpus") },
		func() error { return v.BindEnv("bind") },
		func() error { return v.BindEnv("port") },
		func() error { return v.BindEnv("read_timeout") },
		func() error { return v.BindEnv("read_header_timeout") },
		func() error { return v.BindEnv("write_timeout") },
		func() error { return v.BindEnv("idle_timeout") },
		func() error { return v.BindEnv("base_url") },
		func() error { return v.BindEnv("template_dir") },
	)
	{
		v.SetDefault("max_cpus", 1)
		v.SetDefault("bind", "127.0.0.1")
		v.SetDefault("port", 8080)
		v.SetDefault("read_timeout", time.Duration(0))
		v.SetDefault("read_header_timeout", time.Duration(0))
		v.SetDefault("write_timeout", time.Duration(0))
		v.SetDefault("idle_timeout", time.Duration(0))
		v.SetDefault("base_url", "http://127.0.0.1:8080/")
		v.SetDefault("template_dir", "")
	}
	{
		runCmd.Flags().Int("cpus", v.GetInt("max_cpus"),
			"maximum number of CPUs that can be executing simultaneously")
		runCmd.Flags().String("bind", v.GetString("bind"),
			"interface to which the server will bind")
		runCmd.Flags().Int("port", v.GetInt("port"),
			"port on which the server will listen")
		runCmd.Flags().Duration("read-timeout", v.GetDuration("read_timeout"),
			"maximum duration for reading the entire request, including the body")
		runCmd.Flags().Duration("read-header-timeout", v.GetDuration("read_header_timeout"),
			"amount of time allowed to read request headers")
		runCmd.Flags().Duration("write-timeout", v.GetDuration("write_timeout"),
			"maximum duration before timing out writes of the response")
		runCmd.Flags().Duration("idle-timeout", v.GetDuration("idle_timeout"),
			"maximum amount of time to wait for the next request when keep-alive is enabled")
		runCmd.Flags().Bool("with-profiler", false,
			"enable pprof on /debug/pprof")
		runCmd.Flags().String("base-url", v.GetString("base_url"),
			"hostname (and path) to the root")
		runCmd.Flags().String("tpl-dir", v.GetString("template_dir"),
			"filesystem path to custom template directory")
	}
	db(runCmd)
}
