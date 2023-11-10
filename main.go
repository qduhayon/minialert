/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"minialert/cmd"
	"minialert/logger"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logger.CustomFormatter{})

	cmd.Execute()
}
