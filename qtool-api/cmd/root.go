/*
Copyright © 2022 Alejo Acosta

*/
package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/alejoacosta74/qtool/qtool-api/server"
	"github.com/spf13/cobra"
)

var (
	debug   bool
	address string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "qtool-server",
	Short: "qtool JSON RPC api server",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: runServer,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&address, "address", "a", ":8080", "address to listen on")
	rootCmd.Flags().BoolVarP(&debug, "debug", "d", false, "enable debug mode")
}

func runServer(cmd *cobra.Command, args []string) error {
	s, err := server.NewServer(debug, address)
	if err != nil {
		return err
	}
	// channel to receive os signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println("\nReceived SIGINT/SIGTERM, exiting gracefully...")
		err := s.Stop()
		if err != nil {
			fmt.Println("Error stopping server:", err)
			os.Exit(1)
		}
		os.Exit(0)
	}()

	return s.Start()
}
