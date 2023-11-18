// Copyright (C) 2023 DeepSquare Association
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
Package initc provides subcommands to initialize templates to get started with DeepSquare.
It will initialize a job.<generated name>.yaml with a JSON schema in the cache or /tmp directory.

USAGE:

	deepsquaretui init [command options] [arguments...]

OPTIONS:

	--output value, -o value  Output path.
*/
package initc

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "embed"

	"github.com/deepsquare-io/grid/cli/internal/log"
	"github.com/deepsquare-io/grid/cli/internal/wordlists"
	"github.com/urfave/cli/v2"
)

var (
	output string

	//go:embed template.yaml
	template []byte

	//go:embed job.schema.json
	schema []byte
)

const templateFormat = "# yaml-language-server: $schema=%s\n%s"

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "output",
		Usage:       "Output path.",
		Destination: &output,
		Aliases:     []string{"o"},
	},
}

func prepareFiles() (jerr error) {
	dir, err := os.UserCacheDir()
	if err != nil {
		log.I.Warn("couldn't fetch user cache dir, using tmp...")
		dir = os.TempDir()
	}
	words := strings.Join(wordlists.GetRandomWords(3), "-")
	jobSchemaPath := filepath.Join(dir, ".job.schema.json")
	jobPath := filepath.Join(output, fmt.Sprintf("job.%s.yaml", words))

	// Insert the yaml-language-server parameter
	template = []byte(
		fmt.Sprintf(templateFormat, jobSchemaPath, template),
	)

	if err := os.WriteFile(jobSchemaPath, schema, 0644); err != nil {
		return fmt.Errorf("fail to write %s: %w", jobSchemaPath, err)
	}

	if err := os.WriteFile(jobPath, template, 0644); err != nil {
		return fmt.Errorf("fail to write %s: %w", jobPath, err)
	}

	return nil
}

// Command is the init command used to bootstrap a job workflow file.
var Command = cli.Command{
	Name:  "init",
	Usage: "Bootstrap a job workflow file.",
	Flags: flags,
	Action: func(cCtx *cli.Context) error {
		return prepareFiles()
	},
}
