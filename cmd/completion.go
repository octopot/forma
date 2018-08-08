package cmd

import (
	"github.com/kamilsk/go-kit/pkg/fn"
	"github.com/spf13/cobra"
)

const (
	bashFormat = "bash"
	zshFormat  = "zsh"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Print Bash or Zsh completion",
	RunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Flag("format").Value.String() == bashFormat {
			return cmd.Parent().GenBashCompletion(cmd.OutOrStdout())
		}
		return cmd.Parent().GenZshCompletion(cmd.OutOrStdout())
	},
}

func init() {
	completionCmd.Flags().StringVarP(new(string), "format", "f", zshFormat, "supported formats are bash and zsh")
	fn.Must(func() error { return completionCmd.MarkFlagRequired("format") })
}
