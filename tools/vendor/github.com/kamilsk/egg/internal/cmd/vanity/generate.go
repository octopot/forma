package vanity

import "github.com/spf13/cobra"

func NewGenerateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "generate",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return cmd
}
