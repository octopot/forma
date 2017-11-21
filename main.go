package main

import (
	"log"

	"github.com/kamilsk/form-api/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
