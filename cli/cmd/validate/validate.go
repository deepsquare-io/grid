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

/*
Package submit permits the submission of a job to the DeepSquare Grid.

USAGE:

	dps validate [command options] <job.yaml>

OPTIONS:

	--sbatch.endpoint value  SBatch Service GraphQL endpoint. (default: "https://sbatch.deepsquare.run/graphql") [$SBATCH_ENDPOINT]
*/
package validate

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/deepsquare-io/grid/cli/deepsquare"
	"github.com/deepsquare-io/grid/cli/sbatch"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

var (
	sbatchEndpoint string
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "sbatch.endpoint",
		Value:       deepsquare.DefaultSBatchEndpoint,
		Usage:       "SBatch Service GraphQL endpoint.",
		Destination: &sbatchEndpoint,
		EnvVars:     []string{"SBATCH_ENDPOINT"},
	},
}

// Command is the validate subcommand used to submit jobs.
var Command = cli.Command{
	Name:      "validate",
	Usage:     "Quickly validate a job.",
	ArgsUsage: "<job.yaml>",
	Flags:     flags,
	Action: func(cCtx *cli.Context) error {
		if cCtx.NArg() < 1 {
			return errors.New("missing arguments")
		}
		jobPath := cCtx.Args().First()
		ctx := cCtx.Context
		client := sbatch.NewService(http.DefaultClient, sbatchEndpoint)
		dat, err := os.ReadFile(jobPath)
		if err != nil {
			return err
		}
		var job sbatch.Job
		if err := yaml.Unmarshal(dat, &job); err != nil {
			return err
		}
		_, err = client.Submit(ctx, &job)
		if err == nil {
			fmt.Println("valid")
		}
		return err
	},
}
