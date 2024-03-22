// Copyright (C) 2024 DeepSquare Association
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package cmd

import (
	"fmt"
	"os"

	"github.com/deepsquare-io/grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/grid/sbatch-service/logger"
	"github.com/deepsquare-io/grid/sbatch-service/renderer"
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
		Action: func(_ *cli.Context) error {
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
