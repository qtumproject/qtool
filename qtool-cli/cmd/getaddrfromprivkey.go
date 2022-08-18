/*
Copyright © 2022 Alejo Acosta <alejoacos@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/qtumproject/qtool/pkg/tools"
	"github.com/spf13/cobra"
)

var keytoaddressCmd = &cobra.Command{
	Use:   "getaddrfromprivkey",
	Short: "Gets a base 58 address from a given private key",
	Long: `Converts a EC private key to a blockchain address encoded as base58.
Private key supported formats are WIF Base58 or Hex. 
Supported blockchains are 'btc' and 'qtum'.
Supported networks are 'mainnet' and 'testnet'.
For example:
> qtool getaddrfromprivkey -f b58 cMbgxCJrTYUqgcmiC1berh5DFrtY1KeU4PXZ6NZxgenniF1mXCRk -n testnet

Address: qUbxboqjBRp96j3La8D1RYkyqx5uQbJPoW
`,
	RunE:    runGetaddrfromprivkey,
	Args:    cobra.ExactArgs(1),
	Example: "qtool getaddrfromprivkey -f b58 cMbgxCJrTYUqgcmiC1berh5DFrtY1KeU4PXZ6NZxgenniF1mXCRk -n testnet",
	Annotations: map[string]string{
		"test1": "args:getaddrfromprivkey cMbgxCJrTYUqgcmiC1berh5DFrtY1KeU4PXZ6NZxgenniF1mXCRk -f b58, want:qUbxboqjBRp96j3La8D1RYkyqx5uQbJPoW",
		"test2": "args:getaddrfromprivkey 00821d8c8a3627adc68aa4034fea953b2f5da553fab312db3fa274240bd49f35 -f hex, want:qUbxboqjBRp96j3La8D1RYkyqx5uQbJPoW",
	},
}

var privKeyFormat string

func init() {
	rootCmd.AddCommand(keytoaddressCmd)
	keytoaddressCmd.Flags().StringVarP(&privKeyFormat, "format", "f", "", "The encoding format of the private key (hex or b58)")
	keytoaddressCmd.MarkFlagRequired("format")
}

func runGetaddrfromprivkey(cmd *cobra.Command, args []string) error {
	if privKeyFormat != "hex" && privKeyFormat != "b58" {
		return fmt.Errorf("format must be either 'hex' or 'b58'")
	}
	checkFlags(cmd)

	result, err := tools.GetAddressFromPrivkey(args[0], blockchain, network, privKeyFormat)
	if err != nil {
		return err
	}
	printVerbose(cmd, result)
	fmt.Fprintf(cmd.OutOrStdout(), "\n> Result: %s\n", result.AddressBase58)

	return nil
}
