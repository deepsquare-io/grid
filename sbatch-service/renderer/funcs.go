package renderer

import (
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

func funcMap() template.FuncMap {
	f := sprig.TxtFuncMap()
	f["deref"] = func(i *string) string { return *i }
	f["renderStep"] = RenderStep
	return f
}

func engine() *template.Template {
	return template.New("gotpl").Funcs(funcMap())
}
