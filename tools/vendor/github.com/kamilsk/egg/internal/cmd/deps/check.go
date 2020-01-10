package deps

import "github.com/spf13/cobra"

func NewCheckCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "check",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return cmd
}
