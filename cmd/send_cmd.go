/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"minialert/app"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a battery or cpu value",
	Long:  `Send a battery or cpu value`,
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			logrus.Error(err)
			return
		}
		value, err := strconv.Atoi(args[1])
		if err != nil {
			logrus.Error(err)
			return
		}
		app.SendMetric(url, app.Metric{Datatype: args[0], Value: value})
	},
}

func init() {
	clientCmd.AddCommand(sendCmd)
}
