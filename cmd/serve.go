/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run minialert in server mode",
	Long: `Run minialert in server mode: 
	- listen to updates for battery and CPU values
	- provide alerts in the standard output if battery or CPU fall below a specific threshold.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("serve called with args: %v\n", args)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	serveCmd.Flags().IntP("port", "p", 8080, "Port number")
}
