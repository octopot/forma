package cobra

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.octolab.org/os/shell"
)

const (
	bashFormat       = "bash"
	zshFormat        = "zsh"
	powershellFormat = "powershell"
)

// NewCompletionCommand returns a command that helps to build autocompletion.
//
//  $ source <(cli completion)
//  #
//  # or add into .bash_profile / .zshrc
//  # if [[ -n "$(which cli)" ]]; then
//  #   source <(cli completion)
//  # fi
//  #
//  # or use bash-completion / zsh-completions
//  $ cli completion bash > /path/to/bash_completion.d/cli.sh
//  $ cli completion zsh  > /path/to/zsh-completions/_cli.zsh
//
func NewCompletionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "completion",
		Short: "Print Bash, Zsh or PowerShell completion",
		Long:  "Print Bash, Zsh or PoserShell completion.",
		RunE: func(cmd *cobra.Command, args []string) error {
			sh, err := shell.Classify(os.Getenv("SHELL"), shell.Completion)
			if err != nil {
				return err
			}
			child, args, _ := cmd.Find([]string{sh.String()})
			if child == nil || child == cmd {
				return fmt.Errorf("completion: %s is not supported", sh)
			}
			return child.RunE(child, args)
		},
	}
	cmd.AddCommand(
		&cobra.Command{
			Use:   bashFormat,
			Short: "Print Bash completion",
			Long:  "Print Bash completion.",
			RunE: func(cmd *cobra.Command, args []string) error {
				return root(cmd).GenBashCompletion(cmd.OutOrStdout())
			},
		},
		&cobra.Command{
			Use:   powershellFormat,
			Short: "Print PowerShell completion",
			Long:  "Print PowerShell completion.",
			RunE: func(cmd *cobra.Command, args []string) error {
				return root(cmd).GenPowerShellCompletion(cmd.OutOrStdout())
			},
		},
		&cobra.Command{
			Use:   zshFormat,
			Short: "Print Zsh completion",
			Long:  "Print Zsh completion.",
			RunE: func(cmd *cobra.Command, args []string) error {
				return root(cmd).GenZshCompletion(cmd.OutOrStdout())
			},
		},
	)
	return cmd
}
