package main

import (
	"github.com/deepsquare-io/the-grid/cli/cmd/submit"
	"github.com/deepsquare-io/the-grid/cli/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var app = &cobra.Command{
	Use:   "grid",
	Short: "The DeepSquare Grid Client CLI",
}

func init() {
	app.AddCommand(submit.Command)
}

func main() {
	if err := app.Execute(); err != nil {
		logger.I.Fatal("app crashed", zap.Error(err))
	}
}
