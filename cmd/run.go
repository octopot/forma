package cmd

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/kamilsk/form-api/server"
	"github.com/kamilsk/form-api/server/router/chi"
	"github.com/spf13/cobra"
)

var run = &cobra.Command{
	Use:   "run",
	Short: "Start HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		addr := cmd.Flag("bind").Value.String() + ":" + cmd.Flag("port").Value.String()
		log.Println("starting server at", addr)
		log.Fatal(http.ListenAndServe(addr, chi.NewRouter(
			server.New(), cmd.Flag("with-profiler").Value.String() == "true")))
	},
}

func init() {
	run.Flags().String("bind", func() string {
		env := os.Getenv("BIND")
		if env == "" {
			return "127.0.0.1"
		}
		return env
	}(), "interface to which the server will bind")
	run.Flags().Int("port", func() int {
		env := os.Getenv("PORT")
		if env == "" {
			return 8080
		}
		port, err := strconv.Atoi(env)
		if err != nil {
			panic(err)
		}
		return port
	}(), "port on which the server will listen")
	run.Flags().Bool("with-profiler", false, "enable pprof on /debug/pprof")
}
