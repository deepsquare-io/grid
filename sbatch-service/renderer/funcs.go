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

package renderer

import (
	"fmt"
	"net"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/deepsquare-io/grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/grid/sbatch-service/utils/base58"
	"github.com/kballard/go-shellquote"
)

func squote(str ...interface{}) string {
	out := make([]string, 0, len(str))
	for _, s := range str {
		if s != nil {
			switch s := s.(type) {
			case string:
				s = strings.ReplaceAll(s, "'", "'\"'\"'")
				out = append(out, fmt.Sprintf("'%v'", s))
			default:
				out = append(out, fmt.Sprintf("'%v'", s))
			}
		}
	}
	return strings.Join(out, " ")
}

func escapeSQuote(str ...interface{}) string {
	out := make([]string, 0, len(str))
	for _, s := range str {
		if s != nil {
			switch s := s.(type) {
			case string:
				s = strings.ReplaceAll(s, "'", "'\"'\"'")
				out = append(out, s)
			default:
				out = append(out, fmt.Sprintf("%v", s))
			}
		}
	}
	return strings.Join(out, " ")
}

func quoteEscape(str ...interface{}) string {
	out := make([]string, 0, len(str))
	for _, s := range str {
		if s != nil {
			switch s := s.(type) {
			case string:
				s = strings.ReplaceAll(s, "\"", "\\\"")
				out = append(out, fmt.Sprintf("\"%v\"", s))
			default:
				out = append(out, fmt.Sprintf("\"%v\"", s))
			}
		}
	}
	return strings.Join(out, " ")
}

func ignoreNil(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func FormatImageURL(
	registry *string,
	image string,
	apptainer *bool,
	deepsquareHosted *bool,
) string {
	// Is absolute path
	if strings.HasPrefix(image, "/") {
		return filepath.Clean(image)
	}

	if deepsquareHosted != nil && *deepsquareHosted {
		if registry != nil && *registry != "" {
			return fmt.Sprintf("/opt/software/%s/%s", *registry, image)
		}
		return fmt.Sprintf("/opt/software/%s", image)
	}

	// Is apptainer
	if apptainer != nil && *apptainer {
		if registry != nil && *registry != "" {
			if strings.Contains(*registry, "docker.io") {
				_, _, ok := strings.Cut(image, "/")
				if !ok {
					image = "library/" + image
				}
			}
			return fmt.Sprintf("docker://%s/%s", *registry, image)
		}
		return fmt.Sprintf("docker://%s", image)
	}

	// Is enroot
	if registry != nil && *registry != "" {
		if strings.Contains(*registry, "docker.io") {
			_, _, ok := strings.Cut(image, "/")
			if !ok {
				image = "library/" + image
			}
		}
		return fmt.Sprintf("docker://%s#%s", *registry, image)
	}
	return fmt.Sprintf("docker://%s", image)
}

func funcMap() template.FuncMap {
	f := sprig.TxtFuncMap()
	f["derefStr"] = func(i *string) string { return *i }
	f["derefBool"] = func(i *bool) bool { return *i }
	f["derefInt"] = func(i *int) int { return *i }
	f["octal"] = func(i int) string { return fmt.Sprintf("%o", i) }
	f["renderStep"] = RenderStep
	f["renderStepRun"] = RenderStepRun
	f["renderStepFor"] = RenderStepFor
	f["renderWireguard"] = RenderWireguard
	f["renderVirtualNetwork"] = RenderVirtualNetwork
	f["renderVNet"] = RenderVNet
	f["renderApptainerCommand"] = RenderApptainerCommand
	f["renderSlirp4NetNS"] = RenderSlirp4NetNS
	f["renderPastaNS"] = RenderPastaNS
	f["renderEnrootCommand"] = RenderEnrootCommand
	f["renderStepAsyncLaunch"] = RenderStepAsyncLaunch
	f["renderStepUse"] = func(
		j *model.Job,
		s *model.Step,
		u *model.StepUse,
	) (string, error) {
		return NewStepUseRenderer(base58.Encoder{}).Render(j, s, u)
	}
	f["squote"] = squote
	f["escapeSQuote"] = escapeSQuote
	f["quoteEscape"] = quoteEscape
	f["formatImageURL"] = FormatImageURL
	f["ignoreNil"] = ignoreNil
	f["isCIDRv4"] = func(i string) bool {
		ip, _, err := net.ParseCIDR(i)

		return err == nil && ip.To4() != nil
	}
	f["isCIDRv6"] = func(i string) bool {
		ip, _, err := net.ParseCIDR(i)

		return err == nil && ip.To4() == nil
	}
	f["escapeCommand"] = func(i string) (string, error) {
		split, err := shellquote.Split(i)
		if err != nil {
			return "", err
		}
		return shellquote.Join(split...), nil
	}
	return f
}

func engine() *template.Template {
	return template.New("gotpl").Funcs(funcMap())
}
