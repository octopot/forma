package legacy

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	kit "go.octolab.org/strings"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"gopkg.in/yaml.v2"

	"go.octolab.org/ecosystem/forma/internal/config"
	"go.octolab.org/ecosystem/forma/internal/domain"
	pb "go.octolab.org/ecosystem/forma/internal/server/grpc"
	"go.octolab.org/ecosystem/forma/internal/server/grpc/middleware"
)

const (
	schemaKind   kind = "Schema"
	templateKind kind = "Template"
	inputKind    kind = "Input"
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
			inputKind:    func() interface{} { return &inputReadProxy{} },
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

type inputReadProxy struct {
	Filter struct {
		ID        domain.ID `json:"id" mapstructure:"id" yaml:"id"`
		Condition struct {
			SchemaID  domain.ID `json:"schema_id" mapstructure:"schema_id" yaml:"schema_id"`
			CreatedAt struct {
				Start string `json:"start" mapstructure:"start" yaml:"start"`
				End   string `json:"end"   mapstructure:"end" yaml:"end"`
			} `json:"created_at" mapstructure:"created_at" yaml:"created_at"`
		} `json:"condition" mapstructure:"condition" yaml:"condition"`
	} `json:"filter" mapstructure:"filter" yaml:"filter"`
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
	defer func() { _ = conn.Close() }()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = metadata.AppendToOutgoingContext(ctx,
		middleware.AuthHeader,
		kit.Concat(middleware.AuthScheme, " ", string(cnf.Token)))
	switch request := entity.(type) {

	case *pb.CreateSchemaRequest:
		return pb.NewSchemaClient(conn).Create(ctx, request)
	case *pb.CreateTemplateRequest:
		return pb.NewTemplateClient(conn).Create(ctx, request)

	case *pb.ReadSchemaRequest:
		return pb.NewSchemaClient(conn).Read(ctx, request)
	case *pb.ReadTemplateRequest:
		return pb.NewTemplateClient(conn).Read(ctx, request)

	// TODO issue#180
	// - remove hacks with proxies
	// - remove deps to github.com/mitchellh/mapstructure
	// - use github.com/ghodss/yaml and github.com/grpc-ecosystem/grpc-gateway/runtime instead
	case *inputReadProxy:
		grpcRequest := &pb.ReadInputRequest{}
		if request.Filter.ID != "" {
			grpcRequest.Filter = &pb.ReadInputRequest_Id{Id: request.Filter.ID.String()}
		} else {
			var start, end *time.Time
			if request.Filter.Condition.CreatedAt.Start != "" {
				t, parseErr := time.Parse(time.RFC3339, request.Filter.Condition.CreatedAt.Start)
				if parseErr == nil {
					start = &t
				}
			}
			if request.Filter.Condition.CreatedAt.End != "" {
				t, parseErr := time.Parse(time.RFC3339, request.Filter.Condition.CreatedAt.End)
				if parseErr == nil {
					end = &t
				}
			}
			grpcRequest.Filter = &pb.ReadInputRequest_Condition{Condition: &pb.InputFilter{
				SchemaId: request.Filter.Condition.SchemaID.String(),
				CreatedAt: &pb.TimestampRange{
					Start: pb.Timestamp(start),
					End:   pb.Timestamp(end),
				},
			}}
		}
		return pb.NewInputClient(conn).Read(ctx, grpcRequest)

	case *pb.UpdateSchemaRequest:
		return pb.NewSchemaClient(conn).Update(ctx, request)
	case *pb.UpdateTemplateRequest:
		return pb.NewTemplateClient(conn).Update(ctx, request)

	case *pb.DeleteSchemaRequest:
		return pb.NewSchemaClient(conn).Delete(ctx, request)
	case *pb.DeleteTemplateRequest:
		return pb.NewTemplateClient(conn).Delete(ctx, request)

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
