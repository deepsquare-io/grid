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

package module

import (
	"bytes"
	"reflect"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/deepsquare-io/grid/sbatch-service/graph/model"
	"gopkg.in/yaml.v3"
)

func Render(j *model.Job, s *model.Step, template string) (string, error) {
	tmpl, err := engine().Parse(template)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	if err = tmpl.Execute(&out, struct {
		Job  *model.Job
		Step *model.Step
	}{
		Job:  j,
		Step: s,
	}); err != nil {
		return "", err
	}
	return out.String(), nil
}

func isZeroOfUnderlyingType(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

func toYAML(v interface{}) string {
	if isZeroOfUnderlyingType(v) {
		return ""
	}
	data, err := yaml.Marshal(v)
	if err != nil {
		// Swallow errors inside of a template.
		return ""
	}
	return strings.TrimSuffix(string(data), "\n")
}

func funcMap() template.FuncMap {
	f := sprig.TxtFuncMap()
	f["toYaml"] = toYAML
	delete(f, "env")
	delete(f, "expandenv")
	return f
}

func engine() *template.Template {
	return template.New("gotpl").Funcs(funcMap())
}
