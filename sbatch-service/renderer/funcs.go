package renderer

import (
	"fmt"
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
				s = strings.ReplaceAll(s, "'", "'\\''")
				out = append(out, fmt.Sprintf("'%v'", s))
			default:
				out = append(out, fmt.Sprintf("'%v'", s))
			}
		}
	}
	return strings.Join(out, " ")
}

func funcMap() template.FuncMap {
	f := sprig.TxtFuncMap()
	f["deref"] = func(i *string) string { return *i }
	f["renderStep"] = RenderStep
	f["squote"] = squote
	return f
}

func engine() *template.Template {
	return template.New("gotpl").Funcs(funcMap())
}
