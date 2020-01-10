package tools

import (
	"github.com/izumin5210/gex"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tools",
		Short: "Manage tools",
		Long:  "Manage tools.",
	}
	cfg := &gex.Config{
		OutWriter: cmd.OutOrStdout(),
		ErrWriter: cmd.ErrOrStderr(),
		InReader:  cmd.InOrStdin(),
	}
	cmd.AddCommand(
		NewAddCommand(cfg),
		NewBuildCommand(cfg),
		NewInitCommand(cfg),
		NewListCommand(cfg),
		NewRegenCommand(cfg),
		NewRunCommand(cfg),
	)
	return cmd
}

var (
	_ = []string{
		"github.com/golang/mock/mockgen",
		"github.com/golangci/golangci-lint/cmd/golangci-lint",
		"golang.org/x/tools/cmd/goimports",
	}
	_ = []string{
		"github.com/golang/protobuf/protoc-gen-go",
		"github.com/gogo/protobuf/protoc-gen-gofast",
	}
	_ = []string{
		"github.com/go-swagger/go-swagger/cmd/swagger",
		"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger",
		"github.com/twitchtv/twirp/protoc-gen-twirp",
	}
)
