package font

import (
	"text/template"
)

const t = `{{range .}}
@font-face {
  font-family: '{{.Family}}';
  src: local('{{.Filename}}'), url('{{.BaseURL}}/{{.Family}}/{{.Filename}}') format('woff2');
  font-weight: {{.Weight}};
  font-style: {{.Style}};
}
{{end}}
`

func CSSTemplate() *template.Template {
	return template.Must(template.New("template.css").Parse(t))
}
