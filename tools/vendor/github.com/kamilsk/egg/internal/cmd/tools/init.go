package tools

import (
	"context"
	"path/filepath"

	"github.com/izumin5210/gex"
	"github.com/izumin5210/gex/pkg/tool"
	"github.com/spf13/cobra"
)

func NewInitCommand(cfg *gex.Config) *cobra.Command {
	var (
		force bool
	)
	cmd := &cobra.Command{
		Use: "init",
		RunE: func(*cobra.Command, []string) error {
			ctx, cancel := context.WithCancel(context.TODO())
			defer cancel()

			repository, err := cfg.Create()
			if force {
				path := filepath.Join(cfg.RootDir, cfg.ManifestName)
				m := tool.NewManifest(nil, cfg.ManagerType)
				return tool.NewWriter(cfg.FS).Write(path, m)
			}
			if err != nil {
				return err
			}
			return repository.Add(ctx)
		},
	}
	cmd.Flags().BoolVar(&force, "force", false, "initialize from scratch")
	return cmd
}
