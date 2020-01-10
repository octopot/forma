package make

import "github.com/spf13/cobra"

func NewBuildCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "build",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return cmd
}
