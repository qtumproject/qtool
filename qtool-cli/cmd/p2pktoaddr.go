/*
Copyright © 2022 alejoacos@gmail.com

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/alejoacosta74/qtool/pkg/tools"
)

const OP_CHECKSIG = "ac"

var p2pkToAddrCmd = &cobra.Command{
	Use:   "p2pktoaddr",
	Short: "Gets the b58 encoded address from a p2pk script",
	Long: `Generates the base 58 encoded address from a p2pk script.
Example:
Input: 210299d391f528b9edd07284c7e23df8415232a8ce41531cf460a390ce32b4efd112ac
Output: qUbxboqjBRp96j3La8D1RYkyqx5uQbJPoW`,
	Example: `qtool p2pktoaddr 210299d391f528b9edd07284c7e23df8415232a8ce41531cf460a390ce32b4efd112ac`,
	RunE:    runP2PKHexToB58AddrCmd,
	Args:    cobra.ExactArgs(1),
	Annotations: map[string]string{
		"test1": "args:p2pktoaddr 210299d391f528b9edd07284c7e23df8415232a8ce41531cf460a390ce32b4efd112ac, want:qUbxboqjBRp96j3La8D1RYkyqx5uQbJPoW",
	},
}

func runP2PKHexToB58AddrCmd(cmd *cobra.Command, args []string) error {
	checkFlags(cmd)
	p2pkOutputStr := args[0]
	result, err := tools.AddressFromP2PKScript(p2pkOutputStr, blockchain, network)
	if err != nil {
		return err
	}
	printVerbose(cmd, result)
	fmt.Fprintf(cmd.OutOrStdout(), "\n> Result: %s\n", result.AddressBase58)
	return nil
}

func init() {
	rootCmd.AddCommand(p2pkToAddrCmd)
}
