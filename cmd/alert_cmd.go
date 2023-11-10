/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"minialert/app"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// alertCmd represents the alert command
var alertCmd = &cobra.Command{
	Use:   "alert",
	Short: "Retrieve alert history",
	Long:  `Retrieve alert history`,
	Run: func(cmd *cobra.Command, args []string) {
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			logrus.Error(err)
			return
		}
		app.ListAlert(url)
	},
}

func init() {
	clientCmd.AddCommand(alertCmd)
}
