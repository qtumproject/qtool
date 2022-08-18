/*
Copyright © 2022 Alejo Acosta <alejoacos@gmail.com>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "qtool",
	RunE: func(cmd *cobra.Command, args []string) error {
		if gendocs {
			err := doc.GenMarkdownTree(cmd, "./docs")
			if err != nil {
				return err
			}
		} else {
			cmd.Help()
		}
		return nil
	},
	Version: "1.0.0",
}

var verbose bool
var blockchain string
var network string
var from string
var gendocs bool

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&network, "network", "n", "mainnet", `network type: "testnet" or "mainnet"`)
	rootCmd.PersistentFlags().StringVarP(&blockchain, "blockchain", "b", "qtum", `blockchain: "qtum" or "btc"`)
	rootCmd.Flags().BoolVar(&gendocs, "gendocs", false, "generates documentation in 'docs' folder")
	// rootCmd.AddCommand(convertaddressCmd)
}

// checkFlags outputs a 'warning' message to console if optional flags are not set
func checkFlags(cmd *cobra.Command) {
	if !cmd.Flags().Changed("network") || !cmd.Flags().Changed("blockchain") {
		fmt.Fprintf(cmd.OutOrStdout(), "\nWARNING: 'network' and/or 'blockchain' flags not set, using default values '%s' and '%s'\n", network, blockchain)
	}
}

// printVerbose outputs the result in verbose mode
func printVerbose(cmd *cobra.Command, result interface{}) {
	if verbose {
		b, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "\n> Error: %s\n", err)
		} else {
			fmt.Fprintf(cmd.OutOrStdout(), "\n> Full output:\n"+string(b)+"\n")
		}
	}
}
