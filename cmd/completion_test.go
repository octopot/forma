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
			assert.NoError(t, cmd.Flag("format").Value.Set(tc.format))
			assert.NoError(t, cmd.RunE(cmd, nil))
			assert.Contains(t, buf.String(), tc.expected)
		})
	}
}
