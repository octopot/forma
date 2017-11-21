package cmd

import "github.com/spf13/cobra"

func init() {
	RootCmd.AddCommand(migrate, run)
}

var RootCmd = &cobra.Command{
	Short: "Form API",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
