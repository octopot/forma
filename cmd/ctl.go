package cmd

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/kamilsk/go-kit/pkg/fn"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var (
	controlCmd = &cobra.Command{
		Use:   "ctl",
		Short: "Communicate with Forma server via gRPC",
	}
)

func init() {
	var (
		file = ""
		v    = viper.New()
	)
	fn.Must(
		func() error { return v.BindEnv("forma_token") },
		func() error { return v.BindEnv("grpc_host") },
		func() error {
			v.SetDefault("forma_token", "")
			v.SetDefault("grpc_host", "127.0.0.1:8092")
			return nil
		},
		func() error {
			flags := controlCmd.PersistentFlags()
			flags.StringVarP(&file, "file", "f", file, "entity source (default is stdin)")
			flags.StringVarP(&cnf.Union.GRPCConfig.Interface,
				"grpc-host", "", v.GetString("grpc_host"), "gRPC server host")
			flags.DurationVarP(&cnf.Union.GRPCConfig.Timeout,
				"timeout", "t", time.Second, "connection timeout")
			flags.StringVarP((*string)(&cnf.Union.GRPCConfig.Token),
				"token", "", v.GetString("forma_token"), "user access token")
			return nil
		},
	)
	controlCmd.AddCommand(
		&cobra.Command{
			Use:   "create",
			Short: "Create some kind",
			RunE: func(cmd *cobra.Command, args []string) error {

				var (
					err error
					out struct {
						Kind    string                 `yaml:"kind"`
						Payload map[string]interface{} `yaml:"payload"`
					}
					raw []byte
					src io.Reader = os.Stdin
				)
				if file != "" {
					if src, err = os.Open(file); err != nil {
						return err
					}
				}
				if raw, err = ioutil.ReadAll(src); err != nil {
					return err
				}
				if err = yaml.Unmarshal(raw, &out); err != nil {
					return err
				}

				log.Println("`ctl create` was called", out)
				return nil
			},
		},
		&cobra.Command{
			Use:   "get",
			Short: "Get some kind",
			RunE: func(cmd *cobra.Command, args []string) error {
				log.Println("`ctl get` was called")
				return nil
			},
		},
		&cobra.Command{
			Use:   "update",
			Short: "Update some kind",
			RunE: func(cmd *cobra.Command, args []string) error {
				log.Println("`ctl update` was called")
				return nil
			},
		},
		&cobra.Command{
			Use:   "delete",
			Short: "Delete some kind",
			RunE: func(cmd *cobra.Command, args []string) error {
				log.Println("`ctl delete` was called")
				return nil
			},
		},
	)
}
