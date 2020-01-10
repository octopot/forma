package deps

import "github.com/spf13/cobra"

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deps",
		Short: "Manage dependencies",
		Long:  "Manage dependencies.",
	}
	cmd.AddCommand(NewCheckCommand(), NewListCommand())
	return cmd
}
