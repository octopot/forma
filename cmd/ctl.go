package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.octolab.org/fn"
	"go.octolab.org/strings"
)

const (
	yamlFormat = "yaml"
	jsonFormat = "json"
)

var (
	controlCmd = &cobra.Command{Use: "ctl", Short: "Forma Service Control"}
	createCmd  = &cobra.Command{Use: "create", Short: "Create some kind", RunE: communicate}
	readCmd    = &cobra.Command{Use: "read", Short: "Read some kind", RunE: communicate}
	updateCmd  = &cobra.Command{Use: "update", Short: "Update some kind", RunE: communicate}
	deleteCmd  = &cobra.Command{Use: "delete", Short: "Delete some kind", RunE: communicate}
	schemaCmd  = &cobra.Command{Use: "schema", Short: "Print schemas of another control command", RunE: printSchemas}
)

func init() {
	v := viper.New()
	fn.Must(
		func() error { return v.BindEnv("bind") },
		func() error { return v.BindEnv("grpc_port") },
		func() error { return v.BindEnv("forma_token") },
		func() error {
			v.SetDefault("bind", defaults["bind"])
			v.SetDefault("grpc_port", defaults["grpc_port"])
			v.SetDefault("grpc_host", strings.Concat(v.GetString("bind"), ":", strconv.Itoa(v.GetInt("grpc_port"))))
			v.SetDefault("forma_token", defaults["forma_token"])
			return nil
		},
		func() error {
			flags := controlCmd.PersistentFlags()
			flags.StringVarP(new(string), "filename", "f", "", "entity source (default is standard input)")
			flags.StringVarP(new(string), "output", "o", yamlFormat, fmt.Sprintf(
				"output format, one of: %s|%s",
				jsonFormat, yamlFormat))
			flags.Bool("dry-run", false, "if true, only print the object that would be sent, without sending it")
			flags.StringVarP(&cnf.Union.GRPCConfig.Interface,
				"grpc-host", "", v.GetString("grpc_host"), "gRPC server host")
			flags.DurationVarP(&cnf.Union.GRPCConfig.Timeout,
				"timeout", "t", time.Second, "connection timeout")
			flags.StringVarP((*string)(&cnf.Union.GRPCConfig.Token),
				"token", "", v.GetString("forma_token"), "user access token")
			schemaCmd.Flags().String("for", "", "which command: create, read, update or delete")
			return schemaCmd.MarkFlagRequired("for")
		},
	)
	controlCmd.AddCommand(createCmd, readCmd, updateCmd, deleteCmd, schemaCmd)
}
