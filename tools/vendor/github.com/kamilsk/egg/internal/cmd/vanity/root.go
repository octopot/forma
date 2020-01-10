package vanity

import "github.com/spf13/cobra"

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vanity",
		Short: "Manage vanity URLs",
		Long:  "Manage vanity URLs.",
	}
	cmd.AddCommand(NewGenerateCommand())
	return cmd
}
