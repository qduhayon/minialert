/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Run minialert in HTTP client mode",
	Long: `Run minialert in client mode: 
	- send updates for battery and CPU values
	- get active alerts in the server`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	clientCmd.PersistentFlags().StringP("url", "u", "http://localhost:8080", "server url")
}
