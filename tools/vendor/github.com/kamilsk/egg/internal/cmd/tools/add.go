package tools

import (
	"context"

	"github.com/izumin5210/gex"
	"github.com/izumin5210/gex/pkg/tool"
	"github.com/spf13/cobra"
)

func NewAddCommand(cfg *gex.Config) *cobra.Command {
	var (
		build  bool
		option byte
	)
	cmd := &cobra.Command{
		Use: "add",
		RunE: func(_ *cobra.Command, args []string) error {
			if !build {
				option |= tool.SkipBuild
			}

			ctx, cancel := context.WithCancel(context.WithValue(context.TODO(), tool.Option{}, option))
			defer cancel()

			repository, err := cfg.Create()
			if err != nil {
				return err
			}
			return repository.Add(ctx, args...)
		},
	}
	cmd.Flags().BoolVar(&build, "build", false, "run build after add")
	return cmd
}
