package main

import (
	"github.com/deepsquare-io/the-grid/cli/cmd/submit"
	"github.com/deepsquare-io/the-grid/cli/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var debug bool

var app = &cobra.Command{
	Use:   "grid",
	Short: "The DeepSquare Grid Client CLI",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if debug {
			logger.EnableDebug()
		}
	},
}

func init() {
	flags := app.PersistentFlags()
	flags.BoolVarP(
		&debug,
		"debug",
		"d",
		false,
		"Show debug logging.",
	)
	viper.BindPFlag("DEBUG", flags.Lookup("debug"))
	app.AddCommand(submit.Command)
}

func main() {
	if err := app.Execute(); err != nil {
		logger.I.Fatal("app crashed", zap.Error(err))
	}
}
