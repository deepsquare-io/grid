package template

import (
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

func funcMap() template.FuncMap {
	f := sprig.TxtFuncMap()
	f["deref"] = func(i *string) string { return *i }
	return f
}

func Init() *template.Template {
	return template.New("gotpl").Funcs(funcMap())
}
