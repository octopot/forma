package make

import "github.com/spf13/cobra"

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "make",
		Short: "Manage makefiles",
		Long:  "Manage makefiles.",
	}
	cmd.AddCommand(NewBuildCommand())
	return cmd
}
