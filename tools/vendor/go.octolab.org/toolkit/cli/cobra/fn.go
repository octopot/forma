package cobra

import "github.com/spf13/cobra"

func root(cmd *cobra.Command) *cobra.Command {
	for {
		if !cmd.HasParent() {
			break
		}
		cmd = cmd.Parent()
	}
	return cmd
}
