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
Package initc provides subcommands to initialize templates to get started with DeepSquare.
It will initialize a job.<generated name>.yaml with a JSON schema in the cache or /tmp directory.

USAGE:

	dps init [command options] [arguments...]

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

	"github.com/deepsquare-io/grid/cli/internal/utils"
	"github.com/deepsquare-io/grid/cli/internal/wordlists"
	"github.com/urfave/cli/v2"
)

var (
	output string

	//go:embed template.yaml
	template []byte

	ethHexPK     string
	ethHexPKPath string
)

const jobSchemaPath = "https://raw.githubusercontent.com/deepsquare-io/grid/main/cli/job.schema.json"
const templateFormat = "# yaml-language-server: $schema=%s\n%s"

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "output",
		Usage:       "Output path.",
		Destination: &output,
		Aliases:     []string{"o"},
	},
	&cli.StringFlag{
		Name:        "private-key",
		Usage:       "An hexadecimal private key for ethereum transactions.",
		Value:       "",
		Destination: &ethHexPK,
		Category:    "DeepSquare Settings:",
		EnvVars:     []string{"ETH_PRIVATE_KEY"},
	},
	&cli.StringFlag{
		Name:        "private-key.path",
		Usage:       "Path to an hexadecimal private key for ethereum transactions.",
		Destination: &ethHexPKPath,
		EnvVars:     []string{"ETH_PRIVATE_KEY_PATH"},
		Category:    "DeepSquare Settings:",
		Value:       "",
	},
}

func prepareFiles() (jerr error) {
	words := strings.Join(wordlists.GetRandomWords(3), "-")
	jobPath := filepath.Join(output, fmt.Sprintf("job.%s.yaml", words))

	// Insert the yaml-language-server parameter
	template = []byte(
		fmt.Sprintf(templateFormat, jobSchemaPath, template),
	)

	if err := os.WriteFile(jobPath, template, 0644); err != nil {
		return fmt.Errorf("fail to write %s: %w", jobPath, err)
	}

	fmt.Printf("Job file created at %s\n", jobPath)

	return nil
}

// Command is the init command used to bootstrap a job workflow file.
var Command = cli.Command{
	Name:  "init",
	Usage: "Bootstrap a job workflow file.",
	Flags: flags,
	Action: func(_ *cli.Context) error {
		_, err := utils.GetPrivateKey(ethHexPK, ethHexPKPath)
		if err != nil {
			return err
		}
		return prepareFiles()
	},
}
