package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompletion(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	cmd := completionCmd
	cmd.SetOutput(buf)
	defer cmd.SetOutput(nil)
	{
		tests := []struct {
			name    string
			args    []string
			checker func(assert.TestingT, error, ...interface{}) bool
		}{
			{"empty args", nil, assert.Error},
			{"args with invalid format", []string{"shell"}, assert.Error},
			{"args with valid format (bash)", []string{bashFormat}, assert.NoError},
			{"args with valid format (zsh) ", []string{zshFormat}, assert.NoError},
		}
		for _, test := range tests {
			tc := test
			t.Run(test.name, func(t *testing.T) {
				tc.checker(t, cmd.Args(cmd, tc.args))
			})
		}
	}
	{
		tests := []struct {
			name     string
			format   string
			expected string
		}{
			{"Bash", bashFormat, "# bash completion for form-api"},
			{"Zsh", zshFormat, "#compdef form-api"},
		}
		for _, test := range tests {
			tc := test
			t.Run(test.name, func(t *testing.T) {
				buf.Reset()
				assert.NoError(t, cmd.RunE(cmd, []string{tc.format}))
				assert.Contains(t, buf.String(), tc.expected)
			})
		}
	}
}
