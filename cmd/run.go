package cmd

import (
	"expvar"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"runtime"
	"time"

	pb "github.com/kamilsk/form-api/pkg/server/grpc"

	"github.com/kamilsk/form-api/pkg/dao"
	"github.com/kamilsk/form-api/pkg/server"
	"github.com/kamilsk/form-api/pkg/server/router/chi"
	"github.com/kamilsk/form-api/pkg/service"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start HTTP server",
	RunE: func(cmd *cobra.Command, args []string) error {
		runtime.GOMAXPROCS(asInt(cmd.Flag("cpus").Value))
		addr := cmd.Flag("bind").Value.String() + ":" + cmd.Flag("port").Value.String()

		if err := startGRPC(); err != nil {
			return err
		}
		if asBool(cmd.Flag("with-profiler").Value) {
			go startProfiler()
		}
		if asBool(cmd.Flag("with-monitoring").Value) {
			go startMonitoring()
		}

		handler := chi.NewRouter(
			server.New(
				cmd.Flag("base-url").Value.String(),
				cmd.Flag("tpl-dir").Value.String(),
				service.New(
					dao.Must(dao.Connection(dsn(cmd))),
				),
			),
		)
		srv := &http.Server{Addr: addr, Handler: handler,
			ReadTimeout:       asDuration(cmd.Flag("read-timeout").Value),
			ReadHeaderTimeout: asDuration(cmd.Flag("read-header-timeout").Value),
			WriteTimeout:      asDuration(cmd.Flag("write-timeout").Value),
			IdleTimeout:       asDuration(cmd.Flag("idle-timeout").Value)}
		log.Println("starting server at", addr)
		return srv.ListenAndServe()
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
		v.SetDefault("port", 80)
		v.SetDefault("read_timeout", time.Duration(0))
		v.SetDefault("read_header_timeout", time.Duration(0))
		v.SetDefault("write_timeout", time.Duration(0))
		v.SetDefault("idle_timeout", time.Duration(0))
		v.SetDefault("base_url", "http://localhost/")
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
			"enable pprof on /pprof")
		runCmd.Flags().Bool("with-monitoring", false,
			"enable expvar on /vars")
		runCmd.Flags().String("base-url", v.GetString("base_url"),
			"hostname (and path) to the root")
		runCmd.Flags().String("tpl-dir", v.GetString("template_dir"),
			"filesystem path to custom template directory")
	}
	db(runCmd)
}

func startGRPC() error {
	listener, err := net.Listen("tcp", ":8092")
	if err != nil {
		return err
	}
	go func() {
		srv := grpc.NewServer()
		pb.RegisterSchemaServer(srv, pb.NewSchemaServer())
		pb.RegisterTemplateServer(srv, pb.NewTemplateServer())
		pb.RegisterInputServer(srv, pb.NewInputServer())
		pb.RegisterLogServer(srv, pb.NewLogServer())
		log.Println("starting gRPC at", listener.Addr())
		_ = srv.Serve(listener) // TODO log critical
	}()
	return nil
}

func startMonitoring() {
	mux := &http.ServeMux{}
	expvar.Handler()
	mux.Handle("/monitoring", promhttp.Handler())
	mux.Handle("/vars", expvar.Handler())
	log.Println("starting monitoring at [::]:8091")
	// TODO use net.Listen and http.Serve instead of http.ListenAndServe
	_ = http.ListenAndServe(":8091", mux) // TODO log critical
}

func startProfiler() {
	mux := &http.ServeMux{}
	mux.HandleFunc("/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/pprof/profile", pprof.Profile)
	mux.HandleFunc("/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/pprof/trace", pprof.Trace)
	mux.HandleFunc("/debug/pprof/", pprof.Index) // net/http/pprof.handler.ServeHTTP specificity
	log.Println("starting profiler at [::]:8090")
	// TODO use net.Listen and http.Serve instead of http.ListenAndServe
	_ = http.ListenAndServe(":8090", mux) // TODO log critical
}
