package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const (
	bashFormat = "bash"
	zshFormat  = "zsh"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Print Bash or Zsh completion",

	// TODO issue#150 start
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.Errorf("please provide only %q or %q as an argument", bashFormat, zshFormat)
		}
		if args[0] != bashFormat && args[0] != zshFormat {
			return errors.Errorf("only %q and %q formats are supported, received %q", bashFormat, zshFormat, args[0])
		}
		return nil
	},
	// TODO issue#150 end

	RunE: func(cmd *cobra.Command, args []string) error {
		if args[0] == bashFormat {
			return cmd.Parent().GenBashCompletion(cmd.OutOrStdout())
		}
		return cmd.Parent().GenZshCompletion(cmd.OutOrStdout())
	},
}
