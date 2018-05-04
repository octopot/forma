package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	bashFormat = "bash"
	zshFormat  = "zsh"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Print Bash or Zsh completion",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("one argument is required, received %d arg(s)", len(args))
		}
		if args[0] != bashFormat && args[0] != zshFormat {
			return fmt.Errorf("only %q and %q formats are supported, received %q", bashFormat, zshFormat, args[0])
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if args[0] == bashFormat {
			return cmd.Parent().GenBashCompletion(cmd.OutOrStdout())
		}
		return cmd.Parent().GenZshCompletion(cmd.OutOrStdout())
	},
}
