package cmd

import (
	"github.com/spf13/cobra"

	"github.com/kamilsk/egg/internal/cmd/deps"
	"github.com/kamilsk/egg/internal/cmd/make"
	"github.com/kamilsk/egg/internal/cmd/tools"
	"github.com/kamilsk/egg/internal/cmd/vanity"
)

// New returns the new root command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "egg",
		Short: "Extended go get",
		Long:  "Extended go get - alternative for the standard `go get` with a few little but useful features.",

		SilenceErrors: false,
		SilenceUsage:  true,
	}
	cmd.AddCommand(deps.New(), make.New(), tools.New(), vanity.New())
	return cmd
}
