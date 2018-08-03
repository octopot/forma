package cmd

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/kamilsk/form-api/pkg/server/grpc"
	"github.com/kamilsk/go-kit/pkg/fn"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type schema struct {
	Kind    string                 `yaml:"kind"`
	Payload map[string]interface{} `yaml:"payload"`
}

type factory map[*cobra.Command]map[string]func() interface{}

func (f factory) new(cmd *cobra.Command) (interface{}, error) {
	data, err := f.data(cmd.Flag("file").Value.String())
	if err != nil {
		return nil, err
	}
	builder, ok := f[cmd][data.Kind]
	if !ok {
		return nil, errors.Errorf("unknown payload type %q", data.Kind)
	}
	entity := builder()
	if err = mapstructure.Decode(data.Payload, &entity); err != nil {
		return nil, errors.Wrapf(err, "trying to decode payload to %#v", entity)
	}
	return entity, nil
}

func (f factory) data(file string) (schema, error) {
	var (
		err error
		out schema
		raw []byte
		src io.Reader = os.Stdin
	)
	if file != "" {
		if src, err = os.Open(file); err != nil {
			return out, errors.Wrapf(err, "trying to open file %q", file)
		}
	} else {
		file = "/dev/stdin"
	}
	if raw, err = ioutil.ReadAll(src); err != nil {
		return out, errors.Wrapf(err, "trying to read file %q", file)
	}
	err = yaml.Unmarshal(raw, &out)
	return out, errors.Wrapf(err, "trying to decode file %q as YAML", file)
}

var entities factory

var (
	controlCmd = &cobra.Command{
		Use:   "ctl",
		Short: "Communicate with Forma server via gRPC",
	}
	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create some kind",
		RunE: func(cmd *cobra.Command, args []string) error {
			entity, err := entities.new(cmd)
			if err != nil {
				return err
			}

			log.Printf("`ctl create` was called, %#v\n", entity)
			return nil
		},
	}
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get some kind",
		RunE: func(cmd *cobra.Command, args []string) error {
			entity, err := entities.new(cmd)
			if err != nil {
				return err
			}

			log.Printf("`ctl get` was called, %#v\n", entity)
			return nil
		},
	}
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update some kind",
		RunE: func(cmd *cobra.Command, args []string) error {
			entity, err := entities.new(cmd)
			if err != nil {
				return err
			}

			log.Printf("`ctl update` was called, %#v\n", entity)
			return nil
		},
	}
	deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete some kind",
		RunE: func(cmd *cobra.Command, args []string) error {
			entity, err := entities.new(cmd)
			if err != nil {
				return err
			}

			log.Printf("`ctl delete` was called, %#v\n", entity)
			return nil
		},
	}
)

func init() {
	v := viper.New()
	fn.Must(
		func() error { return v.BindEnv("forma_token") },
		func() error { return v.BindEnv("grpc_host") },
		func() error {
			v.SetDefault("forma_token", "")
			v.SetDefault("grpc_host", "127.0.0.1:8092")
			return nil
		},
		func() error {
			file := ""
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
		func() error {
			entities = factory{
				createCmd: {
					"Schema":   func() interface{} { return &grpc.CreateSchemaRequest{} },
					"Template": func() interface{} { return &grpc.CreateTemplateRequest{} },
				},
				getCmd: {
					"Schema":   func() interface{} { return &grpc.DeleteSchemaRequest{} },
					"Template": func() interface{} { return &grpc.DeleteTemplateRequest{} },
				},
				updateCmd: {
					"Schema":   func() interface{} { return &grpc.UpdateSchemaRequest{} },
					"Template": func() interface{} { return &grpc.UpdateTemplateRequest{} },
				},
				deleteCmd: {
					"Schema":   func() interface{} { return &grpc.DeleteSchemaRequest{} },
					"Template": func() interface{} { return &grpc.DeleteTemplateRequest{} },
				},
			}
			return nil
		},
	)
	controlCmd.AddCommand(createCmd, getCmd, updateCmd, deleteCmd)
}
