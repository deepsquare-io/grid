package initc

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	_ "embed"

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
	date := time.Now().Unix()
	jobSchemaPath := filepath.Join(output, ".job.schema.json")
	jobPath := filepath.Join(output, fmt.Sprintf("job.%d.yaml", date))

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

var Command = cli.Command{
	Name:  "init",
	Usage: "Bootstrap a job workflow file.",
	Flags: flags,
	Action: func(cCtx *cli.Context) error {
		return prepareFiles()
	},
}
