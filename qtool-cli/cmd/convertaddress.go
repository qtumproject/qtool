/*
Copyright © 2022 Alejo Acosta <alejoacos@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/alejoacosta74/qtool/pkg/tools"
	"github.com/spf13/cobra"
)

var convertaddressCmd = &cobra.Command{
	Use:   "convertaddress",
	Short: "Converts a legacy address from one encoding to another",
	Long: `Converts a legacy address from base58 ('b58') encoding to hexadecimal ('hex') encoding and vice versa.
When converting from 'hex' is important to explicity set the flags for 'blockchain' and 'network'.
Example:
> qtool convertaddress -f b58 qUbxboqjBRp96j3La8D1RYkyqx5uQbJPoW
Address: 7926223070547d2d15b2ef5e7383e541c338ffe9
`,
	Example: `qtool convertaddress qUbxboqjBRp96j3La8D1RYkyqx5uQbJPoW -f b58`,
	Args:    cobra.ExactArgs(1),
	RunE:    runConvertAddress,
	Annotations: map[string]string{
		"test1": "args:convertaddress 7926223070547d2d15b2ef5e7383e541c338ffe9 -f hex, want:qUbxboqjBRp96j3La8D1RYkyqx5uQbJPoW",
		"test2": "args:convertaddress qUbxboqjBRp96j3La8D1RYkyqx5uQbJPoW -f b58, want:7926223070547d2d15b2ef5e7383e541c338ffe9",
	},
}

func init() {
	rootCmd.AddCommand(convertaddressCmd)
	convertaddressCmd.Flags().StringVarP(&from, "from", "f", "", "The current encoding format of the address (b58 or hex)")
	convertaddressCmd.MarkFlagRequired("from")

}

func runConvertAddress(cmd *cobra.Command, args []string) error {
	if from != "b58" && from != "hex" {
		return fmt.Errorf("from must be either 'b58' or 'hex'")
	}
	checkFlags(cmd)
	var result *tools.ConvertAddressResult
	var err error
	if from == "b58" {
		result, err = tools.ConvertAddressBase58ToHex(args[0])
	} else {
		result, err = tools.ConvertAddressHexToBase58(args[0], blockchain, network)
	}
	if err != nil {
		return err
	}
	fmt.Fprintf(cmd.OutOrStdout(), "> Result: %s\n", result.Address)
	return nil
}
