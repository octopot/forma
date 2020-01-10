package tools

import (
	"path/filepath"

	"github.com/izumin5210/gex"
	"github.com/izumin5210/gex/pkg/tool"
	"github.com/spf13/cobra"
)

func NewRegenCommand(cfg *gex.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use: "regen",
		RunE: func(*cobra.Command, []string) error {
			_, _ = cfg.Create() // set defaults
			path := filepath.Join(cfg.RootDir, cfg.ManifestName)
			m, err := tool.NewParser(cfg.FS, cfg.ManagerType).Parse(path)
			if err != nil {
				return err
			}
			return tool.NewWriter(cfg.FS).Write(path, m)
		},
	}
	return cmd
}
