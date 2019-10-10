package font

import (
	"html/template"
	"strings"
)

const t = `{{range . -}}
@font-face {
  font-family: '{{.Name}}';
  src: local('{{.FileName}}'), url('{{.Source}}.woff2') format('woff2');
  font-weight: {{ extractWeight .Weight}};
  font-style: {{.Style}};
}
{{end -}}
`

// CSSTemplate returns parsed CSS template.
func CSSTemplate() *template.Template {
	return template.Must(template.New("template.css").Funcs(template.FuncMap{
		"extractWeight": func(weight string) string {
			return strings.Trim(weight, "i")
		},
	}).Parse(t))
}
