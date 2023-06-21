package cmd

import (
	"fmt"
	"os"

	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/logger"
	"github.com/deepsquare-io/the-grid/sbatch-service/renderer"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

var (
	file      string
	RenderCmd = cli.Command{
		Name:  "render",
		Usage: "Validate and render a YAML or JSON to sbatch script",
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:        "file",
				Value:       "script.yaml",
				Usage:       "File to convert.",
				Destination: &file,
				Aliases:     []string{"f"},
			},
		},
		Action: func(ctx *cli.Context) error {
			r := renderer.NewJobRenderer(
				"logger.example.com:443",
				"/usr/local/bin/grid-logger-writer",
			)
			bytes, err := os.ReadFile(file)
			if err != nil {
				logger.I.Error("failed to read file", zap.Error(err))
				return err
			}
			j := struct {
				Job model.Job
			}{}
			if err := yaml.Unmarshal(bytes, &j); err != nil {
				logger.I.Error("failed to parse file", zap.Error(err))
				return err
			}
			out, err := r.RenderJob(&j.Job)
			if err != nil {
				logger.I.Error("failed to render sbatch script from file", zap.Error(err))
				return err
			}
			fmt.Println(out)
			return nil
		},
	}
)
