package cmd

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/qtumproject/qtool/pkg/tools"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

const (
	sampleFileName = "samples.json"
)

func assertCmdSuccess(t *testing.T, cmd *cobra.Command, args []string, want string) {
	output := new(bytes.Buffer)
	cmd.SetOut(output)
	cmd.SetErr(output)
	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	assert.NoError(t, err)
	assert.Contains(t, output.String(), want)
}

func assertCmdError(t *testing.T, cmd *cobra.Command, args []string) {
	output := new(bytes.Buffer)
	cmd.SetOut(output)
	cmd.SetErr(output)
	rootCmd.SetErr(output)
	rootCmd.SetOut(output)
	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	assert.Error(t, err)
}

func loadSampleValues() (*tools.SampleValues, error) {
	p := filepath.Join("../../pkg/tools/testdata", sampleFileName)
	return tools.LoadSampleValues(p)
}
