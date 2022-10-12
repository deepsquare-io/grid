package main

import (
	"os"

	"github.com/deepsquare-io/the-grid/cli/cmd/submit"
	"github.com/deepsquare-io/the-grid/cli/logger"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var app = &cli.App{
	Name:  "grid",
	Usage: "The DeepSquare Grid Client CLI",
	Commands: []*cli.Command{
		submit.Command,
	},
}

func main() {
	if err := app.Run(os.Args); err != nil {
		logger.I.Fatal("app crashed", zap.Error(err))
	}
}
