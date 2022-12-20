/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/qtumproject/qtool/pkg/script"
	"github.com/spf13/cobra"
)

// scriptparserCmd represents the scriptparser command
var scriptparserCmd = &cobra.Command{
	Use:   "scriptparser",
	Short: "Parse a scriptPubKey Vout",
	Long: `Parses a scriptPubKey Vout passed as a 1-line hex string. 
	For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run:  runScriptParser,
	Args: cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(scriptparserCmd)

}

func runScriptParser(cmd *cobra.Command, args []string) {
	scriptPubKey := args[0]
	disasm, err := script.DisasmScript(scriptPubKey)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(disasm)
	}

}
