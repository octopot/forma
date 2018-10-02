package cmd

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	pb "github.com/kamilsk/form-api/pkg/server/grpc"
	kit "github.com/kamilsk/go-kit/pkg/strings"

	"github.com/kamilsk/form-api/pkg/config"
	"github.com/kamilsk/form-api/pkg/server/grpc/middleware"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"gopkg.in/yaml.v2"
)

const (
	schemaKind   kind = "Schema"
	templateKind kind = "Template"
)

var entities factory

func init() {
	entities = factory{
		createCmd: {
			schemaKind:   func() interface{} { return &pb.CreateSchemaRequest{} },
			templateKind: func() interface{} { return &pb.CreateTemplateRequest{} },
		},
		readCmd: {
			schemaKind:   func() interface{} { return &pb.ReadSchemaRequest{} },
			templateKind: func() interface{} { return &pb.ReadTemplateRequest{} },
		},
		updateCmd: {
			schemaKind:   func() interface{} { return &pb.UpdateSchemaRequest{} },
			templateKind: func() interface{} { return &pb.UpdateTemplateRequest{} },
		},
		deleteCmd: {
			schemaKind:   func() interface{} { return &pb.DeleteSchemaRequest{} },
			templateKind: func() interface{} { return &pb.DeleteTemplateRequest{} },
		},
	}
}

func communicate(cmd *cobra.Command, _ []string) error {
	entity, err := entities.new(cmd)
	if err != nil {
		return err
	}
	if dry, _ := cmd.Flags().GetBool("dry-run"); dry {
		cmd.Printf("%T would be sent with data: ", entity)
		if cmd.Flag("output").Value.String() == jsonFormat {
			return json.NewEncoder(cmd.OutOrStdout()).Encode(entity)
		}
		return json.NewEncoder(cmd.OutOrStdout()).Encode(entity)
	}
	response, err := call(cnf.Union.GRPCConfig, entity)
	if err != nil {
		cmd.Println(err)
		return nil
	}
	if cmd.Flag("output").Value.String() == jsonFormat {
		return json.NewEncoder(cmd.OutOrStdout()).Encode(response)
	}
	return yaml.NewEncoder(cmd.OutOrStdout()).Encode(response)
}

func printSchemas(cmd *cobra.Command, _ []string) error {
	var (
		target   *cobra.Command
		builders map[kind]builder
		found    bool
	)
	use := cmd.Flag("for").Value.String()
	for target, builders = range entities {
		if strings.EqualFold(target.Use, use) {
			found = true
			break
		}
	}
	if !found {
		return errors.Errorf("unknown control command %q", use)
	}
	for k, b := range builders {
		_ = yaml.NewEncoder(cmd.OutOrStdout()).Encode(schema{Kind: k, Payload: convert(b())})
		cmd.Println()
	}
	return nil
}

type builder func() interface{}

type kind string

type schema struct {
	Kind    kind                   `yaml:"kind"`
	Payload map[string]interface{} `yaml:"payload"`
}

type factory map[*cobra.Command]map[kind]builder

func (f factory) new(cmd *cobra.Command) (interface{}, error) {
	data, err := f.data(cmd.Flag("filename").Value.String())
	if err != nil {
		return nil, err
	}
	build, found := f[cmd][data.Kind]
	if !found {
		return nil, errors.Errorf("unknown payload type %q", data.Kind)
	}
	entity := build()
	if err = mapstructure.Decode(data.Payload, &entity); err != nil {
		return nil, errors.Wrapf(err, "trying to decode payload to %#v", entity)
	}
	return entity, nil
}

func (f factory) data(name string) (schema, error) {
	var (
		err error
		out schema
		raw []byte
		src io.Reader = os.Stdin
	)
	if name != "" {
		if src, err = os.Open(name); err != nil {
			return out, errors.Wrapf(err, "trying to open file %q", name)
		}
	} else {
		name = "/dev/stdin"
	}
	if raw, err = ioutil.ReadAll(src); err != nil {
		return out, errors.Wrapf(err, "trying to read file %q", name)
	}
	err = yaml.Unmarshal(raw, &out)
	return out, errors.Wrapf(err, "trying to decode file %q as YAML", name)
}

func call(cnf config.GRPCConfig, entity interface{}) (interface{}, error) {
	conn, err := grpc.Dial(cnf.Interface, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrapf(err, "trying to connect to the gRPC server %q", cnf.Interface)
	}
	defer conn.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx,
		middleware.AuthHeader,
		kit.Concat(middleware.AuthScheme, " ", string(cnf.Token)))
	switch request := entity.(type) {
	case *pb.CreateSchemaRequest:
		client := pb.NewSchemaClient(conn)
		return client.Create(ctx, request)
	case *pb.CreateTemplateRequest:
		client := pb.NewTemplateClient(conn)
		return client.Create(ctx, request)
	case *pb.ReadSchemaRequest:
		client := pb.NewSchemaClient(conn)
		return client.Read(ctx, request)
	case *pb.ReadTemplateRequest:
		client := pb.NewTemplateClient(conn)
		return client.Read(ctx, request)
	case *pb.UpdateSchemaRequest:
		client := pb.NewSchemaClient(conn)
		return client.Update(ctx, request)
	case *pb.UpdateTemplateRequest:
		client := pb.NewTemplateClient(conn)
		return client.Update(ctx, request)
	case *pb.DeleteSchemaRequest:
		client := pb.NewSchemaClient(conn)
		return client.Delete(ctx, request)
	case *pb.DeleteTemplateRequest:
		client := pb.NewTemplateClient(conn)
		return client.Delete(ctx, request)
	default:
		return nil, errors.Errorf("unknown type %T", request)
	}
}

func convert(entity interface{}) map[string]interface{} {
	t := reflect.ValueOf(entity).Elem().Type()
	m := make(map[string]interface{}, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		v, ok := f.Tag.Lookup("json")
		if !ok {
			continue
		}
		p := strings.Split(v, ",")
		if p[0] == "-" {
			continue
		}
		switch f.Type.String() {
		case "[]uint8":
			m[p[0]] = "binary"
		default:
			m[p[0]] = f.Type.String()
		}
	}
	return m
}
