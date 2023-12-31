/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"minialert/app"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run minialert in server mode",
	Long: `Run minialert in server mode: 
	- listen to updates for battery and CPU values
	- provide alerts in the standard output if battery or CPU fall below a specific threshold
	- load user-defined thresholds from miniconf.json file`,
	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			port = 8080
		}
		app.Serve(port)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().IntP("port", "p", 8080, "Port number")
}
