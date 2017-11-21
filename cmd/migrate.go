package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var migrate = &cobra.Command{
	Use:   "migrate",
	Short: "Apply database migration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("not implemented yet...")
	},
}
