package main

import (
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/mock"
)

type CmdMock struct {
	mock.Mock
	commands []*cobra.Command
}

func (m *CmdMock) AddCommand(commands ...*cobra.Command) {
	m.commands = commands
	converted := make([]interface{}, 0, len(commands))
	for _, cmd := range commands {
		converted = append(converted, cmd)
	}
	m.Called(converted...)
}

func (m *CmdMock) Execute() error {
	for _, cmd := range m.commands {
		if cmd.Use == "version" {
			cmd.Run(cmd, nil)
			break
		}
	}
	return m.Called().Error(0)
}
