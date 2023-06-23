//go:build ignore

package main

import (
	"fmt"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin/modelgen"
	"github.com/deepsquare-io/the-grid/cli/internal/log"
	"github.com/vektah/gqlparser/v2/ast"
)

func constraintFieldHook(
	td *ast.Definition,
	fd *ast.FieldDefinition,
	f *modelgen.Field,
) (*modelgen.Field, error) {
	if f, err := modelgen.DefaultFieldMutateHook(td, fd, f); err != nil {
		return f, err
	}

	c := fd.Directives.ForName("constraint")
	if c != nil {
		formatConstraint := c.Arguments.ForName("format")

		if formatConstraint != nil {
			f.Tag += " validate:" + formatConstraint.Value.String()
		}
	}

	return f, nil
}

func main() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	// Attaching the mutation function onto modelgen plugin
	p := modelgen.Plugin{
		FieldHook: constraintFieldHook,
	}

	if err = api.Generate(cfg, api.ReplacePlugin(&p)); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}

	if err = os.Remove("./graph/resolver.go"); err != nil {
		log.I.Warn(err.Error())
	}
	if err = os.Remove("./graph/schema.resolvers.go"); err != nil {
		log.I.Warn(err.Error())
	}
	if err = os.Remove("./graph"); err != nil {
		log.I.Warn(err.Error())
	}
	if err = os.Remove("./sbatch/generated.go"); err != nil {
		log.I.Warn(err.Error())
	}
}
