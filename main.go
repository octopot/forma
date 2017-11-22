package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/kamilsk/form-api/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cmd.RootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Show application version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version %s (commit: %s, build date: %s, go version: %s, compiler: %s, platform: %s)\n",
				version, commit, date, runtime.Version(), runtime.Compiler, runtime.GOOS+"/"+runtime.GOARCH)
		},
	})
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
