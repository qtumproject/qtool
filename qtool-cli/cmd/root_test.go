package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Root(t *testing.T) {
	output := new(bytes.Buffer)
	for _, childCmd := range rootCmd.Commands() {
		for key, val := range childCmd.Annotations {
			args, want := getTestInputsFromAnotations(key, val)
			if args != nil {
				t.Run("Testing cmd: "+childCmd.Name(), func(t *testing.T) {
					output.Reset()
					childCmd.SetOutput(output)
					childCmd.SetErr(output)
					rootCmd.SetArgs(args)
					err := rootCmd.Execute()
					assert.NoError(t, err)
					assert.Contains(t, output.String(), want)

				})
			}
		}
	}
}

// getTestInputsFromAnotations returns the args and want from the annotations map of a cobra command
// val expected format: args:arg1 arg2 , want: want1
// Example: "args: getaddrfromprivkey cMbgxCJrTYUqgcmiC1berh5DFrtY1KeU4PXZ6NZxgenniF1mXCRk -f b58 , want: qUbxboqjBRp96j3La8D1RYkyqx5uQbJPoW"
func getTestInputsFromAnotations(key, val string) (args []string, want string) {
	if strings.Contains(strings.ToLower(key), "test") {
		valSplit := strings.Split(val, ",")
		argsSplit := strings.Split(valSplit[0], ":")
		args = strings.Split(argsSplit[1], " ")
		wantSplit := strings.Split(valSplit[1], ":")
		want = strings.TrimSpace(wantSplit[1])
	}
	return
}
