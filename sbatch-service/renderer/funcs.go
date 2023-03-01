package renderer

import (
	"fmt"
	"net"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
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

func FormatImageURL(registry *string, image string, apptainer *bool, deepsquareHosted *bool) string {
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
			return fmt.Sprintf("docker://%s/%s", *registry, image)
		}
		return fmt.Sprintf("docker://%s", image)
	}

	// Is enroot
	if registry != nil && *registry != "" {
		return fmt.Sprintf("%s#%s", *registry, image)
	}
	return image
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
	f["renderApptainerCommand"] = RenderApptainerCommand
	f["renderSlirp4NetNS"] = RenderSlirp4NetNS
	f["renderEnrootCommand"] = RenderEnrootCommand
	f["squote"] = squote
	f["escapeSQuote"] = escapeSQuote
	f["quoteEscape"] = quoteEscape
	f["formatImageURL"] = FormatImageURL
	f["isCIDRv4"] = func(i string) bool {
		ip, _, err := net.ParseCIDR(i)

		return err == nil && ip.To4() != nil
	}
	f["isCIDRv6"] = func(i string) bool {
		ip, _, err := net.ParseCIDR(i)

		return err == nil && ip.To4() == nil
	}
	return f
}

func engine() *template.Template {
	return template.New("gotpl").Funcs(funcMap())
}
