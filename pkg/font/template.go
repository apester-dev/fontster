package font

import "text/template"

const t = `{{range .}}@font-face {
  font-family: '{{.Name}}';
  src: local('{{.Filename}}'), url('{{.BaseURL}}/{{.Name}}/{{.Filename}}.woff2') format('woff2');
  font-weight: {{.Weight}};
  font-style: {{.Style}};
}
{{end}}
`

// CSSTemplate returns parsed CSS template.
func CSSTemplate() *template.Template {
	return template.Must(template.New("template.css").Parse(t))
}
