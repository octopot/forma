package tools

import (
	"context"

	"github.com/izumin5210/gex"
	"github.com/spf13/cobra"
)

func NewListCommand(cfg *gex.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use: "list",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx, cancel := context.WithCancel(context.TODO())
			defer cancel()

			repository, err := cfg.Create()
			if err != nil {
				return err
			}
			tools, err := repository.List(ctx)
			if err != nil {
				return err
			}
			list := make([]interface{}, 0, len(tools))
			for _, tool := range tools {
				list = append(list, tool.Name())
			}
			cmd.Println(list...)
			return nil
		},
	}
	return cmd
}
