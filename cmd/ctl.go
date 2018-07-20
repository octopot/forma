package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	controlCmd = &cobra.Command{
		Use:   "ctl",
		Short: "Communicate with Forma server via gRPC",
	}
)

func init() {
	controlCmd.AddCommand(
		&cobra.Command{
			Use:   "create",
			Short: "Create some kind",
			RunE: func(cmd *cobra.Command, args []string) error {
				log.Println("`ctl create` was called")
				return nil
			},
		},
		&cobra.Command{
			Use:   "get",
			Short: "Get some kind",
			RunE: func(cmd *cobra.Command, args []string) error {
				log.Println("`ctl get` was called")
				return nil
			},
		},
		&cobra.Command{
			Use:   "update",
			Short: "Update some kind",
			RunE: func(cmd *cobra.Command, args []string) error {
				log.Println("`ctl update` was called")
				return nil
			},
		},
		&cobra.Command{
			Use:   "delete",
			Short: "Delete some kind",
			RunE: func(cmd *cobra.Command, args []string) error {
				log.Println("`ctl delete` was called")
				return nil
			},
		},
	)
}
