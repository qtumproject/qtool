/*
Copyright © 2022 Alejo Acosta <alejoacos@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/alejoacosta74/qtool/pkg/tools"
	"github.com/spf13/cobra"
)

var convertprivkeyCmd = &cobra.Command{
	Use:   "convertprivkey",
	Short: "Converts the encoding of a ECDSA private key",
	Long: `Converts a ECDSA private key from base58 ('b58') encoding to hexadecimal ('hex') encoding and vice versa.	
When converting from 'hex' is important to explicity set the flags for 'network'.
Example:
qtool convertprivkey   cMbgxCJrTYUqgcmiC1berh5DFrtY1KeU4PXZ6NZxgenniF1mXCRk -f b58
00821d8c8a3627adc68aa4034fea953b2f5da553fab312db3fa274240bd49f35
`,
	Example: `qtool convertprivkey cMbgxCJrTYUqgcmiC1berh5DFrtY1KeU4PXZ6NZxgenniF1mXCRk -f b58`,
	Args:    cobra.ExactArgs(1),
	RunE:    runConvertPrivateKey,
	Annotations: map[string]string{
		"test1": "args:convertprivkey 00821d8c8a3627adc68aa4034fea953b2f5da553fab312db3fa274240bd49f35 -f hex, want:cMbgxCJrTYUqgcmiC1berh5DFrtY1KeU4PXZ6NZxgenniF1mXCRk",
		"test2": "args:convertprivkey cMbgxCJrTYUqgcmiC1berh5DFrtY1KeU4PXZ6NZxgenniF1mXCRk -f b58, want:00821d8c8a3627adc68aa4034fea953b2f5da553fab312db3fa274240bd49f35",
	},
}

func init() {
	rootCmd.AddCommand(convertprivkeyCmd)
	convertprivkeyCmd.Flags().StringVarP(&from, "from", "f", "", "The current encoding format of the address (b58 or hex)")
	convertprivkeyCmd.MarkFlagRequired("from")
}

func runConvertPrivateKey(cmd *cobra.Command, args []string) error {
	if from != "b58" && from != "hex" {
		return fmt.Errorf("from must be either 'b58' or 'hex'")
	}
	checkFlags(cmd)
	privKey := args[0]
	var result *tools.ConvertPrivateKeyResult
	var err error
	if from == "b58" {
		result, err = tools.ConvertPrivateKeyToHex(privKey)
	} else {
		result, err = tools.ConvertPrivateKeyToWIF(privKey, network)
	}

	if err != nil {
		return err
	}
	fmt.Fprintf(cmd.OutOrStdout(), "\n> Result: %s\n", result.PrivateKey)
	return nil
}
