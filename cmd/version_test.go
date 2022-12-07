package cmd_test

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/KarolosLykos/advent-of-code-gen/cmd"
)

func TestVersionCommand(t *testing.T) {
	command := cmd.NewRootCmd()

	b := bytes.NewBufferString("")

	command.SetArgs([]string{"version"})
	command.SetOut(b)

	_, err := command.ExecuteC()
	require.NoError(t, err)

	out, err := io.ReadAll(b)
	require.NoError(t, err)

	assert.Equal(t, fmt.Sprintln("v0.0.3"), string(out))
}
