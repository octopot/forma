package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/kamilsk/form-api/pkg/server"
	"github.com/kamilsk/form-api/pkg/server/grpc"
	"github.com/kamilsk/form-api/pkg/service"
	"github.com/kamilsk/form-api/pkg/storage"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// invariant
var (
	_ server.Service        = service.New(nil, nil)
	_ service.Storage       = storage.Must()
	_ grpc.ProtectedStorage = storage.Must()
)

func TestApplication_Run(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	app := application{Output: buf}

	tests := []struct {
		name string
		cmd  func() interface {
			AddCommand(...*cobra.Command)
			Execute() error
		}
		expected int
	}{
		{
			"success run",
			func() interface {
				AddCommand(...*cobra.Command)
				Execute() error
			} {
				cmd := &cmdMock{}
				cmd.On("AddCommand", mock.Anything)
				cmd.On("Execute").Return(nil)
				return cmd
			},
			success,
		},
		{
			"failed run",
			func() interface {
				AddCommand(...*cobra.Command)
				Execute() error
			} {
				cmd := &cmdMock{}
				cmd.On("AddCommand", mock.Anything)
				cmd.On("Execute").Return(fmt.Errorf("mocking"))
				return cmd
			},
			failed,
		},
		{
			"panicked run",
			func() interface {
				AddCommand(...*cobra.Command)
				Execute() error
			} {
				cmd := &cmdMock{}
				cmd.On("AddCommand", mock.Anything)
				cmd.On("Execute").Run(func(mock.Arguments) { panic("something unexpected") })
				return cmd
			},
			failed,
		},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			buf.Reset()
			app.Commander = tc.cmd()
			app.Shutdown = func(code int) { panic(assert.Equal(t, tc.expected, code)) }
			assert.Panics(t, func() { app.run() })
			assert.Contains(t, buf.String(), "Version dev")
		})
	}
}

type cmdMock struct {
	mock.Mock
	commands []*cobra.Command
}

func (m *cmdMock) AddCommand(commands ...*cobra.Command) {
	m.commands = commands
	converted := make([]interface{}, 0, len(commands))
	for _, cmd := range commands {
		converted = append(converted, cmd)
	}
	m.Called(converted...)
}

func (m *cmdMock) Execute() error {
	for _, cmd := range m.commands {
		if cmd.Use == "version" {
			cmd.Run(cmd, nil)
			break
		}
	}
	return m.Called().Error(0)
}
