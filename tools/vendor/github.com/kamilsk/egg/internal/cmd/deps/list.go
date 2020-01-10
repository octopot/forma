package deps

import (
	"bytes"
	"context"
	"encoding/json"
	"os/exec"

	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	type (
		Module struct {
			Path    string
			Version string
		}

		Require struct {
			Path     string
			Version  string
			Indirect bool
		}

		Replace struct {
			Old Module
			New Module
		}

		GoMod struct {
			Module  Module
			Go      string
			Require []Require
			Exclude []Module
			Replace []Replace
		}
	)
	var (
		withIndirect bool
	)
	cmd := &cobra.Command{
		Use: "list",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithCancel(context.TODO())
			defer cancel()

			buf := new(bytes.Buffer)
			command := exec.CommandContext(ctx, "go", "mod", "edit", "-json")
			command.Stdout, command.Stderr = buf, cmd.ErrOrStderr()

			if err := command.Run(); err != nil {
				return err
			}

			var output GoMod
			if err := json.NewDecoder(buf).Decode(&output); err != nil {
				return err
			}
			candidates := make([]interface{}, 0, len(output.Require))
			for _, req := range output.Require {
				if !withIndirect && req.Indirect {
					continue
				}
				candidates = append(candidates, req.Path)
			}
			cmd.Println(candidates...)
			return nil
		},
	}
	cmd.Flags().BoolVar(&withIndirect, "with-indirect", false, "include indirect dependencies")
	return cmd
}
